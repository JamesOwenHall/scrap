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

	if next, err := scanner.Next(); err == nil {
		t.Fatal(next)
	} else if err != io.EOF {
		t.Fatal(err)
	}
}
