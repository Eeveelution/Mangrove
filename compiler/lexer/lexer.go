package lexer

import "mangrove/compiler/lexer/token"

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
	case '=':
		tok = newToken(token.ASSIGN, lexer.char)
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case '-':
		tok = newToken(token.MINUS, lexer.char)
	case '*':
		tok = newToken(token.MULTIPLICATION, lexer.char)
	case '/':
		tok = newToken(token.DIVISION, lexer.char)
	case '<':
		tok = newToken(token.LESS_THAN, lexer.char)
	case '>':
		tok = newToken(token.GREATER_THAN, lexer.char)
	case '|':
		tok = newToken(token.BINARY_OR, lexer.char)
	case '^':
		tok = newToken(token.BINARY_XOR, lexer.char)
	case '&':
		tok = newToken(token.BINARY_AND, lexer.char)

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
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()

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

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func canBeInIdentifier(char byte, identifierBeginning bool) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' ||
		//Special handling for numbers in identifiers, they cannot be at the beginning
		//to not confuse them with integers
		(!identifierBeginning && ('0' <= char && char <= '9'))
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}
