namespace Mangrove.Compiler.SyntaxAnalysis;

public class Parser {
    public Token[] InputTokens;
    public int     CurrentPosition;
    public int     ReadPosition;
    public Token   CurrentToken;

    public Parser(IEnumerable<Token> tokens) {
        this.InputTokens = tokens.ToArray();
    }

    public void ReadToken() {
        if (this.ReadPosition >= this.InputTokens.Length) {
            this.CurrentToken = new Token { TokenType = TokenType.EndOfFile };
        } else {
            this.CurrentToken = this.InputTokens[this.ReadPosition];
        }

        this.CurrentPosition = this.ReadPosition;
        this.ReadPosition++;
    }

    public Token PeekToken() {
        if (this.ReadPosition >= this.InputTokens.Length) {
            return new Token { TokenType = TokenType.EndOfFile };
        }

        return this.InputTokens[this.ReadPosition];
    }

    public ExpressionValue.Variable ParseVariable() {
        ExpressionValue.Variable returnVariable;

        this.ReadToken(); //var

        if (this.CurrentToken.TokenType != TokenType.GreaterThan) {
            throw new ExpectedTypeParameterOnVariable(this.CurrentToken.Line, this.CurrentToken.Column);
        }

        this.ReadToken(); //<

        Token typeLiteral = this.CurrentToken;

        Dictionary<TokenType, bool> allowedTypesAndConversion = new Dictionary<TokenType, bool>() {
            [TokenType.Identifier] = true, [TokenType.TypeInt8]   = true, [TokenType.TypeInt16]  = true,
            [TokenType.TypeInt32]  = true, [TokenType.TypeInt64]  = true, [TokenType.TypeUint8]  = true,
            [TokenType.TypeUint16] = true, [TokenType.TypeUint32] = true, [TokenType.TypeUint64] = true,
            [TokenType.TypeDouble] = true, [TokenType.TypeFloat]  = true, [TokenType.TypeString] = true,
            [TokenType.Array]      = true
        };

        if (!allowedTypesAndConversion.ContainsKey(typeLiteral.TokenType)) {
            throw new InvalidVariableType(typeLiteral.Line, typeLiteral.Column, typeLiteral);
        }

        /* WIP Array implementation

        if (typeLiteral.TokenType == TokenType.Array) {
            this.ReadToken(); //arr

            if (this.CurrentToken.TokenType != TokenType.GreaterThan) {
                throw new ExpectedTypeParameterOnArray(this.CurrentToken.Line, this.CurrentToken.Column);
            }

            this.ReadToken(); //<

            Token arrayTypeLiteral = this.CurrentToken;

            if (!allowedTypesAndConversion.ContainsKey(arrayTypeLiteral.TokenType)) {
                throw new InvalidVariableType(arrayTypeLiteral.Line, arrayTypeLiteral.Column, arrayTypeLiteral);
            }

            this.ReadToken(); // type

            if (this.CurrentToken.TokenType == TokenType.Comma) {
                this.ReadToken(); //,

                Token lengthToken = this.CurrentToken;



                returnVariable.TypeValue = new TypeValue(arrayTypeLiteral, lengthToken.Literal)
            } else {

            }
        }

        */

        this.ReadToken(); //>

        Token nameLiteral = this.CurrentToken;

        if (nameLiteral.TokenType != TokenType.Identifier) {
            throw new ExpectedNameOnVariable(this.CurrentToken.Line, this.CurrentToken.Column);
        }
    }

    public Expression ParseNextExpression(OperatorPrescedence prescedence = OperatorPrescedence.Lowest) {
        Expression expression = default;

        switch (this.CurrentToken.TokenType) {
            case TokenType.Variable:
                ExpressionValue.Variable variable = this.ParseVariable();
                break;
        }
    }
}
