package obo

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/db"
	"github.com/bytebase/bytebase/backend/plugin/db/util"
)

const systemSchemas = "'DWEXP','OMC','ORAAUDITOR','LBACSYS','SYS'"

var semVersionRegex = regexp.MustCompile(`[0-9]+(\.[0-9])*`)

func (d *Driver) SyncInstance(ctx context.Context) (*db.InstanceMetadata, error) {
	var fullVersion string
	if err := d.db.QueryRowContext(ctx, "SELECT BANNER FROM v$version").Scan(&fullVersion); err != nil {
		return nil, errors.Wrapf(err, "failed to get version")
	}
	version := fullVersion
	for _, token := range strings.Fields(fullVersion) {
		if semVersionRegex.MatchString(token) {
			version = token
			break
		}
	}

	query := fmt.Sprintf(`
		SELECT username FROM sys.all_users
		WHERE username NOT IN (%s)
		ORDER BY username
	`, systemSchemas)
	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schemas []string
	for rows.Next() {
		var schema string
		if err := rows.Scan(&schema); err != nil {
			return nil, err
		}
		schemas = append(schemas, schema)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	instance := &db.InstanceMetadata{
		Version:   version,
		Databases: nil,
	}
	for _, schema := range schemas {
		instance.Databases = append(instance.Databases, &storepb.DatabaseSchemaMetadata{
			Name: schema,
		})
	}
	return instance, nil
}

func (d *Driver) SyncDBSchema(ctx context.Context) (*storepb.DatabaseSchemaMetadata, error) {
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	columnMap, err := getTableColumns(ctx, tx, d.databaseName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get table columns from database %q", d.databaseName)
	}
	tableMap, err := getTables(ctx, tx, d.databaseName, columnMap)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get tables from database %q", d.databaseName)
	}
	viewMap, err := getViews(ctx, tx, d.databaseName, columnMap)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get views from database %q", d.databaseName)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	databaseMetadata := &storepb.DatabaseSchemaMetadata{
		Name: d.databaseName,
	}
	databaseMetadata.Schemas = append(databaseMetadata.Schemas, &storepb.SchemaMetadata{
		Name:   "",
		Tables: tableMap[d.databaseName],
		Views:  viewMap[d.databaseName],
	})
	return databaseMetadata, nil
}

func getTables(ctx context.Context, txn *sql.Tx, schemaName string, columnMap map[db.TableKey][]*storepb.ColumnMetadata) (map[string][]*storepb.TableMetadata, error) {
	indexMap, err := getIndexes(ctx, txn, schemaName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get indexes")
	}
	// TODO(d): foreign keys.
	tableMap := make(map[string][]*storepb.TableMetadata)
	query := fmt.Sprintf(`
		SELECT OWNER, TABLE_NAME, NUM_ROWS
		FROM all_tables
		WHERE OWNER = '%s'
		ORDER BY TABLE_NAME`, schemaName)

	rows, err := txn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		table := &storepb.TableMetadata{}
		var schemaName string
		// https://github.com/rana/ora/issues/57#issuecomment-179909837
		// NUMBER in Oracle can hold 38 decimal digits, so int64 is not enough with its 19 decimal digits.
		// float64 is a little bit better - not precise enough, but won't overflow.
		var count sql.NullFloat64
		if err := rows.Scan(&schemaName, &table.Name, &count); err != nil {
			return nil, err
		}
		table.RowCount = int64(count.Float64)
		key := db.TableKey{Schema: schemaName, Table: table.Name}
		table.Columns = columnMap[key]
		table.Indexes = indexMap[key]

		tableMap[schemaName] = append(tableMap[schemaName], table)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tableMap, nil
}

