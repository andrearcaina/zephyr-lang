package lexer

import (
	"github.com/andrearcaina/zephyr-lang/token"
)

// A Lexer (in the context of programming) is a component that processes input text
// and converts the text into a sequence of tokens, so it's easily readable by a parser (or interpreter)

type Lexer struct {
	input        string // the input string to be lexed
	position     int    // current position in input (points to curr char)
	readPosition int    // current reading position input (after curr char)
	char         byte   // the curr char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // 0 represents EOF
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = "" // manually set to empty string because inputting "" in newToken converts it to a string with a single character 0
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenCategory, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
