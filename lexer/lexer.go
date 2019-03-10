package lexer

import "github.com/XCapsule/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) (lexer *Lexer) {
	lexer = &Lexer{
		input: input,
	}
	lexer.readChar()
	return
}

func (l *Lexer) NextToken() (t token.Token) {
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		t = l.processEQ()
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case '!':
		t = l.processNEQ()
	case ',':
		t = newToken(token.COMMA, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case 0:
		t = newToken(token.EOF, '0')
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdent()
			t.LookUpIdent()
			return
		} else if isDigit(l.ch) {
			t.Literal = l.readNum()
			t.Type = token.INT
			return
		}
	}
	if t.Type == "" {
		t = token.Token{
			Type:    token.ILLEGAL,
			Literal: "",
		}
	}
	l.readChar()
	return
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
		return
	}
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdent() string {
	p := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[p:l.position]
}

func (l *Lexer) readNum() string {
	p := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[p:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\r' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
	return
}

func (l *Lexer) processEQ() token.Token {
	if l.peekChar() == '=' {
		l.readChar()
		return token.Token{
			Type:    token.EQ,
			Literal: "==",
		}
	}
	return token.Token{
		Type:    token.ASSIGN,
		Literal: "=",
	}
}

func (l *Lexer) processNEQ() token.Token {
	if l.peekChar() == '=' {
		l.readChar()
		return token.Token{
			Type:    token.NEQ,
			Literal: "!=",
		}
	}
	return token.Token{
		Type:    token.BANG,
		Literal: "!",
	}
}

func (l *Lexer) peekChar() (c byte) {
	if l.readPosition >= len(l.input) {
		c = 0
		return
	}
	c = l.input[l.readPosition]
	return
}

func newToken(Type token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    Type,
		Literal: string(literal),
	}
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
