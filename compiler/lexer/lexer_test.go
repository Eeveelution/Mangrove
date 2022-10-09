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
}
