using Mangrove.Compiler.SyntaxAnalysis.CompilerErrors;

namespace Mangrove.Compiler.SyntaxAnalysis;

public class InvalidInfixOperator : CompilerError {
    protected sealed override string Description { get; init; }

    public InvalidInfixOperator(int line, int column, Token receivedToken) : base(2, line, column) {
        this.Description = $"Invalid Infix Operator!";
    }
}
