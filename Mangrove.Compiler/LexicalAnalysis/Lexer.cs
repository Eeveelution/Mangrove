namespace Mangrove.Compiler;

public class Lexer {
    public string InputString;
    public int    CurrentPosition;
    public int    ReadPosition;
    public char   CurrentCharacter;

    public int Line, Column;

    public Lexer(string inputString) {
        this.InputString = inputString;

        this.ReadChar();
    }

    public void ReadChar() {
        if (this.ReadPosition >= this.InputString.Length) {
            this.CurrentCharacter = '\0';
        } else {
            this.CurrentCharacter = this.InputString[this.ReadPosition];
        }

        this.CurrentPosition = this.ReadPosition;
        this.ReadPosition++;

        this.Column++;
    }

    public char PeekChar() {
        if (this.ReadPosition >= this.InputString.Length) {
            return '\0';
        } else {
            return this.InputString[this.ReadPosition];
        }
    }

    public void SkipWhitespace() {
        while (this.CurrentCharacter == ' '
               || this.CurrentCharacter == '\r'
               || this.CurrentCharacter == '\n'
               || this.CurrentCharacter == '\t'
        ) {
            if (this.CurrentCharacter == '\n') {
                this.Line++;
            }

            this.ReadChar();
        }
    }

    private Token NewToken(TokenType type, char character) {
        return new Token {
            Line = this.Line,
            Column = this.Column,
            Literal = $"{character}",
            TokenType = type
        };
    }

    private Token NewToken(TokenType type, string literal) {
        return new Token {
            Line      = this.Line,
            Column    = this.Column,
            Literal   = literal,
            TokenType = type
        };
    }

    private bool IsDigit(char c) {
        return '0' <= c && c <= '9';
    }

    private bool IsOperatorSymbol(char c) {
        return c == '=' ||
               c == '-' ||
               c == '+' ||
               c == '*' ||
               c == '/' ||
               c == '<' ||
               c == '>' ||
               c == '|' ||
               c == '^' ||
               c == '&' ||
               c == '%' ||
               c == ':' ||
               c == '!';
    }

    private bool CanBeInIdentifier(char c, bool identifierBeginning) {
        return c is >= 'a' and <= 'z' || c is >= 'A' and <= 'Z' || c == '_' ||
               //Special handling for numbers in identifiers, they cannot be at the beginning
               //to not confuse them with integers
               (!identifierBeginning && this.IsDigit(c));
    }

    private string ReadIdentifier() {
        int currentPosition = this.CurrentPosition;
        bool identifierBeginning = true;

        while (this.CanBeInIdentifier(this.CurrentCharacter, identifierBeginning)) {
            identifierBeginning = false;

            this.ReadChar();
        }

        return this.InputString[currentPosition..this.CurrentPosition];
    }

    private (string literal, TokenType numberTokenType) ReadNumber() {
        TokenType numberTokenType = TokenType.Int;
        int currentPosition = this.CurrentPosition;

        while (true) {
            if (this.CurrentCharacter == '.') {
                numberTokenType = TokenType.Double;
                this.ReadChar();
            } else if (this.CurrentCharacter == 'f') {
                numberTokenType = TokenType.Float;
                this.ReadChar();
                break;
            } else if (this.IsDigit(this.CurrentCharacter)) {
                this.ReadChar();
            } else {
                break;
            }
        }

        return (this.InputString[currentPosition..this.CurrentPosition], numberTokenType);
    }

    private (string literal, TokenType tokenType) ReadOperator() {
        string literal = "";
        TokenType operatorType = TokenType.Illegal;

        while (this.IsOperatorSymbol(this.CurrentCharacter)) {
            literal += this.CurrentCharacter;

            this.ReadChar();
        }

        operatorType = literal switch {
            ":"   => TokenType.Colon,
            "="   => TokenType.Assign,
            "+"   => TokenType.Plus,
            "-"   => TokenType.Minus,
            "*"   => TokenType.Asterisk,
            "/"   => TokenType.Slash,
            "<"   => TokenType.LessThan,
            ">"   => TokenType.GreaterThan,
            "|"   => TokenType.Pipe,
            "^"   => TokenType.Caret,
            "&"   => TokenType.Ampersand,
            "%"   => TokenType.Percent,
            "!"   => TokenType.ExclamationMark,
            "=="  => TokenType.Equality,
            "!="  => TokenType.Inequality,
            "&&"  => TokenType.LogicalAnd,
            "||"  => TokenType.LogicalOr,
            "<="  => TokenType.LessOrEqual,
            ">="  => TokenType.GreaterOrEqual,
            "+="  => TokenType.PlusAssign,
            "-="  => TokenType.MinusAssign,
            "/="  => TokenType.DivisionAssign,
            "*="  => TokenType.MultiplicationAssign,
            "%="  => TokenType.ModAssign,
            "<<=" => TokenType.BitShiftLeftAssign,
            ">>=" => TokenType.BitShiftRightAssign,
            "&="  => TokenType.AndAssign,
            "^="  => TokenType.XorAssign,
            "|="  => TokenType.OrAssign,
            "::"  => TokenType.Namespace,
            _     => operatorType
        };

        return (literal, operatorType);
    }

    public Token NextToken() {
        this.SkipWhitespace();

        Token token;

        switch (this.CurrentCharacter) {
            case '?':
                token = this.NewToken(TokenType.QuestionMark, this.CurrentCharacter);
                break;
            case '#':
                token = this.NewToken(TokenType.HashSign, this.CurrentCharacter);
                break;
            case '$':
                token = this.NewToken(TokenType.DollarSign, this.CurrentCharacter);
                break;
            case ',':
                token = this.NewToken(TokenType.Comma, this.CurrentCharacter);
                break;
            case ';':
                token = this.NewToken(TokenType.Semicolon, this.CurrentCharacter);
                break;
            case '(':
                token = this.NewToken(TokenType.LeftParenthesis, this.CurrentCharacter);
                break;
            case ')':
                token = this.NewToken(TokenType.RightParenthesis, this.CurrentCharacter);
                break;
            case '{':
                token = this.NewToken(TokenType.LeftBrace, this.CurrentCharacter);
                break;
            case '}':
                token = this.NewToken(TokenType.RightBrace, this.CurrentCharacter);
                break;
            case '[':
                token = this.NewToken(TokenType.LeftBracket, this.CurrentCharacter);
                break;
            case ']':
                token = this.NewToken(TokenType.RightBracket, this.CurrentCharacter);
                break;
            case '"':
                int currentPosition = this.CurrentPosition;

                this.ReadChar();

                while (this.CurrentCharacter != '"') {
                    this.ReadChar();
                }

                this.ReadChar();

                string resultingString = this.InputString[currentPosition..this.CurrentPosition];

                token = this.NewToken(TokenType.String, resultingString);
                break;
            case '\0':
                token = this.NewToken(TokenType.EndOfFile, "");
                break;
            default:
                if (this.CanBeInIdentifier(this.CurrentCharacter, true)) {
                    return this.NewToken(TokenType.Identifier, this.ReadIdentifier());
                }

                if (this.IsDigit(this.CurrentCharacter)) {
                    (string literal, TokenType numberTokenType) = this.ReadNumber();

                    return this.NewToken(numberTokenType, literal);
                }

                if (this.IsOperatorSymbol(this.CurrentCharacter)) {
                    (string literal, TokenType operatorType) = this.ReadOperator();

                    return this.NewToken(operatorType, literal);
                }

                return this.NewToken(TokenType.Illegal, this.CurrentCharacter);
        }

        this.ReadChar();

        return token;
    }
}