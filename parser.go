package scrap

import (
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	scanner *Scanner
	current *Token
	hold    bool
	err     error
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
	if err := p.read(); err != nil {
		return nil, err
	}

	switch p.current.Type {
	case String:
		lit := StringLiteral(p.current.Val.(string))
		p.discard()
		return &lit, nil
	case Ident:
		left := &Identifier{
			Line:   p.current.Line,
			Offset: p.current.Offset,
			Name:   p.current.Val.(string),
		}
		p.discard()

		if err := p.read(); err != nil || p.current.Type != Equals {
			return left, nil
		}
		p.discard()

		right, err := p.ParseExpression()
		if err != nil {
			return nil, err
		}
		p.discard()

		return &Assignment{
			Left:  left,
			Right: right,
		}, nil
	default:
		return nil, &ParseError{Actual: p.current}
	}
}

func (p *Parser) read() error {
	if p.err != nil || p.hold {
		return p.err
	}

	p.current, p.err = p.scanner.Next()
	p.hold = true
	return p.err
}

func (p *Parser) discard() {
	p.hold = false
}
