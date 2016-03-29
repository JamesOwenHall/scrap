package parser

import (
	"fmt"
)

const (
	Identifier TokenType = iota
	String
	OpenParen
	CloseParen
	Comma
	Period
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
	case Comma:
		return "comma"
	case Period:
		return "period"
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

func (t Token) String() string {
	return fmt.Sprintf("%s(%v)@[%d,%d]", t.Type.String(), t.Val, t.Line, t.Offset)
}
