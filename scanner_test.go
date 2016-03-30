package scrap

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	reader := strings.NewReader(`x=foo("bar",.)`)
	scanner := NewScanner(reader)

	expected := []Token{
		Token{Line: 0, Offset: 0, Type: Ident, Val: "x"},
		Token{Line: 0, Offset: 1, Type: Equals, Val: "="},
		Token{Line: 0, Offset: 2, Type: Ident, Val: "foo"},
		Token{Line: 0, Offset: 5, Type: OpenParen, Val: "("},
		Token{Line: 0, Offset: 6, Type: String, Val: "bar"},
		Token{Line: 0, Offset: 11, Type: Comma, Val: ","},
		Token{Line: 0, Offset: 12, Type: Period, Val: "."},
		Token{Line: 0, Offset: 13, Type: CloseParen, Val: ")"},
	}

	for _, exp := range expected {
		if next, err := scanner.Next(); err != nil {
			t.Fatal(err)
		} else if !reflect.DeepEqual(&exp, next) {
			t.Fatal(next)
		}
	}
}

func TestScannerIdent(t *testing.T) {
	reader := strings.NewReader(` foo bar `)
	scanner := NewScanner(reader)

	if next, err := scanner.Next(); err != nil {
		t.Fatal(err)
	} else if next.Type != Ident {
		t.Fatal(next.Type)
	} else if next.Val != "foo" {
		t.Fatal(next.Val)
	}

	if next, err := scanner.Next(); err != nil {
		t.Fatal(err)
	} else if next.Type != Ident {
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
	reader := strings.NewReader(` (,.)= `)
	scanner := NewScanner(reader)

	expected := []TokenType{OpenParen, Comma, Period, CloseParen, Equals}
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
