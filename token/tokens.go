package token

type TokenType string // Allow to use many different values as token types

type Token struct {
    Type    TokenType
    Literal string
}

const (
    
    ILLEGAL = "ILLEGAL" 
    EOF     = "EOF"     

    // Identifiers and literals
    IDENT = "IDENT" 
    INT   = "INT"   
    FLOAT = "FLOAT"
    STRING = "STRING" 

    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    ASTERISK = "*"
    SLASH    = "/"
    EQ       = "=="
    NOT_EQ   = "!="
    LT       = "<"
    GT       = ">"
    LTE      = "<="
    GTE      = ">="

    COMMA     = ","
    COLON     = ":"
    SEMICOLON = ";"
    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"
    LBRACKET  = "["
    RBRACKET  = "]"

    // Indentation tokens
    INDENT = "INDENT" 
    DEDENT = "DEDENT" 

    // Keywords
    DEF      = "DEF"
    CLASS    = "CLASS"
    IF       = "IF"
    ELSE     = "ELSE"
    ELIF     = "ELIF"
    FOR      = "FOR"
    WHILE    = "WHILE"
    RETURN   = "RETURN"
    IMPORT   = "IMPORT"
    FROM     = "FROM"
    AS       = "AS"
    TRY      = "TRY"
    EXCEPT   = "EXCEPT"
    FINALLY  = "FINALLY"
    WITH     = "WITH"
    LAMBDA   = "LAMBDA"
    PASS     = "PASS"
    BREAK    = "BREAK"
    CONTINUE = "CONTINUE"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    NONE     = "NONE"
)

var keywords = map[string]TokenType{
    "def":      DEF,
    "class":    CLASS,
    "if":       IF,
    "else":     ELSE,
    "elif":     ELIF,
    "for":      FOR,
    "while":    WHILE,
    "return":   RETURN,
    "import":   IMPORT,
    "from":     FROM,
    "as":       AS,
    "try":      TRY,
    "except":   EXCEPT,
    "finally":  FINALLY,
    "with":     WITH,
    "lambda":   LAMBDA,
    "pass":     PASS,
    "break":    BREAK,
    "continue": CONTINUE,
    "True":     TRUE,
    "False":    FALSE,
    "None":     NONE,
}

// LookupIdent checks if an identifier is a keyword or a user-defined name.
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}




