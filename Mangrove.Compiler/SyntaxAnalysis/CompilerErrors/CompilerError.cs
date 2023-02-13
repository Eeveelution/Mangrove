namespace Mangrove.Compiler.SyntaxAnalysis.CompilerErrors;

public abstract class CompilerError : Exception {
    protected abstract string Description { get; init; }
    private readonly int _code, _line, _column;

    public CompilerError(int code, int line, int column) {
        this._code   = code;
        this._line   = line;
        this._column = column;
    }

    public override string Message => $"Compiler Error occured! {this._code:0000} ({this._column};{this._line}: {this.Description}";
}
