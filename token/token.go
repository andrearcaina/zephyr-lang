package token

// Token represents a single token in the input given to the lexer

type TokenCategory string

type Token struct {
	Type    TokenCategory
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	FUNCTION = "FUNC"
	LET      = "LET"
	CONST    = "CONST"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenCategory{
	"func":   FUNCTION,
	"let":    LET,
	"const":  CONST,
	"return": RETURN,
}

// LookupIdent checks if the given identifier is a keyword and returns the corresponding token type
func LookupIdentifier(identifier string) TokenCategory {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
