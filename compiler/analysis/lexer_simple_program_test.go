package analysis

import (
	"mangrove/compiler/analysis/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `

func main<args: str[]>: uint32 {
	var testFloat: float = 44.561f;
	var testDouble: double = 88.121;
	var testString: str = "Hello, World!";

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
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.LBRACE, "{"},

		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "testFloat"},
		{token.COLON, ":"},
		{token.TYPE_FLOAT, "float"},
		{token.ASSIGN, "="},
		{token.FLOAT, "44.561f"},
		{token.SEMICOLON, ";"},

		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "testDouble"},
		{token.COLON, ":"},
		{token.TYPE_DOUBLE, "double"},
		{token.ASSIGN, "="},
		{token.DOUBLE, "88.121"},
		{token.SEMICOLON, ";"},

		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "testString"},
		{token.COLON, ":"},
		{token.TYPE_STR, "str"},
		{token.ASSIGN, "="},
		{token.STRING, "\"Hello, World!\""},
		{token.SEMICOLON, ";"},

		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "testFunction"},
		{token.COLON, ":"},
		{token.EMPTY_PARAMETER_LIST, "<>"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},

		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "returnDouble"},
		{token.COLON, ":"},
		{token.LESS_THAN, "<"},
		{token.IDENTIFIER, "num"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.GREATER_THAN, ">"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.LBRACE, "{"},

		{token.RETURN, "return"},
		{token.IDENTIFIER, "num"},
		{token.ASTERISK, "*"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},

		{token.VARIABLE, "var"},
		{token.IDENTIFIER, "finalNumber"},
		{token.COLON, ":"},
		{token.TYPE_UINT32, "uint32"},
		{token.SEMICOLON, ";"},

		{token.IDENTIFIER, "testFunction"},
		{token.LPARENTHESIS, "("},
		{token.RPARENTHESIS, ")"},
		{token.PIPE_SINGLE, "|>"},
		{token.IDENTIFIER, "returnDouble"},
		{token.PIPE_SINGLE, "|>"},
		{token.IDENTIFIER, "finalNumber"},
		{token.SEMICOLON, ";"},

		{token.RETURN, "return"},
		{token.IDENTIFIER, "finalNumber"},
		{token.SEMICOLON, ";"},

		{token.RBRACE, "}"},
	}

	lexer := CreateNewLexer(input)

	for i, test := range tests {
		if i == 31 {
			i := 06
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
