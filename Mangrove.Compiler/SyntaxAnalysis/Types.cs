namespace Mangrove.Compiler.SyntaxAnalysis;

public class TypeValue {
    public TokenType VariableType;
    public object    Value;
    public int       ArrayLength;

    public bool IsArray => this.ArrayLength != 0;

    public TypeValue(TokenType variableType) {
        this.VariableType = variableType;
    }

    public TypeValue(TokenType arrayType, int length) {
        this.VariableType = arrayType;
        this.ArrayLength  = length;
    }

    public override string ToString() => $"Type: {this.Value}";
}