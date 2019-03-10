package lexer

import (
	"github.com/XCapsule/monkey-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "let a = 1 + 3;"
	expectedTokens := []token.Token{
		{
			Type:    token.LET,
			Literal: "let",
		},
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

func TestLexer_NextToken(t *testing.T) {
	input :=
		`let five = 5;
         let ten = 10;
         
		 let add = func(x,y) {
				x+y; };
		 let res = add(five,ten);
		 !-/*5;
	     5 < 10 > 5;
		 if (5 < 10) {
		     return true;
		 } else {
		     return false;
		 }
		 10 == 10;
         10 != 9;

`
	expectedTokens := []token.Token{
		{
			Type:    token.LET,
			Literal: "let",
		}, {
			Type:    token.IDENT,
			Literal: "five",
		}, {
			Type:    token.ASSIGN,
			Literal: "=",
		}, {
			Type:    token.INT,
			Literal: "5",
		}, {
			Type:    token.SEMICOLON,
			Literal: ";",
		}, {
			Type:    token.LET,
			Literal: "let",
		}, {
			Type:    token.IDENT,
			Literal: "ten",
		}, {
			Type:    token.ASSIGN,
			Literal: "=",
		}, {
			Type:    token.INT,
			Literal: "10",
		}, {
			Type:    token.SEMICOLON,
			Literal: ";",
		}, {
			Type:    token.LET,
			Literal: "let",
		}, {
			Type:    token.IDENT,
			Literal: "add",
		},
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.FUNCTION,
			Literal: "func",
		}, {
			Type:    token.LPAREN,
			Literal: "(",
		}, {
			Type:    token.IDENT,
			Literal: "x",
		}, {
			Type:    token.COMMA,
			Literal: ",",
		}, {
			Type:    token.IDENT,
			Literal: "y",
		}, {
			Type:    token.RPAREN,
			Literal: ")",
		}, {
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.IDENT,
			Literal: "x",
		}, {
			Type:    token.PLUS,
			Literal: "+",
		},
		{
			Type:    token.IDENT,
			Literal: "y"},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.LET,
			Literal: "let",
		},
		{
			Type:    token.IDENT,
			Literal: "res",
		},
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.IDENT,
			Literal: "add",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "five",
		},
		{
			Type:    token.COMMA,
			Literal: ",",
		},
		{
			Type:    token.IDENT,
			Literal: "ten",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{

			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.BANG,
			Literal: "!",
		},
		{
			Type:    token.MINUS,
			Literal: "-",
		}, {
			Type:    token.SLASH,
			Literal: "/",
		}, {
			Type:    token.ASTERISK,
			Literal: "*",
		}, {
			Type:    token.INT,
			Literal: "5",
		}, {
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.INT,
			Literal: "5",
		},
		{
			Type:    token.LT,
			Literal: "<",
		},
		{
			Type:    token.INT,
			Literal: "10",
		},
		{
			Type:    token.GT,
			Literal: ">",
		},
		{
			Type:    token.INT,
			Literal: "5",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{Type: token.IF, Literal: "if"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.INT, Literal: "5"},
		{Type: token.LT, Literal: "<"},
		{Type: token.INT, Literal: "10"},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FALSE, Literal: "false"},
		{Type: token.SEMICOLON, Literal: ";"},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.INT,
			Literal: "10",
		},
		{
			Type:    token.EQ,
			Literal: "==",
		},
		{
			Type:    token.INT,
			Literal: "10",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		}, {
			Type:    token.INT,
			Literal: "10",
		}, {
			Type:    token.NEQ,
			Literal: "!=",
		},
		{
			Type:    token.INT,
			Literal: "9",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.EOF,
			Literal: "0",
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
