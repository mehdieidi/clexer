// Package lexer implements a lexer for C language.
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

// NextToken returns next token in the src code using it's lexer l.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '-':
		if l.peekChar() == '-' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DEC, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUSAS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '>' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ARROW, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.MINUS, l.ch)
		}

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
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.DIVAS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '/' {
			l.readChar()
			col++
			l.readChar()
			l.lineComment(&tok)
			return tok
		} else if l.peekChar() == '*' {
			l.readChar()
			col++
			l.readChar()
			l.blockComment(&tok)
			return tok
		} else {
			col++
			tok = newToken(token.SLASH, l.ch)
		}

	case '*':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUSAS, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.ASTERISK, l.ch)
		}

	case '<':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.STEQ, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '<' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LSHIFT, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.ST, l.ch)
		}

	case '>':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTEQ, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '>' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.RSHIFT, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.GT, l.ch)
		}

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

	case '"':
		if isLetter(l.peekChar()) || isEscapeSequence(l.peekChar()) {
			txt := l.text()
			tok = token.Token{Type: token.TEXT, Literal: txt}
			col += 2
		}

	case '[':
		col++
		tok = newToken(token.LBRACK, l.ch)

	case ']':
		col++
		tok = newToken(token.RBRACK, l.ch)

	case ':':
		col++
		tok = newToken(token.COLON, l.ch)

	case '?':
		col++
		tok = newToken(token.QUESTION, l.ch)

	case ';':
		col++
		tok = newToken(token.SEMICOLON, l.ch)

	case '%':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PERCENTAS, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.PERCENT, l.ch)
		}

	case '&':
		if l.peekChar() == '&' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AMPERAS, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.AMPER, l.ch)
		}

	case '^':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.XORAS, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.XOR, l.ch)
		}

	case '|':
		if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PIPEAS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '|' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.PIPE, l.ch)
		}

	case '~':
		col++
		tok = newToken(token.TILDA, l.ch)

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
		if l.peekChar() == '+' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.INC, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			col += 2
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUSAS, Literal: string(ch) + string(l.ch)}
		} else {
			col++
			tok = newToken(token.PLUS, l.ch)
		}

	case '{':
		col++
		blockNo++
		tok = newToken(token.LBRACE, l.ch)

	case '}':
		col++
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
			l.readNum(&tok)
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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
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

func (l *Lexer) text() string {
	l.readChar()
	position := l.position
	for l.ch != '"' {
		col++
		l.readChar()
	}
	return l.input[position:l.position]
}

func isEscapeSequence(ch byte) bool {
	if ch == '\n' || ch == '\r' || ch == '\a' || ch == '\b' || ch == '\f' || ch == '\t' || ch == '\v' || ch == '\\' || ch == '\'' || ch == '\000' {
		return true
	}
	return false
}

func (l *Lexer) lineComment(tok *token.Token) {
	position := l.position
	for l.ch != '\n' {
		col++
		l.readChar()
	}

	if l.ch == '\r' {
		l.readChar()
	}

	*tok = token.Token{Type: token.COMMENT, Literal: l.input[position:l.position], Row: row, Col: col}

	row++
	col = 0

	l.readChar()
}

func (l *Lexer) blockComment(tok *token.Token) {
	position := l.position
	for !(l.ch == '*' && l.peekChar() == '/') {
		if l.ch == '\n' {
			row++
			col = 0
			l.readChar()
		} else if l.ch == '\r' {
			l.readChar()
		} else {
			col++
			l.readChar()
		}
	}

	*tok = token.Token{Type: token.COMMENT, Literal: l.input[position:l.position], Row: row, Col: col}

	l.readChar()

	if l.ch == '\r' {
		l.readChar()
	}

	col = 0

	l.readChar()
}

func (l *Lexer) readNum(tok *token.Token) {
	position := l.position
	for isDigit(l.ch) {
		col++
		l.readChar()
	}

	intPart := l.input[position:l.position]

	if l.ch == '.' {
		l.readChar()
		col++
		position = l.position
		for isDigit(l.ch) {
			col++
			l.readChar()
		}
		decPart := l.input[position:l.position]

		*tok = token.Token{Type: token.DECIMAL, Literal: string(intPart) + "." + string(decPart), Row: row, Col: col, BlockNo: blockNo}

		return
	}

	*tok = token.Token{Type: token.INTEGER, Literal: string(intPart), Row: row, Col: col, BlockNo: blockNo}
}
