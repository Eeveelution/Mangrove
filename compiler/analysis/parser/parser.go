package parser

import (
	"mangrove/compiler/analysis"
	"mangrove/compiler/analysis/token"
)

type Parser struct {
	tokens       []token.Token
	position     int
	readPosition int
	tok          token.Token
}

func CreateNewParser(lexer *analysis.Lexer) *Parser {
	var tokens []token.Token
	var tok token.Token

	for {
		tok = lexer.NextToken()

		if tok.Type == token.EOF {
			break
		} else {
			tokens = append(tokens, tok)
		}
	}

	return &Parser{
		tokens: tokens,
	}
}

func (parser *Parser) PeekNextToken() token.Token {
	return parser.tokens[parser.readPosition]
}

func (parser *Parser) NextToken() {
	if parser.readPosition >= len(parser.tokens) {
		parser.tok = token.Token{Type: token.EOF}
	} else {
		parser.tok = parser.tokens[parser.readPosition]
	}

	parser.position = parser.readPosition
	parser.readPosition += 1
}

func (parser *Parser) NextExpression() Expression {
	var expression Expression
	var tok token.Token

	for {
		tok = parser.PeekNextToken()

		if tok.Type == token.EOF {
			break
		}

		switch tok.Type {
		case token.VARIABLE:
			expression = parser.ParseVariableDeclarationExpression()
			return expression
		}
	}

	return expression
}
