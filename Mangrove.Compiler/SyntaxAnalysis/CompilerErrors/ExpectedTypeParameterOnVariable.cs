using Mangrove.Compiler.SyntaxAnalysis.CompilerErrors;

namespace Mangrove.Compiler.SyntaxAnalysis;

public class ExpectedTypeParameterOnVariable : CompilerError {
    protected sealed override string Description { get; init; }

    public ExpectedTypeParameterOnVariable(int line, int column) : base(1, line, column) {
        this.Description = "A type parameter on this variable would've been expected; Example var<int>";
    }
}
