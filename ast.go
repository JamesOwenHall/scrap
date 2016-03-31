package scrap

import (
	"fmt"
)

type UnknownVariable struct {
	Line, Offset int
	Name         string
}

func (u *UnknownVariable) Error() string {
	return fmt.Sprintf("(%d,%d) unknown variable \"%d\"", u.Line, u.Offset, u.Name)
}

type Expression interface {
	Eval(p *Program) (interface{}, error)
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

type StringLiteral string

func (s *StringLiteral) Eval(_ *Program) (interface{}, error) {
	return string(*s), nil
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

type FunctionCall struct {
	Line, Offset int
	Name         string
	Arguments    []Expression
}

func (f *FunctionCall) Eval(p *Program) (interface{}, error) {
	panic("TODO")
}
