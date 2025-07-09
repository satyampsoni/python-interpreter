package lexer

import (
    "interpreter/token"
)

type Lexer struct {
    input         string
    position      int
    readPosition  int
    ch            byte
    currentIndent int
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    if l.ch == '\n' {
        l.readChar()
        tok = l.handleIndentation()
        if tok.Type != token.ILLEGAL {
            return tok
        }
    }

    l.skipWhitespace()

    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '*':
        tok = newToken(token.ASTERISK, l.ch)
    case '/':
        tok = newToken(token.SLASH, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case '[':
        tok = newToken(token.LBRACKET, l.ch)
    case ']':
        tok = newToken(token.RBRACKET, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case ':':
        tok = newToken(token.COLON, l.ch)
    case '>':
        tok = newToken(token.GT, l.ch)
    case '"':
        tok.Literal = l.readString('"')
        tok.Type = token.STRING
        return tok
    case '\'':
        tok.Literal = l.readString('\'')
        tok.Type = token.STRING
        return tok
    case 0:
        if l.currentIndent > 0 {
            return l.handleEOF()
        }
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber()
            tok.Type = token.INT
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' {
        l.readChar()
    }
}

// Manages Python-style indentation tokens
func (l *Lexer) handleIndentation() token.Token {
    indentLevel := 0

    for l.ch == ' ' || l.ch == '\t' {
        if l.ch == ' ' {
            indentLevel++
        } else if l.ch == '\t' {
            indentLevel += 4
        }
        l.readChar()
    }

    if indentLevel > l.currentIndent {
        l.currentIndent = indentLevel
        return token.Token{Type: token.INDENT, Literal: ""}
    } else if indentLevel < l.currentIndent {
        l.currentIndent -= 4
        return token.Token{Type: token.DEDENT, Literal: ""}
    }

    return token.Token{Type: token.ILLEGAL, Literal: ""}
}

func (l *Lexer) handleEOF() token.Token {
    if l.currentIndent > 0 {
        l.currentIndent -= 4
        return token.Token{Type: token.DEDENT, Literal: ""}
    }
    return token.Token{Type: token.EOF, Literal: ""}
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) || l.ch == '.' {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readString(quote byte) string {
    position := l.position + 1
    for {
        l.readChar()
        if l.ch == quote || l.ch == 0 {
            break
        }
    }
    str := l.input[position:l.position]
    l.readChar()
    return str
}

func isLetter(ch byte) bool {
    return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}