func getTableColumns(ctx context.Context, txn *sql.Tx, schemaName string) (map[db.TableKey][]*storepb.ColumnMetadata, error) {
	columnsMap := make(map[db.TableKey][]*storepb.ColumnMetadata)
	query := fmt.Sprintf(`
	SELECT
		OWNER,
		TABLE_NAME,
		COLUMN_NAME,
		DATA_TYPE,
		COLUMN_ID,
		DATA_DEFAULT,
		NULLABLE
	FROM sys.all_tab_columns
	WHERE OWNER = '%s' AND COLUMN_ID IS NOT NULL
	ORDER BY TABLE_NAME, COLUMN_ID`, schemaName)

	rows, err := txn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		column := &storepb.ColumnMetadata{}
		var schemaName, tableName, nullable string
		var defaultStr sql.NullString
		if err := rows.Scan(&schemaName, &tableName, &column.Name, &column.Type, &column.Position, &defaultStr, &nullable); err != nil {
			return nil, err
		}
		if defaultStr.Valid {
			// Store in Default field (migration from DefaultExpression to Default)
			column.Default = defaultStr.String
		}
		isNullBool, err := util.ConvertYesNo(nullable)
		if err != nil {
			return nil, err
		}
		column.Nullable = isNullBool
		// TODO(d): add collation.

		key := db.TableKey{Schema: schemaName, Table: tableName}
		columnsMap[key] = append(columnsMap[key], column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columnsMap, nil
}

func getIndexes(ctx context.Context, txn *sql.Tx, schemaName string) (map[db.TableKey][]*storepb.IndexMetadata, error) {
	indexMap := make(map[db.TableKey][]*storepb.IndexMetadata)

	expressionsMap := make(map[db.IndexKey][]string)

	queryColumn := fmt.Sprintf(`
		SELECT TABLE_OWNER, TABLE_NAME, INDEX_NAME, COLUMN_NAME
		FROM sys.all_ind_columns
		WHERE TABLE_OWNER = '%s'
		ORDER BY TABLE_NAME, INDEX_NAME, COLUMN_POSITION`, schemaName)

	colRows, err := txn.QueryContext(ctx, queryColumn)
	if err != nil {
		return nil, err
	}
	defer colRows.Close()
	for colRows.Next() {
		var schemaName, tableName, indexName, columnName string
		if err := colRows.Scan(&schemaName, &tableName, &indexName, &columnName); err != nil {
			return nil, err
		}
		key := db.IndexKey{Schema: schemaName, Table: tableName, Index: indexName}
		expressionsMap[key] = append(expressionsMap[key], columnName)
	}
	if err := colRows.Err(); err != nil {
		return nil, err
	}
	queryExpression := fmt.Sprintf(`
		SELECT TABLE_OWNER, TABLE_NAME, INDEX_NAME, COLUMN_EXPRESSION, COLUMN_POSITION
		FROM sys.all_ind_expressions
		WHERE TABLE_OWNER = '%s'
		ORDER BY TABLE_NAME, INDEX_NAME, COLUMN_POSITION`, schemaName)
	expRows, err := txn.QueryContext(ctx, queryExpression)
	if err != nil {
		return nil, err
	}
	defer expRows.Close()
	for expRows.Next() {
		var schemaName, tableName, indexName, columnExpression string
		var position int
		if err := expRows.Scan(&schemaName, &tableName, &indexName, &columnExpression, &position); err != nil {
			return nil, err
		}
		key := db.IndexKey{Schema: schemaName, Table: tableName, Index: indexName}
		// Position starts from 1.
		expIndex := position - 1
		if expIndex >= len(expressionsMap[key]) {
			return nil, errors.Errorf("expression %q position %v out of range for index %q.%q.%q", columnExpression, position, schemaName, tableName, indexName)
		}
		expressionsMap[key][expIndex] = columnExpression
	}
	if err := expRows.Err(); err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`
		SELECT OWNER, TABLE_NAME, INDEX_NAME, UNIQUENESS, INDEX_TYPE
		FROM sys.all_indexes
		WHERE OWNER = '%s'
		ORDER BY TABLE_NAME, INDEX_NAME`, schemaName)
	rows, err := txn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		index := &storepb.IndexMetadata{}
		var schemaName, tableName, unique string
		// INDEX_TYPE is NORMAL, or FUNCTION-BASED NORMAL.
		if err := rows.Scan(&schemaName, &tableName, &index.Name, &unique, &index.Type); err != nil {
			return nil, err
		}

		index.Unique = unique == "UNIQUE"
		indexKey := db.IndexKey{Schema: schemaName, Table: tableName, Index: index.Name}
		index.Expressions = expressionsMap[indexKey]

		key := db.TableKey{Schema: schemaName, Table: tableName}
		indexMap[key] = append(indexMap[key], index)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return indexMap, nil
}

func getViews(ctx context.Context, txn *sql.Tx, schemaName string, columnMap map[db.TableKey][]*storepb.ColumnMetadata) (map[string][]*storepb.ViewMetadata, error) {
	viewMap := make(map[string][]*storepb.ViewMetadata)

	query := fmt.Sprintf(`
		SELECT OWNER, VIEW_NAME, TEXT
		FROM sys.all_views
		WHERE OWNER = '%s'
		ORDER BY view_name
	`, schemaName)

	rows, err := txn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		view := &storepb.ViewMetadata{}
		var schemaName string
		if err := rows.Scan(&schemaName, &view.Name, &view.Definition); err != nil {
			return nil, err
		}
		key := db.TableKey{Schema: schemaName, Table: view.Name}
		view.Columns = columnMap[key]

		viewMap[schemaName] = append(viewMap[schemaName], view)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return viewMap, nil
}
