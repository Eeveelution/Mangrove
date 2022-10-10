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
	BREAK    = "BREAK"
	DEFAULT  = "DEFAULT"

	//Types
	TYPE_VOID = "T_VOID"

	TYPE_CHAR   = "T_CHAR"
	TYPE_UINT8  = "T_UINT8"
	TYPE_UINT16 = "T_UINT16"
	TYPE_UINT32 = "T_UINT32"
	TYPE_UINT64 = "T_UINT64"
	TYPE_INT8   = "T_INT8"
	TYPE_INT16  = "T_INT16"
	TYPE_INT32  = "T_INT32"
	TYPE_INT64  = "T_INT64"

	TYPE_STR = "T_STR"

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
	"char":    TYPE_CHAR,
	"uint8":   TYPE_UINT8,
	"uint16":  TYPE_UINT16,
	"uint32":  TYPE_UINT32,
	"uint64":  TYPE_UINT64,
	"int8":    TYPE_INT8,
	"int16":   TYPE_INT16,
	"int32":   TYPE_INT32,
	"int64":   TYPE_INT64,
	"str":     TYPE_STR,
	"float":   TYPE_FLOAT,
	"double":  TYPE_DOUBLE,
	"if":      IF,
	"else":    ELSE,
	"true":    TRUE,
	"false":   FALSE,
	"while":   WHILE,
	"for":     FOR,
	"switch":  SWITCH,
	"case":    CASE,
	"break":   BREAK,
	"default": DEFAULT,
}

func LookupIdentifier(identifier string) TokenType {
	tokenType, exists := keywords[identifier]

	if exists {
		return tokenType
	}

	return IDENTIFIER
}
