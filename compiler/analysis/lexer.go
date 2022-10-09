package analysis

import "mangrove/compiler/analysis/token"

type Lexer struct {
	input        string //Input String
	position     int    //Position in the input (points to current char)
	readPosition int    //Current reading position in input (points to current char + 1)
	char         byte   //Current Char under examination

	Line   int
	Column int
}

func CreateNewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: input,
		Line:  1,
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

	lexer.Column += 1
}

func (lexer *Lexer) newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
		Line:    lexer.Line,
		Column:  lexer.Column,
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	tok.Line = lexer.Line
	tok.Column = lexer.Column

	switch lexer.char {

	case '?':
		tok = lexer.newToken(token.QUESTIONMARK, lexer.char)
	case '$':
		tok = lexer.newToken(token.DOLLAR_SIGN, lexer.char)
	case '#':
		tok = lexer.newToken(token.HASH_SIGN, lexer.char)
	case ',':
		tok = lexer.newToken(token.COMMA, lexer.char)
	case ';':
		tok = lexer.newToken(token.SEMICOLON, lexer.char)
	case ':':
		tok = lexer.newToken(token.COLON, lexer.char)

	case '(':
		tok = lexer.newToken(token.LPARENTHESIS, lexer.char)
	case ')':
		tok = lexer.newToken(token.RPARENTHESIS, lexer.char)
	case '{':
		tok = lexer.newToken(token.LBRACE, lexer.char)
	case '}':
		tok = lexer.newToken(token.RBRACE, lexer.char)
	case '[':
		tok = lexer.newToken(token.LBRACKET, lexer.char)
	case ']':
		tok = lexer.newToken(token.RBRACKET, lexer.char)
	case '"':
		tok.Type = token.STRING
		tok.Literal = lexer.readString()

		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

		return tok
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
			tok = lexer.newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()

	return tok
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		if lexer.char == '\n' {
			lexer.Line += 1
			lexer.Column = 0
		}

		lexer.readChar()
	}
}
