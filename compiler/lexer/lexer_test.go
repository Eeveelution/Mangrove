package lexer

import (
	"mangrove/compiler/lexer/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+*(){}/,;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MULTIPLICATION, "*"},
		{token.LPARENTHESIS, "("},
		{token.RPARENTHESIS, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.DIVISION, "/"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := CreateNewLexer(input)

	for i, test := range tests {
		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tok.Type wrong. Expected: %q, Got: %q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - tok.Literal wrong. Expected: %q, Got: %q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
