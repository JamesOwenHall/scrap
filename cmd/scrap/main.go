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

	files := make([]io.Reader, 0, flag.NArg())

	// Parse all of the input files.
	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			Errorf("error: can't open file %s", flag.Arg(i))
			os.Exit(1)
		}
		defer file.Close()

		files = append(files, file)
	}

	expressions := []scrap.Expression{}
	parser := scrap.NewParser(io.MultiReader(files...))
	for {
		expr, err := parser.ParseExpression()
		if err == io.EOF {
			break
		} else if err != nil {
			Errorf("error: %s", err.Error())
			os.Exit(1)
		}

		expressions = append(expressions, expr)
	}

	// Run all of the expressions.
	program := scrap.NewProgram()
	for _, expr := range expressions {
		_, err := expr.Eval(program)
		if err != nil {
			Errorf("%s", err.Error())
			os.Exit(1)
		}
	}
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, color.RedString(format, args...))
}
