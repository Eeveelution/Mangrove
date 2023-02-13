using Mangrove.Compiler.SyntaxAnalysis.CompilerErrors;

namespace Mangrove.Compiler.SyntaxAnalysis;

public class ExpectedNameOnVariable : CompilerError {
    protected sealed override string Description { get; init; }

    public ExpectedNameOnVariable(int line, int column) : base(1, line, column) {
        this.Description = "A following var<type> is the variable name.";
    }
}
