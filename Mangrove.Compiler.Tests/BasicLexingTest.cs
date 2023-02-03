namespace Mangrove.Compiler.Tests;

public class Tests {
    [SetUp]
    public void Setup() {}

    [Test]
    public void TestBasicLexing() {
        string testInput = """
            var<TcpServer> server = [::TcpServer newWithAddress: "127.0.0.1" port: 13381];
            var<float> testFloat = 4.5147281f;
            var<double> testDouble = 4.5147281;
        """;

        Lexer lexer = new Lexer(testInput);

        Token[] expectedTokens = new [] {
            new Token { TokenType = TokenType.Identifier, Literal = "var"},
            new Token { TokenType = TokenType.LessThan },
            new Token { TokenType = TokenType.Identifier, Literal = "TcpServer"},
            new Token { TokenType = TokenType.GreaterThan },
            new Token { TokenType = TokenType.Identifier, Literal = "server"},
            new Token { TokenType = TokenType.Assign },
            new Token { TokenType = TokenType.LeftBracket },
            new Token { TokenType = TokenType.Namespace },
            new Token { TokenType = TokenType.Identifier, Literal = "TcpServer"},
            new Token { TokenType = TokenType.Identifier, Literal = "newWithAddress"},
            new Token { TokenType = TokenType.Colon },
            new Token { TokenType = TokenType.String, Literal     = "\"127.0.0.1\""},
            new Token { TokenType = TokenType.Identifier, Literal = "port"},
            new Token { TokenType = TokenType.Colon },
            new Token { TokenType = TokenType.Int, Literal = "13381"},
            new Token { TokenType = TokenType.RightBracket},
            new Token { TokenType = TokenType.Semicolon},
            //Line 2
            new Token { TokenType = TokenType.Identifier, Literal = "var"},
            new Token { TokenType = TokenType.LessThan },
            new Token { TokenType = TokenType.Identifier, Literal = "float"},
            new Token { TokenType = TokenType.GreaterThan },
            new Token { TokenType = TokenType.Identifier, Literal = "testFloat"},
            new Token { TokenType = TokenType.Assign },
            new Token { TokenType = TokenType.Float, Literal = "4.5147281f"},
            new Token { TokenType = TokenType.Semicolon},
            //Line 3
            new Token { TokenType = TokenType.Identifier, Literal = "var"},
            new Token { TokenType = TokenType.LessThan },
            new Token { TokenType = TokenType.Identifier, Literal = "double"},
            new Token { TokenType = TokenType.GreaterThan },
            new Token { TokenType = TokenType.Identifier, Literal = "testDouble"},
            new Token { TokenType = TokenType.Assign },
            new Token { TokenType = TokenType.Double, Literal = "4.5147281"},
            new Token { TokenType = TokenType.Semicolon},
        };

        Token currentToken;
        int i = 0;

        while ((currentToken = lexer.NextToken()).TokenType != TokenType.EndOfFile) {
            Token currentExpected = expectedTokens[i];

            if (currentExpected.TokenType != currentToken.TokenType) {
                Assert.Fail($"TokenType different! Index: {i}; Expected: {currentExpected.TokenType}; Received: {currentToken.TokenType}");
            }

            if (!string.IsNullOrEmpty(currentExpected.Literal) && currentToken.Literal != currentExpected.Literal) {
                Assert.Fail($"Literal different! Index: {i}; Expected: {currentExpected.Literal}; Got: {currentToken.Literal}");
            }

            i++;
        }

        Assert.Pass();
    }
}
