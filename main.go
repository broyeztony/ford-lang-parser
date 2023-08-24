package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	program := `
	{
		42; 
		"Hello";
	}
	`

	fmt.Println(program)
	fmt.Println("----------------------")
	fmt.Println()

	parser := NewParser(program)
	ast := parser.parse()

	fmt.Println(ast)

	j, _ := json.MarshalIndent(ast, "", "  ")
	fmt.Println(string(j))
}
