package main

import (
	"bufio"
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
		Repl()
		return
	}

	program := scrap.NewProgram()
	for _, arg := range flag.Args() {
		if err := program.RunFile(arg); err != nil {
			Errorf("%s", err.Error())
			os.Exit(1)
		}
	}
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	program := scrap.NewProgram()

	for {
		fmt.Print("=> ")
		if !scanner.Scan() {
			break
		}

		expr, err := scrap.ParseString(scanner.Text())
		if err == io.EOF {
			continue
		} else if err != nil {
			Errorf("%s", err.Error())
			continue
		}

		val, err := expr.Eval(program)
		if err != nil {
			Errorf("%s", err.Error())
			continue
		}

		fmt.Println(val)
	}

	if err := scanner.Err(); err != nil {
		Errorf("%s", err.Error())
		os.Exit(1)
	}
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, color.RedString(format, args...))
}
