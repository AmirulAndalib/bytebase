package mysql

import (
	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
)

type queryTypeListener struct {
	*mysql.BaseMySQLParserListener

	result base.QueryType
}

func (l *queryTypeListener) EnterSimpleStatement(ctx *mysql.SimpleStatementContext) {
	if !isTopLevel(ctx) {
		return
	}

	switch {
	case ctx.AlterStatement() != nil,
		ctx.CreateStatement() != nil,
		ctx.DropStatement() != nil,
		ctx.RenameTableStatement() != nil,
		ctx.TruncateTableStatement() != nil,
		ctx.ImportStatement() != nil:
		l.result = base.DDL
	case ctx.CallStatement() != nil,
		ctx.DeleteStatement() != nil,
		ctx.DoStatement() != nil,
		ctx.HandlerStatement() != nil,
		ctx.InsertStatement() != nil,
		ctx.LoadStatement() != nil,
		ctx.ReplaceStatement() != nil,
		ctx.UpdateStatement() != nil,
		ctx.TransactionOrLockingStatement() != nil,
		ctx.ReplicationStatement() != nil,
		ctx.PreparedStatement() != nil:
		l.result = base.DML
	case ctx.ShowStatement() != nil:
		l.result = base.SelectInfoSchema
	case ctx.UtilityStatement() != nil:
		if ctx.UtilityStatement().DescribeStatement() != nil {
			l.result = base.SelectInfoSchema
		}
		if ctx.UtilityStatement().ExplainStatement() != nil {
			l.result = base.Explain
		}
	}
}

func (l *queryTypeListener) EnterSelectStatement(ctx *mysql.SelectStatementContext) {
	if !isTopLevel(ctx.GetParent()) {
		return
	}

	// MySQL cannot use SELECT ... INTO .. FROM ... syntax to create a new table or insert into an existing table.
	// So we can safely assume it's a SELECT statement.

	accessTables := getAccessTables("", ctx)
	allSystems, _ := isMixedQuery(accessTables, true)
	if allSystems {
		l.result = base.SelectInfoSchema
	} else {
		l.result = base.Select
	}
}