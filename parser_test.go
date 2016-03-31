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

func TestParserFunctionCall(t *testing.T) {
	expr, err := ParseString(`foo(bar, "baz")`)
	if err != nil {
		t.Fatal(err)
	}

	if fc, ok := expr.(*FunctionCall); !ok {
		t.Fatalf("%T", expr)
	} else if fc.Name != "foo" {
		t.Fatal(fc.Name)
	} else if len(fc.Arguments) != 2 {
		t.Fatal(fc.Arguments)
	} else if ident, ok := fc.Arguments[0].(*Identifier); !ok || ident.Name != "bar" {
		t.Fatal(fc.Arguments[0])
	} else if str, ok := fc.Arguments[1].(*StringLiteral); !ok || *str != "baz" {
		t.Fatal(*str)
	}
}
