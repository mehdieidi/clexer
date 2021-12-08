// Package token contains tokens of C language used in lexical analyzer
package token

// tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT   = "IDENTIFIER"
	TEXT    = "TEXT"
	COMMENT = "COMMENT"
	INTEGER = "INTEGER"
	DECIMAL = "DECIMAL"

	// operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	EXCLA     = "!"
	ST        = "<"
	GT        = ">"
	EQ        = "=="
	NOT_EQ    = "!="
	INC       = "++"
	DEC       = "--"
	PERCENT   = "%"
	STEQ      = "<="
	GTEQ      = ">="
	PLUSAS    = "+="
	MINUSAS   = "-="
	MULTAS    = "*="
	DIVAS     = "/="
	PERCENTAS = "%="
	AMPERAS   = "&="
	XORAS     = "^="
	PIPEAS    = "|="
	LSHIFTAS  = "<<="
	RSHIFTAS  = ">>="
	AMPER     = "&"
	PIPE      = "|"
	AND       = "&&"
	OR        = "||"
	XOR       = "^"
	ARROW     = "->"
	TILDA     = "~"
	RSHIFT    = ">>"
	LSHIFT    = "<<"
	QUESTION  = "?"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	// parentheses
	LBRACK = "["
	RBRACK = "]"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	AUTO     = "AUTO"
	BREAK    = "BREAK"
	CASE     = "CASE"
	CHAR     = "CHAR"
	CONST    = "CONST"
	CONTINUE = "CONTINUE"
	DEFAULT  = "DEFAULT"
	DO       = "DO"
	INT      = "INT"
	LONG     = "LONG"
	REGISTER = "REGISTER"
	SHORT    = "SHORT"
	SIGNED   = "SIGNED"
	SIZEOF   = "SIZEOF"
	STATIC   = "STATIC"
	STRUCT   = "STRUCT"
	SWITCH   = "SWITCH"
	TYPEDEF  = "TYPEDEF"
	UNION    = "UNION"
	UNSIGNED = "UNSIGNED"
	VOID     = "VOID"
	VOLATILE = "VOLATILE"
	WHILE    = "WHILE"
	DOUBLE   = "DOUBLE"
	ENUM     = "ENUM"
	EXTERN   = "EXTERN"
	FLOAT    = "FLOAT"
	FOR      = "FOR"
	GOTO     = "GOTO"
)

type Token struct {
	Type    TokenType
	Literal string
	Row     uint
	Col     uint
	BlockNo uint
}

type TokenType string

var keywords = map[string]TokenType{
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"auto":     AUTO,
	"break":    BREAK,
	"case":     CASE,
	"char":     CHAR,
	"const":    CONST,
	"continue": CONTINUE,
	"default":  DEFAULT,
	"do":       DO,
	"int":      INT,
	"long":     LONG,
	"register": REGISTER,
	"short":    SHORT,
	"signed":   SIGNED,
	"sizeof":   SIZEOF,
	"static":   STATIC,
	"struct":   STRUCT,
	"switch":   SWITCH,
	"typedef":  TYPEDEF,
	"union":    UNION,
	"unsigned": UNSIGNED,
	"void":     VOID,
	"volatile": VOLATILE,
	"while":    WHILE,
	"double":   DOUBLE,
	"enum":     ENUM,
	"extern":   EXTERN,
	"float":    FLOAT,
	"for":      FLOAT,
	"goto":     GOTO,
}

// checks to see if the ident is actually an identifier or a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
