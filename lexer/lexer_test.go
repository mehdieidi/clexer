package lexer

import (
	"testing"

	"github.com/MehdiEidi/clexer/token"
)

func TestNextToken(t *testing.T) {
	input := `int main() {
		int x = 4;

	}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedRow     uint
		expectedCol     uint
		expectedBlockNo uint
	}{
		{token.INT, "int", 0, 0, 0},
		{token.IDENT, "main", 0, 4, 0},
		{token.LPAREN, "(", 0, 8, 0},
		{token.RPAREN, ")", 0, 9, 0},
		{token.LBRACE, "{", 0, 11, 0},
		{token.INT, "int", 1, 2, 1},
		{token.IDENT, "x", 1, 6, 1},
		{token.ASSIGN, "=", 1, 8, 1},
		{token.INTEGER, "4", 1, 10, 1},
		{token.SEMICOLON, ";", 1, 11, 1},
		{token.RBRACE, "}", 3, 0, 0},
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

		// if tok.Col != tt.expectedCol {
		// 	t.Fatalf("tests[%d] - col number wrong. expected=%d, got=%d", i, tt.expectedCol, tok.Col)
		// }

		// if tok.BlockNo != tt.expectedBlockNo {
		// 	t.Fatalf("tests[%d] - block number wrong. expected=%d, got=%d", i, tt.expectedBlockNo, tok.BlockNo)
		// }
	}
}
