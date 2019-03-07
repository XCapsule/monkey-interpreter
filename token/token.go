package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var identsMap = map[string]TokenType{
	"func": FUNCTION,
	"let":  LET,
}

type TokenType string

type Token struct {
	Type        TokenType
	Literal     string
	Column, Row int32
}

func (t *Token) LookUpIdent() {
	t.Type = identsMap[t.Literal]
	if t.Type == "" {
		t.Type = IDENT
	}
}
