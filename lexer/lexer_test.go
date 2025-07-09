package lexer

import (
    "interpreter/token"
    "testing"
	// "fmt"
)

func TestNextToken(t *testing.T) {
	input := `def myFunc():
    x = 42
    if x > 10:
        print("Hello, World!")
    else:
        pass`

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.DEF, "def"},
        {token.IDENT, "myFunc"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.COLON, ":"},
        {token.INDENT, ""},
        {token.IDENT, "x"},
        {token.ASSIGN, "="},
        {token.INT, "42"},
        {token.IF, "if"},
        {token.IDENT, "x"},
        {token.GT, ">"},
        {token.INT, "10"},
        {token.COLON, ":"},
        {token.INDENT, ""},
        {token.IDENT, "print"},
        {token.LPAREN, "("},
        {token.STRING, `"Hello, World!"`},
        {token.RPAREN, ")"},
        {token.DEDENT, ""},
        {token.ELSE, "else"},
        {token.COLON, ":"},
        {token.INDENT, ""},
        {token.PASS, "pass"},
        {token.DEDENT, ""},
        {token.DEDENT, ""},
        {token.EOF, ""},
    }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }
}