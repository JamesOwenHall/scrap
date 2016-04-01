package scrap

type Program struct {
	Vars  map[string]interface{}
	Funcs map[string]func([]interface{}) (interface{}, error)
}

func NewProgram() *Program {
	return &Program{
		Vars:  make(map[string]interface{}),
		Funcs: make(map[string]func([]interface{}) (interface{}, error)),
	}
}
