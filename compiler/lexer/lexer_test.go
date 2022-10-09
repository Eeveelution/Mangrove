package lexer

import (
	"mangrove/compiler/lexer/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `

func main<args: str[]>: uint32 {
	var testFunction: <>: uint32 {
		return 3;
	}

	var returnDouble: <num: uint32>: uint32 {
		return num * 2;
	}

	var finalNumber: uint32;

	testFunction() |> returnDouble |> finalNumber;

	return finalNumber;
}	

	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "func"},
		{token.IDENTIFIER, "main"},
		{token.LESS_THAN, "<"},
		{token.IDENTIFIER, "args"},
		{token.COLON, ":"},
		{token.TYPE_STR, "str"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.GREATER_THAN, ">"},
		{token.COLON, ":"}, //10
		{token.TYPE_UINT32, "uint32"},
		{token.LBRACE, "{"},
		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "testFunction"},
		{token.COLON, ":"},
		{token.LESS_THAN, "<"},
		{token.GREATER_THAN, ">"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.LBRACE, "{"}, //20
		{token.RETURN, "return"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "returnDouble"},
		{token.COLON, ":"},
		{token.LESS_THAN, "<"},
		{token.IDENTIFIER, "num"},
		{token.COLON, ":"}, //30
		{token.TYPE_UINT32, "uint32"},
		{token.GREATER_THAN, ">"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "num"},
		{token.MULTIPLICATION, "*"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"}, //40
		{token.RBRACE, "}"},
		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "finalNumber"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "testFunction"},
		{token.LPARENTHESIS, "("},
		{token.RPARENTHESIS, ")"},
		{token.BINARY_OR, "|"}, //50
		{token.GREATER_THAN, ">"},
		{token.IDENTIFIER, "returnDouble"},
		{token.BINARY_OR, "|"},
		{token.GREATER_THAN, ">"},
		{token.IDENTIFIER, "finalNumber"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "finalNumber"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"}, //60
	}

	lexer := CreateNewLexer(input)

	for i, test := range tests {
		if i == 21 {
			i := 0
			i = i
		}

		tok := lexer.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tok.Type wrong. Expected: %q, Got: %q", i, test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - tok.Literal wrong. Expected: %q, Got: %q", i, test.expectedLiteral, tok.Literal)
		}
	}
}
