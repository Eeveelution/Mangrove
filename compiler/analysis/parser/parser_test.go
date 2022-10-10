package parser

import (
	"mangrove/compiler/analysis"
	"mangrove/compiler/analysis/token"
	"testing"
)

func TestParser(t *testing.T) {
	input := `

var testVariable3: std::threading::thread;
var testVariable: float = 40f;
var testVariable2: float;

`

	lexer := analysis.CreateNewLexer(input)
	parser := CreateNewParser(lexer)

	tests := []struct {
		expectedType   ExpressionType
		expectedValues Expression
	}{
		{
			VARIABLE_DECLARATION,
			&VariableDeclarationExpression{
				Name:                "testVariable3",
				ImmediateAssignment: false,
				Type:                TYPE_COMPLEX,
				AsComplexType:       "std::threading::thread",
			},
		},
		{
			VARIABLE_DECLARATION,
			&VariableDeclarationExpression{
				Name:                "testVariable",
				ImmediateAssignment: true,
				Type:                TYPE_TRADITIONAL,
				AsTraditional: token.Token{
					Type:    token.FLOAT,
					Literal: "40f",
				},
			},
		},
		{
			VARIABLE_DECLARATION,
			&VariableDeclarationExpression{
				Name:                "testVariable2",
				Type:                TYPE_TRADITIONAL,
				ImmediateAssignment: false,
				AsTraditional: token.Token{
					Type:    token.FLOAT,
					Literal: "",
				},
			},
		},
	}

	for i, test := range tests {
		if i == 31 {
			i = i
		}

		expression := parser.NextExpression()

		switch gottenExpression := expression.(type) {
		case *VariableDeclarationExpression:
			switch typedTest := test.expectedValues.(type) {
			case *VariableDeclarationExpression:
				if gottenExpression.Name == typedTest.Name && gottenExpression.ImmediateAssignment == typedTest.ImmediateAssignment && gottenExpression.Type == typedTest.Type {
					switch typedTest.Type {
					case TYPE_TRADITIONAL:
						if gottenExpression.AsTraditional != typedTest.AsTraditional {
							t.Fatalf(
								`tests[%d] - FAIL! Value does not match!
Expected Name     : %s; Got: %s
Expected ImmAssign: %t; Got: %t
Expected Type     : %d; Got: %d
Expected Value    : %s; Got: %s
								`, i,
								typedTest.Name, gottenExpression.Name,
								typedTest.ImmediateAssignment, gottenExpression.ImmediateAssignment,
								typedTest.Type, gottenExpression.Type,
								typedTest.AsTraditional.Literal, gottenExpression.AsTraditional.Literal,
							)
						}
					case TYPE_COMPLEX:
						if gottenExpression.AsComplexValue != typedTest.AsComplexValue {
							t.Fatalf(
								`tests[%d] - FAIL! Value does not match!
Expected Name     : %s; Got: %s
Expected ImmAssign: %t; Got: %t
Expected Type     : %d; Got: %d
Expected Value    : %s; Got: %s
								`, i,
								typedTest.Name, gottenExpression.Name,
								typedTest.ImmediateAssignment, gottenExpression.ImmediateAssignment,
								typedTest.Type, gottenExpression.Type,
								typedTest.AsTraditional.Literal, gottenExpression.AsTraditional.Literal,
							)
						}
					}

					t.Logf("tests[%d] - PASS", i)
				} else {
					t.Fatalf(
						`tests[%d] - FAIL!
Expected Name     : %s; Got: %s
Expected ImmAssign: %t; Got: %t
Expected Type     : %d; Got: %d
						`, i,
						typedTest.Name, gottenExpression.Name,
						typedTest.ImmediateAssignment, gottenExpression.ImmediateAssignment,
						typedTest.Type, gottenExpression.Type,
					)
				}
			default:
				t.Fatalf("tests[%d] - tok.Type wrong. Expected: %q, Got: %q", i, test.expectedType, expression.GetType())
			}
		}
	}
}
