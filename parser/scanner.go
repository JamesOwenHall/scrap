package parser

import (
	"bufio"
	"errors"
	"io"
)

var ErrUnexpectedInput = errors.New("unexpected input")

type Scanner struct {
	reader  *bufio.Reader
	buf     []rune
	current rune
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		reader: bufio.NewReader(r),
		buf:    make([]rune, 0, 1024),
	}
}

func (s *Scanner) Next() (*Token, error) {
	s.buf = s.buf[:0]

	var err error
	s.current, _, err = s.reader.ReadRune()
	if err != nil {
		return nil, err
	}

	for isSpace(s.current) {
		s.current, _, err = s.reader.ReadRune()
		if err != nil {
			return nil, err
		}
	}

	switch {
	case isAlpha(s.current):
		return s.readIdentifier(), nil
	default:
		return nil, ErrUnexpectedInput
	}
}

func (s *Scanner) readIdentifier() *Token {
	s.buf = append(s.buf, s.current)

	var err error
	for {
		s.current, _, err = s.reader.ReadRune()
		if err != nil {
			break
		}

		if !isAlpha(s.current) && !isDigit(s.current) {
			break
		}

		s.buf = append(s.buf, s.current)
	}

	return &Token{
		Type: Identifier,
		Val:  string(s.buf),
	}
}

func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || r == '_'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}
