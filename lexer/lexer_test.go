package lexer

import (
	"github.com/XCapsule/monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "a=1+3;"
	expectedTokens := []token.Token{
		{
			Type:    token.IDENT,
			Literal: "a",
			//Column:  1,
			//Row:     1,
		},
		{
			Type:    token.ASSIGN,
			Literal: "=",
			//Column:  2,
			//Row:     1,
		},
		{
			Type:    token.INT,
			Literal: "1",
			//Column:  3,
			//Row:     1,
		},
		{
			Type:    token.PLUS,
			Literal: "+",
			//Column:  4,
			//Row:     1,
		},
		{
			Type:    token.INT,
			Literal: "3",
			//Column:  5,
			//Row:     1,
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
			//Column:  0,
			//Row:     0,
		},
	}
	lex := New(input)
	for _, exp := range expectedTokens {
		var tok token.Token
		tok = lex.NextToken()
		if exp.Type != tok.Type || exp.Literal != tok.Literal || exp.Column != tok.Column || exp.Row != tok.Row {
			t.Errorf("Not Equal Exp(%v),Get(%v)", exp, tok)
		}
	}
}
