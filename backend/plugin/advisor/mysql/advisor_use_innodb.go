package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
)

const (
	innoDB              string = "innodb"
	defaultStorageEngin string = "default_storage_engine"
)

var _ advisor.Advisor = (*UseInnoDBAdvisor)(nil)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLUseInnoDB, &UseInnoDBAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLUseInnoDB, &UseInnoDBAdvisor{})
}

// UseInnoDBAdvisor is the advisor checking for using InnoDB engine.
type UseInnoDBAdvisor struct {
}

// Check checks for using InnoDB engine.
func (*UseInnoDBAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	list, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	checker := &useInnoDBChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	for _, stmt := range list {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type useInnoDBChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
}

// EnterCreateTable is called when production createTable is entered.
func (c *useInnoDBChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.CreateTableOptions() == nil {
		return
	}
	for _, tableOption := range ctx.CreateTableOptions().AllCreateTableOption() {
		if tableOption.ENGINE_SYMBOL() != nil && tableOption.EngineRef() != nil {
			if tableOption.EngineRef().TextOrIdentifier() == nil {
				continue
			}
			engine := mysqlparser.NormalizeMySQLTextOrIdentifier(tableOption.EngineRef().TextOrIdentifier())
			if strings.ToLower(engine) != innoDB {
				content := "CREATE " + ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
				line := tableOption.GetStart().GetLine()
				c.addAdvice(content, line)
				break
			}
		}
	}
}

func (c *useInnoDBChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.AlterTableActions() == nil || ctx.AlterTableActions().AlterCommandList() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList().AlterList() == nil {
		return
	}
	code := advisor.Ok
	for _, option := range ctx.AlterTableActions().AlterCommandList().AlterList().AllCreateTableOptionsSpaceSeparated() {
		for _, op := range option.AllCreateTableOption() {
			if op.ENGINE_SYMBOL() != nil {
				if op.EngineRef() == nil {
					continue
				}
				engine := op.EngineRef().GetText()
				if strings.ToLower(engine) != innoDB {
					code = advisor.NotInnoDBEngine
					break
				}
			}
		}
	}

	if code != advisor.Ok {
		content := "ALTER " + ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
		line := ctx.GetStart().GetLine()
		c.addAdvice(content, line)
	}
}

func (c *useInnoDBChecker) EnterSetStatement(ctx *mysql.SetStatementContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	code := advisor.Ok
	if ctx.StartOptionValueList() == nil {
		return
	}

	startOptionValueList := ctx.StartOptionValueList()
	if startOptionValueList.OptionValueNoOptionType() == nil {
		return
	}
	optionValueNoOptionType := startOptionValueList.OptionValueNoOptionType()
	if optionValueNoOptionType.InternalVariableName() == nil {
		return
	}
	name := optionValueNoOptionType.InternalVariableName().GetText()
	if strings.ToLower(name) != defaultStorageEngin {
		return
	}
	if optionValueNoOptionType.SetExprOrDefault() != nil {
		engine := optionValueNoOptionType.SetExprOrDefault().GetText()
		if strings.ToLower(engine) != innoDB {
			code = advisor.NotInnoDBEngine
		}
	}

	if code != advisor.Ok {
		content := ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
		line := ctx.GetStart().GetLine()
		c.addAdvice(content, line)
	}
}

func (c *useInnoDBChecker) addAdvice(content string, lineNumber int) {
	lineNumber += c.baseLine
	c.adviceList = append(c.adviceList, &storepb.Advice{
		Status:        c.level,
		Code:          advisor.NotInnoDBEngine.Int32(),
		Title:         c.title,
		Content:       fmt.Sprintf("\"%s;\" doesn't use InnoDB engine", content),
		StartPosition: common.ConvertANTLRLineToPosition(lineNumber),
	})
}
