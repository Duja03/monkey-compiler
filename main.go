package main

import (
	"fmt"
	"monkey-interpreter/repl"
	"os"
)

func main() {
	fmt.Println("This is the Monkey programming language!")
	fmt.Print("Feel free to type in commands\n\n")

	repl.Start(os.Stdin, os.Stdout)
}
