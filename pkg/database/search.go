package database

import (
	"strings"

	"github.com/AdeAttwood/SoleMail/pkg/database/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
)

func tagComp(opp, value string, thread Thread) (bool, error) {
	switch opp {
	case "=":
		return thread.HasTag(value), nil
	case "!":
		return !thread.HasTag(value), nil
	}

	return false, errors.New("Only operators '=' and '!' as supported for field 'tag'")
}

func fromComp(opp, value string, thread Thread) (bool, error) {
	switch opp {
	case "=":
		return thread.From == value, nil
	case "!":
		return thread.From != value, nil
	}

	return false, errors.New("Only operators '=' and '!' as supported for field 'from'")
}

func threadComp(field, opp, value string, thread Thread) (bool, error) {
	if field == "tag" {
		return tagComp(opp, value, thread)
	}

	if field == "from" {
		return fromComp(opp, value, thread)
	}

	return false, errors.Errorf("Invalid field '%s'", field)
}

type ThreadMatchListener struct {
	*parser.BaseQueryListener
	stack  boolstack
	thread Thread
	errors []error
}

func NewThreadMatchListener(thread Thread) *ThreadMatchListener {
	return &ThreadMatchListener{
		stack:  boolstack{},
		thread: thread,
	}
}

func (this *ThreadMatchListener) ExitExpressionItem(ctx *parser.ExpressionItemContext) {
	opp := "="
	if ctx.Operator() != nil {
		opp = ctx.Operator().GetText()
	}

	field := ctx.Word().GetText()
	match := ""

	matchValue := ctx.MatchValue().(*parser.MatchValueContext)
	if matchValue.Word() != nil {
		match = matchValue.Word().GetText()
	} else if matchValue.QuotedString() != nil {
		match = strings.Trim(matchValue.QuotedString().GetText(), "\"")
	}

	result, err := threadComp(field, opp, match, this.thread)
	if err != nil {
		this.errors = append(this.errors, err)
		return
	}

	this.stack.push(result)
}

func (this *ThreadMatchListener) ExitJoinExpression(ctx *parser.JoinExpressionContext) {
	left, err := this.stack.pop()
	if err != nil {
		this.errors = append(this.errors, err)
		return
	}

	right, err := this.stack.pop()
	if err != nil {
		this.errors = append(this.errors, err)
		return
	}

	if ctx.Join() != nil && ctx.Join().GetText() == "or" {
		this.stack.push(left || right)
	} else {
		this.stack.push(left && right)
	}
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func (this *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	this.errors = append(this.errors, errors.New(msg))
}

func MatchThread(query string, thread Thread) (bool, error) {
	input := antlr.NewInputStream(query)
	error_listener := &ErrorListener{}

	lexer := parser.NewQueryLexer(input)
	lexer.AddErrorListener(error_listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	if len(error_listener.errors) > 0 {
		return false, error_listener.errors[0]
	}

	p := parser.NewQueryParser(stream)
	p.AddErrorListener(error_listener)
	p.BuildParseTrees = true

	tree := p.Query()
	if len(error_listener.errors) > 0 {
		return false, error_listener.errors[0]
	}

	l := NewThreadMatchListener(thread)
	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	if len(l.errors) > 0 {
		return false, l.errors[0]
	}

	return l.stack.pop()
}
