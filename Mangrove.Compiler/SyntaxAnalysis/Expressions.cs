namespace Mangrove.Compiler.SyntaxAnalysis;

public abstract record ExpressionValue {
    public record Variable(TypeValue TypeValue, string Name);
}

public class Expression {
    public ExpressionValue Value;

    public Expression(ExpressionValue value) {
        this.Value = value;

    }

    public override string ToString() => $"Expression: {this.Value}";
}