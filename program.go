package scrap

type Program struct {
	Vars map[string]interface{}
}

func NewProgram() *Program {
	return &Program{
		Vars: make(map[string]interface{}),
	}
}
