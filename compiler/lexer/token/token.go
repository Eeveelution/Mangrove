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

	//Operators
	ASSIGN         = "="
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPARENTHESIS = "("
	RPARENTHESIS = ")"
	LBRACE       = "{"
	RBRACE       = "}"

	//Keywords
	FUNCTION = "FUNC"
	VARIABLE = "VAR"
)
