// Code generated from Query.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterEmbedExpression is called when production embedExpression is entered.
func (s *BaseQueryListener) EnterEmbedExpression(ctx *EmbedExpressionContext) {}

// ExitEmbedExpression is called when production embedExpression is exited.
func (s *BaseQueryListener) ExitEmbedExpression(ctx *EmbedExpressionContext) {}

// EnterBoolExpression is called when production boolExpression is entered.
func (s *BaseQueryListener) EnterBoolExpression(ctx *BoolExpressionContext) {}

// ExitBoolExpression is called when production boolExpression is exited.
func (s *BaseQueryListener) ExitBoolExpression(ctx *BoolExpressionContext) {}

// EnterJoinExpression is called when production joinExpression is entered.
func (s *BaseQueryListener) EnterJoinExpression(ctx *JoinExpressionContext) {}

// ExitJoinExpression is called when production joinExpression is exited.
func (s *BaseQueryListener) ExitJoinExpression(ctx *JoinExpressionContext) {}

// EnterExpressionItem is called when production expressionItem is entered.
func (s *BaseQueryListener) EnterExpressionItem(ctx *ExpressionItemContext) {}

// ExitExpressionItem is called when production expressionItem is exited.
func (s *BaseQueryListener) ExitExpressionItem(ctx *ExpressionItemContext) {}

// EnterMatchValue is called when production matchValue is entered.
func (s *BaseQueryListener) EnterMatchValue(ctx *MatchValueContext) {}

// ExitMatchValue is called when production matchValue is exited.
func (s *BaseQueryListener) ExitMatchValue(ctx *MatchValueContext) {}

// EnterJoin is called when production join is entered.
func (s *BaseQueryListener) EnterJoin(ctx *JoinContext) {}

// ExitJoin is called when production join is exited.
func (s *BaseQueryListener) ExitJoin(ctx *JoinContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseQueryListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseQueryListener) ExitOperator(ctx *OperatorContext) {}
