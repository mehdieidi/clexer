// Package token contains tokens of C language used in lexical analyzer.
package token

// tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT   = "IDENTIFIER"
	TEXT    = "TEXT"
	INTEGER = "INTEGER"
	DECIMAL = "DECIMAL"

	// operators
	ASSIGN    = "OPERATOR: ="
	PLUS      = "OPERATOR: +"
	MINUS     = "OPERATOR: -"
	ASTERISK  = "OPERATOR: *"
	SLASH     = "OPERATOR: /"
	EXCLA     = "OPERATOR: !"
	ST        = "OPERATOR: <"
	GT        = "OPERATOR: >"
	EQ        = "OPERATOR: =="
	NOT_EQ    = "OPERATOR: !="
	INC       = "OPERATOR: ++"
	DEC       = "OPERATOR: --"
	PERCENT   = "OPERATOR: %"
	STEQ      = "OPERATOR: <="
	GTEQ      = "OPERATOR: >="
	PLUSAS    = "OPERATOR: +="
	MINUSAS   = "OPERATOR: -="
	MULTAS    = "OPERATOR: *="
	DIVAS     = "OPERATOR: /="
	PERCENTAS = "OPERATOR: %="
	AMPERAS   = "OPERATOR: &="
	XORAS     = "OPERATOR: ^="
	PIPEAS    = "OPERATOR: |="
	LSHIFTAS  = "OPERATOR: <<="
	RSHIFTAS  = "OPERATOR: >>="
	AMPER     = "OPERATOR: &"
	PIPE      = "OPERATOR: |"
	AND       = "OPERATOR: &&"
	OR        = "OPERATOR: ||"
	XOR       = "OPERATOR: ^"
	ARROW     = "OPERATOR: ->"
	TILDA     = "OPERATOR: ~"
	RSHIFT    = "OPERATOR: >>"
	LSHIFT    = "OPERATOR: <<"
	QUESTION  = "OPERATOR: ?"
	DOT       = "OPERATOR: ."

	// delimiters
	COMMA     = "DELIMITER: ,"
	SEMICOLON = "DELIMITER: ;"
	COLON     = "DELIMITER: :"

	// brackets
	LBRACK = "BRACKET: ["
	RBRACK = "BRACKET: ]"
	LPAREN = "BRACKET: ("
	RPAREN = "BRACKET: )"
	LBRACE = "BRACKET: {"
	RBRACE = "BRACKET: }"

	// keywords
	IF       = "KEYWORD: IF"
	ELSE     = "KEYWORD: ELSE"
	RETURN   = "KEYWORD: RETURN"
	AUTO     = "KEYWORD: AUTO"
	BREAK    = "KEYWORD: BREAK"
	CASE     = "KEYWORD: CASE"
	CHAR     = "KEYWORD: CHAR"
	CONST    = "KEYWORD: CONST"
	CONTINUE = "KEYWORD: CONTINUE"
	DEFAULT  = "KEYWORD: DEFAULT"
	DO       = "KEYWORD: DO"
	INT      = "KEYWORD: INT"
	LONG     = "KEYWORD: LONG"
	REGISTER = "KEYWORD: REGISTER"
	SHORT    = "KEYWORD: SHORT"
	SIGNED   = "KEYWORD: SIGNED"
	SIZEOF   = "KEYWORD: SIZEOF"
	STATIC   = "KEYWORD: STATIC"
	STRUCT   = "KEYWORD: STRUCT"
	SWITCH   = "KEYWORD: SWITCH"
	TYPEDEF  = "KEYWORD: TYPEDEF"
	UNION    = "KEYWORD: UNION"
	UNSIGNED = "KEYWORD: UNSIGNED"
	VOID     = "KEYWORD: VOID"
	VOLATILE = "KEYWORD: VOLATILE"
	WHILE    = "KEYWORD: WHILE"
	DOUBLE   = "KEYWORD: DOUBLE"
	ENUM     = "KEYWORD: ENUM"
	EXTERN   = "KEYWORD: EXTERN"
	FLOAT    = "KEYWORD: FLOAT"
	FOR      = "KEYWORD: FOR"
	GOTO     = "KEYWORD: GOTO"
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

// LookupIdent checks to see if the ident is actually an identifier or a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// New returns a new Token with the given params.
func New(tokenType TokenType, literal string, row uint, col uint, blockNo uint) Token {
	return Token{Type: tokenType, Literal: literal, Row: row, Col: col, BlockNo: blockNo}
}
