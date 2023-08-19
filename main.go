package main

func main() {
	program := `
		import ("abc", "def");

		let x = 3;
		let y = -2.76437842 + x;
		
		def myFunc {
			{ x };
		}
	`

	parser := NewParser(program)
	parser.parse()
}
