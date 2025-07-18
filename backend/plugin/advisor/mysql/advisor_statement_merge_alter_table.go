package mysql

// Framework code is generated by the generator.

import (
	"context"
	"fmt"
	"slices"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
)

var (
	_ advisor.Advisor = (*StatementMergeAlterTableAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLMergeAlterTable, &StatementMergeAlterTableAdvisor{})
}

// StatementMergeAlterTableAdvisor is the advisor checking for merging ALTER TABLE statements.
type StatementMergeAlterTableAdvisor struct {
}

// Check checks for merging ALTER TABLE statements.
func (*StatementMergeAlterTableAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementMergeAlterTableChecker{
		level:    level,
		title:    string(checkCtx.Rule.Type),
		tableMap: make(map[string]tableStatement),
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.generateAdvice(), nil
}

type statementMergeAlterTableChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	text       string
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	tableMap   map[string]tableStatement
}

type tableStatement struct {
	name     string
	count    int
	lastLine int
}

func (checker *statementMergeAlterTableChecker) EnterQuery(ctx *mysql.QueryContext) {
	checker.text = ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
}

// EnterCreateTable is called when production createTable is entered.
func (checker *statementMergeAlterTableChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableName() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableName(ctx.TableName())
	checker.tableMap[tableName] = tableStatement{
		name:     tableName,
		count:    1,
		lastLine: checker.baseLine + ctx.GetStart().GetLine(),
	}
}

// EnterAlterTable is called when production alterTable is entered.
func (checker *statementMergeAlterTableChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableRef() == nil {
		return
	}
	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	table, ok := checker.tableMap[tableName]
	if !ok {
		table = tableStatement{
			name:  tableName,
			count: 0,
		}
	}
	table.count++
	table.lastLine = checker.baseLine + ctx.GetStart().GetLine()
	checker.tableMap[tableName] = table
}

func (checker *statementMergeAlterTableChecker) generateAdvice() []*storepb.Advice {
	var tableList []tableStatement
	for _, table := range checker.tableMap {
		tableList = append(tableList, table)
	}
	slices.SortFunc(tableList, func(i, j tableStatement) int {
		if i.lastLine < j.lastLine {
			return -1
		}
		if i.lastLine > j.lastLine {
			return 1
		}
		return 0
	})

	for _, table := range tableList {
		if table.count > 1 {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:        checker.level,
				Code:          advisor.StatementRedundantAlterTable.Int32(),
				Title:         checker.title,
				Content:       fmt.Sprintf("There are %d statements to modify table `%s`", table.count, table.name),
				StartPosition: common.ConvertANTLRLineToPosition(table.lastLine),
			})
		}
	}

	return checker.adviceList
}
