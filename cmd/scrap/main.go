package main

import (
	"flag"
	"fmt"
	"io"
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
	expressions := []scrap.Expression{}

	// Parse all of the input files.
	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			Errorf("error: can't open file %s", flag.Arg(i))
			os.Exit(1)
		}

		for {
			expr, err := scrap.Parse(file)
			if err == io.EOF {
				break
			} else if err != nil {
				Errorf("error: %s", err.Error())
				file.Close()
				os.Exit(1)
			}

			expressions = append(expressions, expr)
		}

		file.Close()
	}

	// Run all of the expressions.
	for _, expr := range expressions {
		_, err := expr.Eval(program)
		if err != nil {
			Errorf("error: %s", err.Error())
			os.Exit(1)
		}
	}
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, color.RedString(format, args...))
}
