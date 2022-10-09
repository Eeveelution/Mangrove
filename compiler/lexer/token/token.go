package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	BYTE       = "BYTE"
	DOUBLE     = "DOUBLE"
	FLOAT      = "FLOAT"

	//Operators
	ASSIGN         = "="
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"
	LESS_THAN      = "<"
	GREATER_THAN   = ">"
	BINARY_OR      = "|"
	BINARY_XOR     = "^"
	BINARY_AND     = "&"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPARENTHESIS = "("
	RPARENTHESIS = ")"
	LBRACE       = "{"
	RBRACE       = "}"
	LBRACKET     = "["
	RBRACKET     = "]"
	QOUTE        = "\""

	/*     Keywords     */
	FUNCTION = "FUNC"
	VARIABLE = "VAR"
	RETURN   = "RETURN"

	//Types
	TYPE_VOID   = "VOID"
	TYPE_UINT32 = "UINT32"
	TYPE_STR    = "STR"
	TYPE_FLOAT  = "TDOUBLE"
	TYPE_DOUBLE = "TFLOAT"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"var":    VARIABLE,
	"return": RETURN,
	"void":   TYPE_VOID,
	"uint32": TYPE_UINT32,
	"str":    TYPE_STR,
	"float":  TYPE_FLOAT,
	"double": TYPE_DOUBLE,
}

func LookupIdentifier(identifier string) TokenType {
	tokenType, exists := keywords[identifier]

	if exists {
		return tokenType
	}

	return IDENTIFIER
}
