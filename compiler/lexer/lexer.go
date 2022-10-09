package lexer

type Lexer struct {
	input        string //Input String
	position     int    //Position in the input (points to current char)
	readPosition int    //Current reading position in input (points to current char + 1)
	ch           byte   //Current Char under examination
}

func CreateNewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	return lexer
}
