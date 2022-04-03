// Code generated from Query.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterEmbedExpression is called when entering the embedExpression production.
	EnterEmbedExpression(c *EmbedExpressionContext)

	// EnterBoolExpression is called when entering the boolExpression production.
	EnterBoolExpression(c *BoolExpressionContext)

	// EnterJoinExpression is called when entering the joinExpression production.
	EnterJoinExpression(c *JoinExpressionContext)

	// EnterExpressionItem is called when entering the expressionItem production.
	EnterExpressionItem(c *ExpressionItemContext)

	// EnterMatchValue is called when entering the matchValue production.
	EnterMatchValue(c *MatchValueContext)

	// EnterJoin is called when entering the join production.
	EnterJoin(c *JoinContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitEmbedExpression is called when exiting the embedExpression production.
	ExitEmbedExpression(c *EmbedExpressionContext)

	// ExitBoolExpression is called when exiting the boolExpression production.
	ExitBoolExpression(c *BoolExpressionContext)

	// ExitJoinExpression is called when exiting the joinExpression production.
	ExitJoinExpression(c *JoinExpressionContext)

	// ExitExpressionItem is called when exiting the expressionItem production.
	ExitExpressionItem(c *ExpressionItemContext)

	// ExitMatchValue is called when exiting the matchValue production.
	ExitMatchValue(c *MatchValueContext)

	// ExitJoin is called when exiting the join production.
	ExitJoin(c *JoinContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
