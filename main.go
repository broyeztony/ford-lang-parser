package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	program := `
	let x = { a };
	`

	fmt.Println(program)
	fmt.Println("----------------------\n")

	parser := NewParser(program)
	ast := parser.parse()

	j, _ := json.MarshalIndent(ast, "", "  ")
	fmt.Println(string(j))
}
