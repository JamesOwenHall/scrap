package parser

import (
	"bufio"
	"errors"
	"io"
)

// ErrUnexpectedInput represents any kind of unexpected input in the reader.
var ErrUnexpectedInput = errors.New("unexpected input")

// Scanner reads tokens from an io.Reader.
type Scanner struct {
	reader  *bufio.Reader
	buf     []rune
	current rune
	hold    bool
	err     error
}

// NewScanner returns a new scanner over the reader r.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		reader: bufio.NewReader(r),
		buf:    make([]rune, 0, 1024),
	}
}

// Next returns the next token from the reader, if any.  Returns EOF when there
// are no more to read.
func (s *Scanner) Next() (*Token, error) {
	// Clear the buffer.
	s.clear()

	if err := s.read(); err != nil {
		return nil, err
	}

	// Skip the spaces.
	for isSpace(s.current) {
		s.discardCurrent()
		if err := s.read(); err != nil {
			return nil, err
		}
	}

	// Figure out how to parse the next token.
	switch {
	case isAlpha(s.current):
		return s.readIdentifier(), nil
	case s.current == '"':
		return s.readString()
	case s.current == '(':
		return s.readSingle(OpenParen), nil
	case s.current == ')':
		return s.readSingle(CloseParen), nil
	case s.current == '=':
		return s.readSingle(Equals), nil
	default:
		return nil, ErrUnexpectedInput
	}
}

func (s *Scanner) readIdentifier() *Token {
	s.appendCurrent()

	for {
		if err := s.read(); err != nil {
			break
		}

		if !isAlpha(s.current) && !isDigit(s.current) {
			break
		}

		s.appendCurrent()
	}

	return &Token{
		Type: Identifier,
		Val:  string(s.buf),
	}
}

func (s *Scanner) readString() (*Token, error) {
	s.discardCurrent()

	for {
		if err := s.read(); err != nil {
			return nil, err
		}

		if s.current == '"' {
			s.discardCurrent()
			break
		}

		if s.current == '\\' {
			s.discardCurrent()
			if err := s.read(); err != nil {
				return nil, err
			}

			if s.current != '\\' && s.current != '"' {
				return nil, ErrUnexpectedInput
			}
		}

		s.appendCurrent()
	}

	return &Token{
		Type: String,
		Val:  string(s.buf),
	}, nil
}

func (s *Scanner) readSingle(typ TokenType) *Token {
	s.appendCurrent()
	return &Token{
		Type: typ,
		Val:  string(s.buf),
	}
}

func (s *Scanner) read() error {
	if s.err != nil || s.hold {
		return s.err
	}

	s.current, _, s.err = s.reader.ReadRune()
	s.hold = true
	return s.err
}

func (s *Scanner) appendCurrent() {
	s.buf = append(s.buf, s.current)
	s.hold = false
}

func (s *Scanner) discardCurrent() {
	s.hold = false
}

func (s *Scanner) clear() {
	s.buf = s.buf[:0]
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
