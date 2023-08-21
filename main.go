package main

import "fmt"

func main() {
	program := `
		def keywork {
			print("hello", 4);
			return true;
		}

		keywork();
	`

	fmt.Println("@ program", program, "\n")

	parser := NewParser(program)
	ast := parser.parse()

	fmt.Println(ast)
}
