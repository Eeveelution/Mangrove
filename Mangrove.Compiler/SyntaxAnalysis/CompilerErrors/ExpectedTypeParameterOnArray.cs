using Mangrove.Compiler.SyntaxAnalysis.CompilerErrors;

namespace Mangrove.Compiler.SyntaxAnalysis;

public class ExpectedTypeParameterOnArray : CompilerError {
    protected sealed override string Description { get; init; }

    public ExpectedTypeParameterOnArray(int line, int column) : base(1, line, column) {
        this.Description = "A type parameter on this array would've been expected; Example var<arr<int>>";
    }
}
