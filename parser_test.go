package scrap

import (
	"testing"
)

func TestParserLiteral(t *testing.T) {
	expr, err := ParseString(`"bar"`)
	if err != nil {
		t.Fatal(err)
	}

	if lit, ok := expr.(*StringLiteral); !ok {
		t.Fatalf("%T", expr)
	} else if *lit != "bar" {
		t.Fatal(*lit)
	}
}

func TestParserAssignment(t *testing.T) {
	expr, err := ParseString(`foo = "bar"`)
	if err != nil {
		t.Fatal(err)
	}

	if assign, ok := expr.(*Assignment); !ok {
		t.Fatalf("%T", expr)
	} else if assign.Left.Name != "foo" {
		t.Fatal(assign.Left)
	} else if str, ok := assign.Right.(*StringLiteral); !ok {
		t.Fatalf("%T", assign.Right)
	} else if *str != "bar" {
		t.Fatal(*str)
	}
}
