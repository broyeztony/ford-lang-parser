package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	program := `
	def square {
    	return _.x * _.x;
	}
	`

	fmt.Println(program)
	fmt.Println("----------------------")
	fmt.Println()

	parser := NewParser(program)
	ast := parser.parse()

	fmt.Println(encode(ast))
}

func encode(ast interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	encoder.Encode(ast)
	jsonString := buffer.String()
	return strings.TrimRight(jsonString, "\n")
}
