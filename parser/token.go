package parser

type TokenType int

const (
	Identifier TokenType = iota
	String
	OpenParen
	CloseParen
	Equals
)

type Token struct {
	Line, Offset int
	Type         TokenType
	Val          interface{}
}
