package lexer

import (
	"os"
	"testing"

	"github.com/MehdiEidi/clexer/token"
)

func TestNextToken(t *testing.T) {
	src, _ := os.ReadFile("../input/source.c")
	input := string(src)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedRow     uint
		expectedCol     uint
		expectedBlockNo uint
	}{
		{token.INT, "int", 1, 3, 0},
		{token.IDENT, "main", 1, 8, 0},
		{token.LPAREN, "(", 1, 9, 0},
		{token.RPAREN, ")", 1, 10, 0},
		{token.LBRACE, "{", 1, 12, 1},
		{token.INT, "int", 2, 7, 1},
		{token.IDENT, "x", 2, 9, 1},
		{token.ASSIGN, "=", 2, 11, 1},
		{token.INTEGER, "4", 2, 13, 1},
		{token.SEMICOLON, ";", 2, 14, 1},
		{token.RBRACE, "}", 3, 1, 0},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Row != tt.expectedRow {
			t.Fatalf("tests[%d] - row number wrong. expected=%d, got=%d", i, tt.expectedRow, tok.Row)
		}

		if tok.Col != tt.expectedCol {
			t.Fatalf("tests[%d] - col number wrong. expected=%d, got=%d", i, tt.expectedCol, tok.Col)
		}

		if tok.BlockNo != tt.expectedBlockNo {
			t.Fatalf("tests[%d] - block number wrong. expected=%d, got=%d", i, tt.expectedBlockNo, tok.BlockNo)
		}
	}
}
