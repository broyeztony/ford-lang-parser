package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStatement(t *testing.T) {

	program := `
	;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "type": "EmptyStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
