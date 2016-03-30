package scrap

import (
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	scanner *Scanner
}

func Parse(reader io.Reader) (Expression, error) {
	p := Parser{scanner: NewScanner(reader)}
	return p.ParseExpression()
}

func ParseString(input string) (Expression, error) {
	return Parse(strings.NewReader(input))
}

type ParseError struct {
	Expected []TokenType
	Actual   *Token
}

func (e *ParseError) Error() string {
	if e.Actual == nil {
		return "unexpected EOF"
	}

	switch len(e.Expected) {
	case 0:
		return fmt.Sprintf("(%d,%d) unexpected %s", e.Actual.Line, e.Actual.Offset, e.Actual.Type)
	case 1:
		return fmt.Sprintf("(%d,%d) unexpected %s; expected %s", e.Actual.Line, e.Actual.Offset, e.Actual.Type, e.Expected[0])
	default:
		expected := make([]string, 0, len(e.Expected))
		for _, e := range e.Expected {
			expected = append(expected, e.String())
		}

		commaSeparated := strings.Join(expected[:len(expected)-1], ", ")
		return fmt.Sprintf("(%d,%d) unexpected %s; expected %s or %s", e.Actual.Line, e.Actual.Offset, e.Actual.Type, commaSeparated, expected[len(expected)-1])
	}
}

func (p *Parser) ParseExpression() (Expression, error) {
	next, err := p.scanner.Next()
	if err != nil {
		return nil, err
	}

	switch next.Type {
	case String:
		lit := StringLiteral(next.Val.(string))
		return &lit, nil
	}

	if next.Type != Ident {
		return nil, &ParseError{
			Expected: []TokenType{Ident},
			Actual:   next,
		}
	}
	ident := next.Val.(string)

	next, err = p.scanner.Next()
	if err != nil {
		return nil, err
	}
	if next.Type != Equals {
		return nil, &ParseError{
			Expected: []TokenType{Equals},
			Actual:   next,
		}
	}

	right, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	return &Assignment{
		Left:  ident,
		Right: right,
	}, nil
}
