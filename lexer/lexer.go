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

// New creates a new Lexer instance with the given input string and initializes the first character to be read
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input string and updates the position and readPosition
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // 0 represents EOF
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token from the input string based on the current character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

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
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads an identifier from the input string and returns it as a string
func (l *Lexer) readIdentifier() string {
	startPosition := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

// readNumber reads a number from the input string and returns it as a string
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// isLetter checks if the given byte is a letter
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

// skipWhitespace skips over any whitespace characters in the input string
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// newToken creates a new token with the specified type and character
func newToken(tokenType token.TokenCategory, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
