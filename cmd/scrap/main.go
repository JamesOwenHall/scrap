package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JamesOwenHall/scrap"

	"github.com/fatih/color"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		Errorf("error: missing input file")
		os.Exit(1)
	}

	program := scrap.NewProgram()
	for _, arg := range flag.Args() {
		if err := program.RunFile(arg); err != nil {
			Errorf("%s", err.Error())
			os.Exit(1)
		}
	}
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, color.RedString(format, args...))
}
