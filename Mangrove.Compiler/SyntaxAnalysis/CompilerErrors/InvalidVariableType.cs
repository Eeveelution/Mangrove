using Mangrove.Compiler.SyntaxAnalysis.CompilerErrors;

namespace Mangrove.Compiler.SyntaxAnalysis;

public class InvalidVariableType : CompilerError {
    protected sealed override string Description { get; init; }

    public InvalidVariableType(int line, int column, Token receivedToken) : base(2, line, column) {
        this.Description = $"Received variable type is invalid! Received variable type: {receivedToken.TokenType}";
    }
}
