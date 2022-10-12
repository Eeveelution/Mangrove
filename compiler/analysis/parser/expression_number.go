package parser

import (
	"mangrove/compiler/analysis/token"
)

type NumberExpression struct {
	Token token.Token
}

func (expression *NumberExpression) GetType() ExpressionType {
	return NUMBER
}

