package parser

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	reader := strings.NewReader(`x=foo("bar")`)
	scanner := NewScanner(reader)

	expected := []Token{
		Token{Type: Identifier, Val: "x"},
		Token{Type: Equals, Val: "="},
		Token{Type: Identifier, Val: "foo"},
		Token{Type: OpenParen, Val: "("},
		Token{Type: String, Val: "bar"},
		Token{Type: CloseParen, Val: ")"},
	}

	for _, exp := range expected {
		if next, err := scanner.Next(); err != nil {
			t.Fatal(err)
		} else if !reflect.DeepEqual(&exp, next) {
			t.Fatal(next.Type)
		}
	}
}

func TestScannerIdentifier(t *testing.T) {
	reader := strings.NewReader(` foo bar `)
	scanner := NewScanner(reader)

	if next, err := scanner.Next(); err != nil {
		t.Fatal(err)
	} else if next.Type != Identifier {
		t.Fatal(next.Type)
	} else if next.Val != "foo" {
		t.Fatal(next.Val)
	}

	if next, err := scanner.Next(); err != nil {
		t.Fatal(err)
	} else if next.Type != Identifier {
		t.Fatal(next.Type)
	} else if next.Val != "bar" {
		t.Fatal(next.Val)
	}

	if next, err := scanner.Next(); err != io.EOF {
		t.Fatal(next, err)
	}
}

func TestScannerString(t *testing.T) {
	reader := strings.NewReader(` "foo \"\\ bar" `)
	scanner := NewScanner(reader)

	if next, err := scanner.Next(); err != nil {
		t.Fatal(err)
	} else if next.Type != String {
		t.Fatal(next.Type)
	} else if next.Val != `foo "\ bar` {
		t.Fatal(next.Val)
	}

	if next, err := scanner.Next(); err != io.EOF {
		t.Fatal(next, err)
	}
}

func TestScannerPunctuation(t *testing.T) {
	reader := strings.NewReader(` ()= `)
	scanner := NewScanner(reader)

	expected := []TokenType{OpenParen, CloseParen, Equals}
	for _, exp := range expected {
		if next, err := scanner.Next(); err != nil {
			t.Fatal(err)
		} else if next.Type != exp {
			t.Fatal(next.Type)
		}
	}

	if next, err := scanner.Next(); err != io.EOF {
		t.Fatal(next, err)
	}
}
