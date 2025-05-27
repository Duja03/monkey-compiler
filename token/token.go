package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals:
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators:
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters:
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords:
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
