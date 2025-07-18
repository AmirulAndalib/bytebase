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
	_ advisor.Advisor = (*IndexTypeAllowListAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLIndexTypeAllowList, &IndexTypeAllowListAdvisor{})
}

// IndexTypeAllowListAdvisor is the advisor checking for index types.
type IndexTypeAllowListAdvisor struct {
}

func (*IndexTypeAllowListAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parser result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}

	checker := &indexTypeAllowListChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}
	payload, err := advisor.UnmarshalStringArrayTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker.allowList = payload.List

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type indexTypeAllowListChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	allowList  []string
}

func (checker *indexTypeAllowListChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableName() == nil {
		return
	}
	if ctx.TableElementList() == nil {
		return
	}

	for _, tableElement := range ctx.TableElementList().AllTableElement() {
		if tableElement == nil || tableElement.TableConstraintDef() == nil {
			continue
		}
		checker.handleConstraintDef(tableElement.TableConstraintDef())
	}
}

func (checker *indexTypeAllowListChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableRef() == nil {
		return
	}
	if ctx.AlterTableActions() == nil || ctx.AlterTableActions().AlterCommandList() == nil || ctx.AlterTableActions().AlterCommandList().AlterList() == nil {
		return
	}

	for _, alterListItem := range ctx.AlterTableActions().AlterCommandList().AlterList().AllAlterListItem() {
		if alterListItem == nil {
			continue
		}
		if alterListItem.ADD_SYMBOL() != nil && alterListItem.TableConstraintDef() != nil {
			checker.handleConstraintDef(alterListItem.TableConstraintDef())
		}
	}
}

func (checker *indexTypeAllowListChecker) handleConstraintDef(ctx mysql.ITableConstraintDefContext) {
	switch ctx.GetType_().GetTokenType() {
	case mysql.MySQLParserINDEX_SYMBOL, mysql.MySQLParserKEY_SYMBOL, mysql.MySQLParserPRIMARY_SYMBOL, mysql.MySQLParserUNIQUE_SYMBOL, mysql.MySQLParserFULLTEXT_SYMBOL, mysql.MySQLParserSPATIAL_SYMBOL:
	default:
		return
	}

	indexType := "BTREE"
	if ctx.IndexNameAndType() != nil && ctx.IndexNameAndType().IndexType() != nil {
		indexType = ctx.IndexNameAndType().IndexType().GetText()
	} else {
		if ctx.FULLTEXT_SYMBOL() != nil {
			indexType = "FULLTEXT"
		} else if ctx.SPATIAL_SYMBOL() != nil {
			indexType = "SPATIAL"
		}
	}
	checker.validateIndexType(indexType, ctx.GetStart().GetLine())
}

func (checker *indexTypeAllowListChecker) EnterCreateIndex(ctx *mysql.CreateIndexContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.CreateIndexTarget() == nil || ctx.CreateIndexTarget().TableRef() == nil || ctx.CreateIndexTarget().KeyListVariants() == nil {
		return
	}

	indexType := "BTREE"
	if ctx.IndexNameAndType() != nil && ctx.IndexNameAndType().IndexType() != nil {
		indexType = ctx.IndexNameAndType().IndexType().GetText()
	} else {
		if ctx.FULLTEXT_SYMBOL() != nil {
			indexType = "FULLTEXT"
		} else if ctx.SPATIAL_SYMBOL() != nil {
			indexType = "SPATIAL"
		}
	}
	checker.validateIndexType(indexType, ctx.GetStart().GetLine())
}

// validateIndexType checks if the index type is in the allow list.
func (checker *indexTypeAllowListChecker) validateIndexType(indexType string, line int) {
	if slices.Contains(checker.allowList, indexType) {
		return
	}

	checker.adviceList = append(checker.adviceList, &storepb.Advice{
		Status:        checker.level,
		Code:          advisor.IndexTypeNotAllowed.Int32(),
		Title:         checker.title,
		Content:       fmt.Sprintf("Index type `%s` is not allowed", indexType),
		StartPosition: common.ConvertANTLRLineToPosition(checker.baseLine + line),
	})
}
