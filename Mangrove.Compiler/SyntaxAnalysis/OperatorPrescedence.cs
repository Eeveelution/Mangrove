namespace Mangrove.Compiler.SyntaxAnalysis;

public enum OperatorPrescedence {
    Lowest = 1,
    Assign,
    LogicalOr,
    LogicalAnd,
    BitwiseOr,
    BitwiseAnd,
    Equality,
    Relational,
    Add,
    Mul,
    Not,
    Namespace
}

[Flags]
public enum ExpressionOperatorType : byte {
    Prefix = 1 << 0,
    Infix  = 1 << 1,
}
