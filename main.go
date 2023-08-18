package main

func main() {
	program := `
		x;
		y;
	`

	parser := NewParser(program)
	parser.parse()
}
