package main

import (
	"fmt"
	"strings"
)

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(input string) *Parser {

	parser := &Parser{}
	parser.tokenizer = NewTokenizer(strings.TrimSpace(input))

	return parser
}

func (p *Parser) parse() {
	fmt.Println("@ program", p.tokenizer.input, "\n")

	for p.tokenizer.hasMoreTokens() {
		token := p.tokenizer.getNextToken()

		if token != nil {
			fmt.Println(token.toString())
		}
	}
}
