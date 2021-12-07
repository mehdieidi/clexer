// Package lexer implements a lexer for c language.
package lexer

import (
	"github.com/MehdiEidi/clexer/token"
)

var (
	row     uint = 1
	col     uint
	blockNo uint
)

type Lexer struct {
	input        string // the src code
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

// New gets a src code, generates and returns a lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads one char and procedes the positions.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII code for NULL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns next token in the src code using it's lexer l.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '-':
		col++
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.EXCLA, l.ch)
		}
	case '/':
		col++
		tok = newToken(token.SLASH, l.ch)
	case '*':
		col++
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		col++
		tok = newToken(token.ST, l.ch)
	case '>':
		col++
		tok = newToken(token.GT, l.ch)
	case '=':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		col++
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		col++
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		col++
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		col++
		tok = newToken(token.COMMA, l.ch)
	case '+':
		col++
		tok = newToken(token.PLUS, l.ch)
	case '{':
		col++
		blockNo++
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		blockNo--
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			tok.Col = col
			tok.Row = row
			tok.BlockNo = blockNo

			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INTEGER
			tok.Literal = l.readNumber()

			tok.Col = col
			tok.Row = row
			tok.BlockNo = blockNo

			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	tok.Col = col
	tok.Row = row
	tok.BlockNo = blockNo

	l.readChar()
	return tok
}

// skipWhiteSpace skips spaces, tabs, linefeeds, and also updates col and row numbers.
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		switch l.ch {
		case ' ':
			col++
		case '\t':
			col += 4
		case '\n':
			col = 0
			row++
		}
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		col++
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		col++
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
