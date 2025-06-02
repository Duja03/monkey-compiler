# Monkey Interpreter in Go

This is my implementation of the **Monkey** programming language interpreter, written in Go.  
The project is based on the books [**"Writing An Interpreter In Go"**](https://interpreterbook.com) and [**"Writing A Compiler In Go"**](https://compilerbook.com
) by Thorsten Ball.

It was built entirely from scratch, covering everything from lexical analysis to evaluation.  
The goal was to understand how interpreters work at a fundamental level.

## Features

- *Lexer* (tokenizing source code)
- *Pratt Parser* (generating AST)
- *Evaluator* (interpreter engine)
- First-class functions and closures
- Lexical scope
- Built-in types: integers, booleans, strings, arrays, hashes
- Built-in functions: `len`, `first`, `last`, `rest`, `push`, etc.
- Error handling and environment isolation
- *REPL* (interactive prompt)

## Getting Started

### Requirements

- Go 1.18 or higher (tested with 1.21)

### Install and Run

```bash
git clone https://github.com/your-username/monkey-interpreter.git
cd monkey-interpreter
go run main.go
```
