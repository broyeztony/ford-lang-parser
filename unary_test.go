package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnaryExpression(t *testing.T) {

	program := `
	-x;
	!x;
	`
	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "argument": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": "-",
        "type": "UnaryExpression"
      },
      "type": "ExpressionStatement"
    },
    {
      "expression": {
        "argument": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": "!",
        "type": "UnaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
