package lexer

import (
	"testing"
	// "rodrigorvsn/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.assign, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		token := lexer.NextToken()
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong, expected=%q, received=%q", i, tt.expectedType, token.Type)
		}
		if token.Literal != tt.expectedLiteral {
			t.Fatalf(("tests[%d] - literal wrong expected=%q, received=%q"), i, tt.expectedLiteral, token.Literal)
		}
	}
}
