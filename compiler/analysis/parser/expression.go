package parser

type ExpressionType string

const (
	VARIABLE_DECLARATION ExpressionType = "VARIABLE_DECLARATION"
)

type Expression interface {
	GetType() ExpressionType
}
