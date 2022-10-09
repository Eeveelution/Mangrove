package analysis

import "mangrove/compiler/analysis/token"

func canBeInIdentifier(char byte, identifierBeginning bool) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' ||
		//Special handling for numbers in identifiers, they cannot be at the beginning
		//to not confuse them with integers
		(!identifierBeginning && isDigit(char))
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

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
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

func (lexer *Lexer) readString() string {
	position := lexer.position

	lexer.readChar()

	for lexer.char != '"' {
		lexer.readChar()
	}

	lexer.readChar()

	return lexer.input[position:lexer.position]
}
