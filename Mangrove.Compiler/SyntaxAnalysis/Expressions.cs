namespace Mangrove.Compiler.SyntaxAnalysis;

public abstract record ExpressionValue {
    public record Variable(TypeValue TypeValue, string Name) : ExpressionValue;
    public record InfixExpression(ExpressionOperator Operator, Expression Left, Expression Right) : ExpressionValue;
}

public class Expression {
    public ExpressionValue Value;

    public Expression(ExpressionValue value) {
        this.Value = value;

    }

    public override string ToString() => $"Expression: {this.Value}";
}