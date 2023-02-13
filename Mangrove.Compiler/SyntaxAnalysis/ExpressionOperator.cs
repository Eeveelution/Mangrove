namespace Mangrove.Compiler.SyntaxAnalysis;

public enum ExpressionOperatorPrescedence {
    Lowest = 1,
    Assign,
    Namespace,
    LogicalOr,
    LogicalAnd,
    BitwiseOr,
    BitwiseAnd,
    Equality,
    Relational,
    Add,
    Mul,
    Not,
    //source: rollie
    LogicalBitShiftLeft,
    LogicalBitShiftRight,
    LogicalXor,
}

[Flags]
public enum ExpressionOperatorType : byte {
    Prefix = 1 << 0,
    Infix  = 1 << 1,
}

public enum ExpressionOperatorAssociationDirection {
    Left,
    Right
}

public abstract record ExpressionOperator {
    public abstract ExpressionOperatorPrescedence Prescedence { get; }
    public abstract ExpressionOperatorType Type { get; }

    public static ExpressionOperator FromTokenType(TokenType tokenType) {
        return tokenType switch {
            TokenType.Assign               => new Assign(),
            TokenType.Plus                 => new Add(),
            TokenType.Minus                => new Sub(),
            TokenType.Asterisk             => new Mul(),
            TokenType.Slash                => new Div(),
            TokenType.LogicalAnd           => new LogicalAnd(),
            TokenType.LogicalOr            => new LogicalOr(),
            TokenType.Caret                => new Xor(),
            TokenType.Percent              => new Rem(),
            TokenType.BitShiftLeft         => new BitShiftLeft(),
            TokenType.BitShiftRight        => new BitShiftRight(),
            TokenType.ExclamationMark      => new Not(),
            TokenType.Inequality           => new NotEq(),
            TokenType.Equality             => new Eq(),
            TokenType.GreaterThan          => new GreaterThan(),
            TokenType.GreaterOrEqual       => new GreaterThanOrEqual(),
            TokenType.LessThan             => new LessThan(),
            TokenType.LessOrEqual          => new LessThanOrEqual(),
            TokenType.AndAssign            => new OperatorWithAssignment(new BitwiseAnd()),
            TokenType.OrAssign             => new OperatorWithAssignment(new BitwiseOr()),
            TokenType.XorAssign            => new OperatorWithAssignment(new Xor()),
            TokenType.PlusAssign           => new OperatorWithAssignment(new Add()),
            TokenType.MinusAssign          => new OperatorWithAssignment(new Sub()),
            TokenType.MultiplicationAssign => new OperatorWithAssignment(new Mul()),
            TokenType.DivisionAssign       => new OperatorWithAssignment(new Div()),
            TokenType.RemAssign            => new OperatorWithAssignment(new Rem()),
            TokenType.BitShiftLeftAssign   => new OperatorWithAssignment(new BitShiftLeft()),
            TokenType.BitShiftRightAssign  => new OperatorWithAssignment(new BitShiftRight()),
            _                              => throw new ArgumentOutOfRangeException(nameof(tokenType), "what")
        };
    }

    public record Assign : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Assign;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record OperatorWithAssignment(ExpressionOperator Wrapped) : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Assign;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record Add : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Add;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix | ExpressionOperatorType.Prefix;
    }

    public record Sub : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Add;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix | ExpressionOperatorType.Prefix;
    }

    public record Mul : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Mul;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record Div : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Mul;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record Rem : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Mul;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record LogicalAnd : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.LogicalAnd;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record LogicalOr : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.LogicalOr;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record Eq : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Equality;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record NotEq : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Equality;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record GreaterThan : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Relational;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record GreaterThanOrEqual : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Relational;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record LessThan : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Relational;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record LessThanOrEqual : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Relational;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Infix;
    }

    public record Not : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.Not;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Prefix;
    }

    public record BitwiseAnd : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.BitwiseAnd;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Prefix;
    }

    public record BitwiseOr : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.BitwiseOr;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Prefix;
    }

    public record BitShiftLeft : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.LogicalBitShiftLeft;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Prefix;
    }

    public record BitShiftRight : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.LogicalBitShiftRight;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Prefix;
    }

    public record Xor : ExpressionOperator {
        public override ExpressionOperatorPrescedence Prescedence => ExpressionOperatorPrescedence.LogicalXor;
        public override ExpressionOperatorType Type => ExpressionOperatorType.Prefix;
    }
}
