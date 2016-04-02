package scrap

import (
	"fmt"
	"strings"
)

type UnknownVariable struct {
	Line, Offset int
	Name         string
}

func (u *UnknownVariable) Error() string {
	return fmt.Sprintf("(%d,%d) unknown variable \"%s\"", u.Line, u.Offset, u.Name)
}

type UnknownFunction struct {
	Line, Offset int
	Name         string
}

func (u *UnknownFunction) Error() string {
	return fmt.Sprintf("(%d,%d) unknown function \"%s\"", u.Line, u.Offset, u.Name)
}

type Expression interface {
	Eval(p *Program) (interface{}, error)
	String() string
}

type Identifier struct {
	Line, Offset int
	Name         string
}

func (i *Identifier) Eval(p *Program) (interface{}, error) {
	if val, exists := p.Vars[i.Name]; !exists {
		return nil, &UnknownVariable{Line: i.Line, Offset: i.Offset, Name: i.Name}
	} else {
		return val, nil
	}
}

func (i *Identifier) String() string {
	return i.Name
}

type StringLiteral string

func (s *StringLiteral) Eval(_ *Program) (interface{}, error) {
	return string(*s), nil
}

func (s *StringLiteral) String() string {
	return `"` + string(*s) + `"`
}

type Assignment struct {
	Left  *Identifier
	Right Expression
}

func (a *Assignment) Eval(p *Program) (interface{}, error) {
	result, err := a.Right.Eval(p)
	if err != nil {
		return nil, err
	}

	p.Vars[a.Left.Name] = result
	return result, nil
}

func (a *Assignment) String() string {
	return fmt.Sprintf("%s = %s", a.Left.String(), a.Right.String())
}

type FunctionCall struct {
	Line, Offset int
	Name         string
	Arguments    []Expression
}

func (f *FunctionCall) Eval(p *Program) (interface{}, error) {
	if fn, exists := p.Funcs[f.Name]; !exists {
		return nil, &UnknownFunction{Line: f.Line, Offset: f.Offset, Name: f.Name}
	} else {
		values := make([]interface{}, 0, len(f.Arguments))
		for _, arg := range f.Arguments {
			if val, err := arg.Eval(p); err != nil {
				return nil, err
			} else {
				values = append(values, val)
			}
		}

		return fn(values)
	}
}

func (f *FunctionCall) String() string {
	args := make([]string, 0, len(f.Arguments))
	for _, arg := range f.Arguments {
		args = append(args, arg.String())
	}

	return fmt.Sprintf("%s(%s)", strings.Join(args, ", "))
}
