package parser

import (
	"mangrove/compiler/analysis"
	"testing"
)

func TestParser(t *testing.T) {
	input := `

var testVariable: float = 40f;
var testVariable2: float;

`

	lexer := analysis.CreateNewLexer(input)
	parser := CreateNewParser(lexer)

	tests := []struct {
		expectedType   ExpressionType
		expectedValues interface{}
	}{
		{VARIABLE_DECLARATION, &VariableDeclarationExpression{}},
	}

	for i, test := range tests {
		if i == 31 {
			i := 06
			i = i
		}

		expression := parser.NextExpression()

		if expression.GetType() != test.expectedType {
			t.Fatalf("tests[%d] - tok.Type wrong. Expected: %q, Got: %q", i, test.expectedType, expression.GetType())
		}

		if expression != test.expectedValues {
			t.Fatalf("tests[%d] - tok.Literal wrong.", i)
		}
	}
}
