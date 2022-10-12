package parser

type ExpressionType string

const (
	NUMBER ExpressionType = "NUMBER"
	VARIABLE_DECLARATION ExpressionType = "VARIABLE_DECLARATION"
)

type Expression interface {
	GetType() ExpressionType
}
