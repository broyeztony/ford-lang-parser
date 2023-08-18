package main

import "fmt"

type Parser struct {
	input     string
	tokenizer *Tokenizer
}

func NewParser(input string) *Parser {
	parser := &Parser{
		input: input,
	}

	parser.tokenizer = NewTokenizer(input)

	return parser
}

func (p *Parser) parse() {
	fmt.Println(p.input)
}
