package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("playground.ford")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	program := string(data)

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
