package scrap

import (
	"fmt"
	"io"
	"os"
)

type Program struct {
	Vars  map[string]interface{}
	Funcs map[string]func([]interface{}) (interface{}, error)
}

func NewProgram() *Program {
	return &Program{
		Vars: make(map[string]interface{}),
		Funcs: map[string]func([]interface{}) (interface{}, error){
			"print": print,
		},
	}
}

func (p *Program) RunFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse all expressions.
	parser := NewParser(file)
	expressions := []Expression{}
	for {
		expr, err := parser.ParseExpression()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		expressions = append(expressions, expr)
	}

	// Run all of the expressions.
	program := NewProgram()
	for _, expr := range expressions {
		_, err := expr.Eval(program)
		if err != nil {
			return err
		}
	}

	return nil
}

func print(args []interface{}) (interface{}, error) {
	str := fmt.Sprint(args...)
	fmt.Println(str)
	return str, nil
}
