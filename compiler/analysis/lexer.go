package analysis

import "mangrove/compiler/analysis/token"

type Lexer struct {
	input        string //Input String
	position     int    //Position in the input (points to current char)
	readPosition int    //Current reading position in input (points to current char + 1)
	char         byte   //Current Char under examination
}

func CreateNewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	lexer.readChar()

	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.char {

	case '?':
		tok = newToken(token.QUESTIONMARK, lexer.char)
	case '$':
		tok = newToken(token.DOLLAR_SIGN, lexer.char)
	case '#':
		tok = newToken(token.HASH_SIGN, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case ':':
		tok = newToken(token.COLON, lexer.char)

	case '(':
		tok = newToken(token.LPARENTHESIS, lexer.char)
	case ')':
		tok = newToken(token.RPARENTHESIS, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case '[':
		tok = newToken(token.LBRACKET, lexer.char)
	case ']':
		tok = newToken(token.RBRACKET, lexer.char)
	case '"':
		tok = newToken(token.QOUTE, lexer.char)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if canBeInIdentifier(lexer.char, true) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)

			return tok
		} else if isDigit(lexer.char) {
			numLiteral, numType := lexer.readNumber()

			tok.Literal = numLiteral
			tok.Type = numType

			return tok
		} else if isOperatorSymbol(lexer.char) {
			operatorLiteral, operatorType := lexer.readOperator()

			tok.Literal = operatorLiteral
			tok.Type = operatorType

			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()

	return tok
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	identifierBeginning := true

	for canBeInIdentifier(lexer.char, identifierBeginning) {
		lexer.readChar()
		identifierBeginning = false
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() (literal string, tokenType token.TokenType) {
	position := lexer.position

	numberTokenType := token.INT

	for {
		if isDigit(lexer.char) {
			lexer.readChar()
		} else if lexer.char == '.' {
			numberTokenType = token.DOUBLE
			lexer.readChar()
		} else if lexer.char == 'f' {
			numberTokenType = token.FLOAT
			lexer.readChar()
			break
		} else {
			break
		}
	}

	return lexer.input[position:lexer.position], token.TokenType(numberTokenType)
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func canBeInIdentifier(char byte, identifierBeginning bool) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' ||
		//Special handling for numbers in identifiers, they cannot be at the beginning
		//to not confuse them with integers
		(!identifierBeginning && isDigit(char))
}

func isOperatorSymbol(char byte) bool {
	return char == '=' ||
		char == '+' ||
		char == '-' ||
		char == '*' ||
		char == '/' ||
		char == '<' ||
		char == '>' ||
		char == '|' ||
		char == '^' ||
		char == '&' ||
		char == '%' ||
		char == '!'
}

func (lexer *Lexer) readOperator() (literal string, tokenType token.TokenType) {
	operator := ""
	var operatorType token.TokenType = token.ILLEGAL

	for {
		if isOperatorSymbol(lexer.char) {
			operator += string(lexer.char)

			lexer.readChar()
		} else {
			break
		}
	}

	switch operator {
	case "=":
		operatorType = token.ASSIGN
	case "+":
		operatorType = token.PLUS
	case "-":
		operatorType = token.MINUS
	case "*":
		operatorType = token.ASTERISK
	case "/":
		operatorType = token.SLASH
	case "<":
		operatorType = token.LESS_THAN
	case ">":
		operatorType = token.GREATER_THAN
	case "|":
		operatorType = token.PIPE
	case "^":
		operatorType = token.CARET
	case "&":
		operatorType = token.AMPERSAND
	case "%":
		operatorType = token.PERCENT
	case "!":
		operatorType = token.BANG
	case "==":
		operatorType = token.EQUALITY
	case "!=":
		operatorType = token.INEQUALITY
	case "&&":
		operatorType = token.LOGICAL_AND
	case "||":
		operatorType = token.LOGICAL_OR
	case "<=":
		operatorType = token.LESS_OR_EQUAL
	case ">=":
		operatorType = token.GREATER_OR_EQUAL
	case "+=":
		operatorType = token.PLUS_ASSIGN
	case "-=":
		operatorType = token.MINUS_ASSIGN
	case "*=":
		operatorType = token.MULTIPLICATION_ASSIGN
	case "/=":
		operatorType = token.DIVISION_ASSIGN
	case "%=":
		operatorType = token.MULTIPLICATION_ASSIGN
	case "<<=":
		operatorType = token.BSL_ASSIGN
	case ">>=":
		operatorType = token.BSR_ASSIGN
	case "&=":
		operatorType = token.AND_ASSIGN
	case "^=":
		operatorType = token.XOR_ASSIGN
	case "|=":
		operatorType = token.OR_ASSIGN
	case "|>":
		operatorType = token.PIPE_SINGLE
	case "|>=":
		operatorType = token.PIPE_MULTIPLE
	case "<>":
		operatorType = token.EMPTY_PARAMETER_LIST
	}

	return operator, operatorType
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}
