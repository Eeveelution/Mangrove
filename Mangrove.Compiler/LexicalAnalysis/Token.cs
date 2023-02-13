namespace Mangrove.Compiler;

public enum TokenType {
    Unset,
    Illegal,
    EndOfFile,

    //Identifiers and literals
    Identifier,
    Int,
    Float,
    Double,
    String,
    Byte,
    Void,

    //Operators
    Assign,
    Plus,
    Minus,
    Asterisk,
    Slash,
    LessThan,
    GreaterThan,
    Pipe,
    Caret,
    Ampersand,
    Percent,
    ExclamationMark,
    Tilde,
    BitShiftLeft,
    BitShiftRight,

    //Multichar Operators
    Equality,
    Inequality,
    LogicalAnd,
    LogicalOr,
    LessOrEqual,
    GreaterOrEqual,

    PlusAssign,
    MinusAssign,
    MultiplicationAssign,
    DivisionAssign,
    RemAssign,
    BitShiftLeftAssign,
    BitShiftRightAssign,
    AndAssign,
    XorAssign,
    OrAssign,

    //Delimiters
    Comma,
    Semicolon,
    Colon,
    Namespace,

    LeftParenthesis,
    RightParenthesis,
    LeftBrace,
    RightBrace,
    LeftBracket,
    RightBracket,
    Quote,

    //Keywords
    Function,
    Variable,
    Return,
    If,
    Else,
    True,
    False,
    While,
    For,
    Switch,
    Case,
    Default,
    Break,
    Use,
    Class,
    Static,
    Array,


    //Types
    TypeVoid,
    TypeInt64,
    TypeInt32,
    TypeInt16,
    TypeInt8,
    TypeUint64,
    TypeUint32,
    TypeUint16,
    TypeUint8,
    TypeString,
    TypeFloat,
    TypeDouble,

    //Others,
    QuestionMark,
    DollarSign,
    HashSign
}

public struct Token {
    public TokenType TokenType;
    public string    Literal;
    public int       Line, Column;
    public bool      IsOperator;
}
