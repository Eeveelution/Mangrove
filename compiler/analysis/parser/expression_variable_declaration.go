package parser

import (
	"mangrove/compiler/analysis/token"
	"mangrove/misc/logging"
)

type VariableType uint32

const (
	TYPE_TRADITIONAL VariableType = 0
	TYPE_COMPLEX     VariableType = 1
	TYPE_FUNCTION    VariableType = 2
	TYPE_TUPLE       VariableType = 3
)

type VariableDeclarationExpression struct {
	Name                string
	ImmediateAssignment bool
	AssignedValue       token.Token

	Type          VariableType
	AsTraditional token.Token
	//Type name for example std::threading::thread
	AsComplexType  string
	AsComplexValue string
}

func (expression *VariableDeclarationExpression) GetType() ExpressionType {
	return VARIABLE_DECLARATION
}

func (parser *Parser) ParseVariableDeclarationExpression() Expression {
	declExpression := VariableDeclarationExpression{}

	//Expecting Next token to be token.VARIABLE
	parser.NextToken()

	if parser.tok.Type != token.VARIABLE {
		logging.FatalUnexpectedTokenError(1, parser.tok, "Expected 'var' for the beginning of variable declaration; weird how you got here, this code usually gets run when the Parser finds a 'var' hmm...")
	}

	//Expecting Next Token to be token.IDENTIFIER
	parser.NextToken()

	if parser.tok.Type != token.IDENTIFIER {
		logging.FatalUnexpectedTokenError(1, parser.tok, "Expected variable name after `var` unnamed variables aren't allowed!")
	}

	declExpression.Name = parser.tok.Literal

	//Expecting next Token to be token.COLON
	parser.NextToken()

	if parser.tok.Type != token.COLON {
		logging.FatalUnexpectedTokenError(1, parser.tok, "Expected a type definition for variable!")
	}

	//We're now either expecting a unnamed type list or a single type token or identifier
	parser.NextToken()

	switch parser.tok.Type {
	case token.LESS_THAN:
		//Expecting a unnamed type list
		declExpression.Type = TYPE_FUNCTION
	case token.LBRACKET:
		//Expecting a tuple
		declExpression.Type = TYPE_TUPLE
	case token.IDENTIFIER:
		//Expecting a Identifier saying a more complex type, potentially with namespaces
		declExpression.Type = TYPE_COMPLEX
	case token.TYPE_FLOAT:
		declExpression.Type = TYPE_TRADITIONAL
		declExpression.AsTraditional = parser.tok
	case token.TYPE_UINT32:
		declExpression.Type = TYPE_TRADITIONAL
		declExpression.AsTraditional = parser.tok
	case token.TYPE_DOUBLE:
		declExpression.Type = TYPE_TRADITIONAL
		declExpression.AsTraditional = parser.tok
	case token.TYPE_STR:
		declExpression.Type = TYPE_TRADITIONAL
		declExpression.AsTraditional = parser.tok
	}

	if declExpression.Type == TYPE_COMPLEX {
		constructedType := ""

		for {
			constructedType += parser.tok.Literal

			parser.NextToken()

			if parser.tok.Type == token.SEMICOLON || parser.tok.Type == token.ASSIGN {
				break
			}
		}

		declExpression.AsComplexType = constructedType
	}

	if parser.tok.Type == token.ASSIGN && declExpression.Type != TYPE_FUNCTION {
		declExpression.ImmediateAssignment = true

		parser.NextToken()

		assignValue := ""

		for {
			assignValue += parser.tok.Literal

			parser.NextToken()

			if parser.tok.Type == token.SEMICOLON {
				break
			}
		}

	}

	return &declExpression
}
