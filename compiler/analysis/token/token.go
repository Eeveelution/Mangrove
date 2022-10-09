package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string

	Line   int
	Column int
}

const (
	UNSET   = "UNSET"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	BYTE       = "BYTE"
	DOUBLE     = "DOUBLE"
	FLOAT      = "FLOAT"
	STRING     = "STRING"

	//Operators
	ASSIGN       = "="
	PLUS         = "+"
	MINUS        = "-"
	ASTERISK     = "*"
	SLASH        = "/"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	PIPE         = "|"
	CARET        = "^"
	AMPERSAND    = "&"
	PERCENT      = "%"
	BANG         = "!"

	//Multichar Operators
	EQUALITY         = "=="
	INEQUALITY       = "!="
	LOGICAL_AND      = "&&"
	LOGICAL_OR       = "||"
	LESS_OR_EQUAL    = "<="
	GREATER_OR_EQUAL = ">="

	PLUS_ASSIGN           = "+="
	MINUS_ASSIGN          = "-="
	MULTIPLICATION_ASSIGN = "*="
	DIVISION_ASSIGN       = "/="
	MOD_ASSIGN            = "%="
	BSL_ASSIGN            = "<<="
	BSR_ASSIGN            = ">>="
	AND_ASSIGN            = "&="
	XOR_ASSIGN            = "^="
	OR_ASSIGN             = "|="

	PIPE_SINGLE   = "|>"
	PIPE_MULTIPLE = "|>="

	//Delimiters
	COMMA                = ","
	SEMICOLON            = ";"
	COLON                = ":"
	NAMESPACE            = "::"
	EMPTY_PARAMETER_LIST = "<>"

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
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	WHILE    = "WHILE"
	FOR      = "FOR"
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	DEFAULT  = "DEFAULT"

	//Types
	TYPE_VOID   = "T_VOID"
	TYPE_UINT32 = "T_UINT32"
	TYPE_STR    = "T_STR"
	TYPE_FLOAT  = "T_TDOUBLE"
	TYPE_DOUBLE = "T_FLOAT"

	//Others

	QUESTIONMARK = "?"
	DOLLAR_SIGN  = "$"
	HASH_SIGN    = "#"
)

var keywords = map[string]TokenType{
	"func":    FUNCTION,
	"var":     VARIABLE,
	"return":  RETURN,
	"void":    TYPE_VOID,
	"uint32":  TYPE_UINT32,
	"str":     TYPE_STR,
	"float":   TYPE_FLOAT,
	"double":  TYPE_DOUBLE,
	"if":      TYPE_DOUBLE,
	"else":    TYPE_DOUBLE,
	"true":    TYPE_DOUBLE,
	"false":   TYPE_DOUBLE,
	"while":   TYPE_DOUBLE,
	"for":     TYPE_DOUBLE,
	"switch":  TYPE_DOUBLE,
	"case":    TYPE_DOUBLE,
	"default": TYPE_DOUBLE,
}

func LookupIdentifier(identifier string) TokenType {
	tokenType, exists := keywords[identifier]

	if exists {
		return tokenType
	}

	return IDENTIFIER
}
