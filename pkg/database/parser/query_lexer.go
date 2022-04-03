// Code generated from Query.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 59, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8,
	38, 10, 8, 13, 8, 14, 8, 39, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 46, 10, 9, 12,
	9, 14, 9, 49, 11, 9, 3, 9, 3, 9, 3, 10, 6, 10, 54, 10, 10, 13, 10, 14,
	10, 55, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 3, 2, 5, 4, 2, 67, 92, 99, 124, 4, 2, 36, 36, 94, 94,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 62, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2,
	2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 32,
	3, 2, 2, 2, 13, 34, 3, 2, 2, 2, 15, 37, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2,
	19, 53, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7,
	43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 99, 2, 2, 26, 27, 7, 112, 2, 2,
	27, 28, 7, 102, 2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 113, 2, 2, 30, 31,
	7, 116, 2, 2, 31, 10, 3, 2, 2, 2, 32, 33, 7, 35, 2, 2, 33, 12, 3, 2, 2,
	2, 34, 35, 7, 60, 2, 2, 35, 14, 3, 2, 2, 2, 36, 38, 9, 2, 2, 2, 37, 36,
	3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2,
	40, 16, 3, 2, 2, 2, 41, 47, 7, 36, 2, 2, 42, 46, 10, 3, 2, 2, 43, 44, 7,
	94, 2, 2, 44, 46, 9, 3, 2, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46,
	49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2,
	2, 49, 47, 3, 2, 2, 2, 50, 51, 7, 36, 2, 2, 51, 18, 3, 2, 2, 2, 52, 54,
	9, 4, 2, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	55, 56, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 8, 10, 2, 2, 58, 20, 3,
	2, 2, 2, 7, 2, 39, 45, 47, 55, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'and'", "'or'", "'!'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "NotEqual", "Colon", "Word", "QuotedString", "SPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "NotEqual", "Colon", "Word", "QuotedString",
	"SPACE",
}

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *QueryLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewQueryLexer(input antlr.CharStream) *QueryLexer {
	l := new(QueryLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerT__0         = 1
	QueryLexerT__1         = 2
	QueryLexerT__2         = 3
	QueryLexerT__3         = 4
	QueryLexerNotEqual     = 5
	QueryLexerColon        = 6
	QueryLexerWord         = 7
	QueryLexerQuotedString = 8
	QueryLexerSPACE        = 9
)
