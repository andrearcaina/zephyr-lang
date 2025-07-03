package lexer

import (
	"testing"

	"github.com/andrearcaina/zephyr-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;*;`

	tests := []struct {
		expectedType    token.TokenCategory
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ASTERISK, "*"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, expectedToken := range tests {
		nextToken := l.NextToken()

		if nextToken.Type != expectedToken.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expectedToken.expectedType, nextToken.Type)
		}

		if nextToken.Literal != expectedToken.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expectedToken.expectedLiteral, nextToken.Literal)
		}
	}
}
