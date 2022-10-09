package parser

import "go/token"

type VariableDeclarationExpression struct {
	Name                string
	ImmediateAssignment bool
	AssignedValue       token.Token
}

func (expression *VariableDeclarationExpression) GetType() ExpressionType {
	return VARIABLE_DECLARATION
}

func (parser *Parser) ParseVariableDeclarationExpression() Expression {
	
}
