package parser

const (
	Identifier TokenType = iota
	String
	OpenParen
	CloseParen
	Equals
)

type TokenType int

func (t TokenType) String() string {
	switch t {
	case Identifier:
		return "identifier"
	case String:
		return "string"
	case OpenParen:
		return "open-paren"
	case CloseParen:
		return "close-paren"
	case Equals:
		return "equals"
	default:
		panic("unreachable")
	}
}

type Token struct {
	Line, Offset int
	Type         TokenType
	Val          interface{}
}
