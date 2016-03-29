package parser

import (
	"io"
	"strings"
	"testing"
)

func TestScannerIdentifier(t *testing.T) {
	reader := strings.NewReader(" foo bar ")
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
	reader := strings.NewReader(" ()= ")
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
