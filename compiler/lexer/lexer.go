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
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.LPARENTHESIS, lexer.char)
	case ')':
		tok = newToken(token.RPARENTHESIS, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lexer.readChar()

	return tok
}
