// Package lexer implements a lexer for C language.
package lexer

import (
	"github.com/MehdiEidi/clexer/token"
)

type Lexer struct {
	input        string // the src code
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
	row          uint
	col          uint
	blockNo      uint
}

// NextToken returns next token in the src code using it's lexer l.
func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhiteSpace()
	l.skipComments()
	l.skipPreprocessorDirectives()

	switch l.ch {
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.DEC, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.MINUSAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.ARROW, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.MINUS, string(l.ch), 0, 0, 0)
		}

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.NOT_EQ, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.EXCLA, string(l.ch), 0, 0, 0)
		}

	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.DIVAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.SLASH, string(l.ch), 0, 0, 0)
		}

	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.MULTAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.ASTERISK, string(l.ch), 0, 0, 0)
		}

	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.STEQ, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.LSHIFT, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.ST, string(l.ch), 0, 0, 0)
		}

	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.GTEQ, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.RSHIFT, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.GT, string(l.ch), 0, 0, 0)
		}

	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.EQ, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.ASSIGN, string(l.ch), 0, 0, 0)
		}

	case '"':
		curCol := l.col
		txt := l.readText()
		tok = token.New(token.TEXT, txt, l.row, curCol, l.blockNo)
		l.readChar()
		return tok

	case '.':
		tok = token.New(token.DOT, string(l.ch), 0, 0, 0)

	case '[':
		tok = token.New(token.LBRACK, string(l.ch), 0, 0, 0)

	case ']':
		tok = token.New(token.RBRACK, string(l.ch), 0, 0, 0)

	case ':':
		tok = token.New(token.COLON, string(l.ch), 0, 0, 0)

	case '?':
		tok = token.New(token.QUESTION, string(l.ch), 0, 0, 0)

	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch), 0, 0, 0)

	case '%':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.PERCENTAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.PERCENT, string(l.ch), 0, 0, 0)
		}

	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.AND, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.AMPERAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.AMPER, string(l.ch), 0, 0, 0)
		}

	case '^':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.XORAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.XOR, string(l.ch), 0, 0, 0)
		}

	case '|':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.PIPEAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.OR, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.PIPE, string(l.ch), 0, 0, 0)
		}

	case '~':
		tok = token.New(token.TILDA, string(l.ch), 0, 0, 0)

	case '(':
		tok = token.New(token.LPAREN, string(l.ch), 0, 0, 0)

	case ')':
		tok = token.New(token.RPAREN, string(l.ch), 0, 0, 0)

	case ',':
		tok = token.New(token.COMMA, string(l.ch), 0, 0, 0)

	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.INC, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.PLUSAS, string(ch)+string(l.ch), l.row, l.col-1, l.blockNo)
			l.readChar()
			return tok
		} else {
			tok = token.New(token.PLUS, string(l.ch), 0, 0, 0)
		}

	case '{':
		l.blockNo++
		tok = token.New(token.LBRACE, string(l.ch), 0, 0, 0)

	case '}':
		l.blockNo--
		tok = token.New(token.RBRACE, string(l.ch), 0, 0, 0)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			curCol := l.col

			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			tok.Col = curCol
			tok.Row = l.row
			tok.BlockNo = l.blockNo

			return tok
		} else if isDigit(l.ch) {
			curCol := l.col
			tok = l.readNum()
			tok.Col = curCol
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch), 0, 0, 0)
		}
	}

	tok.Col = l.col
	tok.Row = l.row
	tok.BlockNo = l.blockNo

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
	l.col++
}

// skipWhiteSpace skips whitespaces also updates col and row numbers.
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		switch l.ch {
		case '\t':
			l.col++
		case '\n':
			l.readChar()
			l.col--
			l.row++
			continue
		case '\r':
			l.readChar()
			l.col = 1
			continue
		}
		l.readChar()
	}
}

// readIdentifier reads and returns a whole word.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// peekChar just returns the next char to be read.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readText returns the string between ""
func (l *Lexer) readText() string {
	l.readChar() // skip starting "
	position := l.position
	for l.ch != '"' {
		// EOF
		if l.ch == 0 {
			break
		}
		l.readChar()
		l.skipWhiteSpace()
	}
	return l.input[position:l.position]
}

// readNum reads a number either decimal or integer and returns the appropriate token.
func (l *Lexer) readNum() (tok token.Token) {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	intPart := l.input[position:l.position]

	if l.ch == '.' {
		l.readChar()

		position = l.position
		for isDigit(l.ch) {
			l.readChar()
		}

		decPart := l.input[position:l.position]

		tok = token.New(token.DECIMAL, intPart+"."+decPart, l.row, l.col, l.blockNo)
		return
	}

	tok = token.New(token.INTEGER, intPart, l.row, l.col, l.blockNo)
	return
}

// skipLineComment just ignores a line comment.
func (l *Lexer) skipLineComment() {
	for l.ch != '\r' {
		l.readChar()
	}

	l.readChar()

	if l.ch == '\n' {
		l.row++
		l.readChar()
	}
	l.col = 1
}

// skipBlockComment just ignores a block comment.
func (l *Lexer) skipBlockComment() {
	for !(l.ch == '*' && l.peekChar() == '/') {
		if l.ch == '\n' {
			l.row++
			l.readChar()
			l.col--
		} else if l.ch == '\r' {
			l.readChar()
			l.col = 1
		} else {
			l.readChar()
		}
	}

	l.readChar() // skip *
	l.readChar() // skip /

	l.skipWhiteSpace()
}

// skipComments just ignores any sort of comment.
func (l *Lexer) skipComments() {
	for l.ch == '/' && (l.peekChar() == '/' || l.peekChar() == '*') {
		if l.ch == '/' && l.peekChar() == '/' {
			// skip slash slash
			l.readChar()
			l.readChar()
			l.skipLineComment()
		}
		if l.ch == '/' && l.peekChar() == '*' {
			l.skipBlockComment()
		}
		l.skipWhiteSpace()
	}
}

func (l *Lexer) skipPreprocessorDirectives() {
	for l.ch == '#' {
		for l.ch != '\n' {
			l.readChar()
		}
		l.skipWhiteSpace()
	}
}

// New gets a src code, generates and returns a lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input, col: 1, row: 1}
	l.readChar() // get to the first char
	l.col--
	return l
}

// isLetter returns true if ch is a letter. _ is considered as a letter too.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit returns true if ch is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
