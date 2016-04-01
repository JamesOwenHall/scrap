package scrap

import (
	"fmt"
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

func print(args []interface{}) (interface{}, error) {
	fmt.Println(args...)
	return nil, nil
}
