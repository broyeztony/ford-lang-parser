package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogicalExpression(t *testing.T) {

	program := `
	x > 0 && y < 1;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "left": {
            "name": "x",
            "type": "Identifier"
          },
          "operator": ">",
          "right": {
            "type": "NumericLiteral",
            "value": 0
          },
          "type": "BinaryExpression"
        },
        "operator": "&&",
        "right": {
          "left": {
            "name": "y",
            "type": "Identifier"
          },
          "operator": "<",
          "right": {
            "type": "NumericLiteral",
            "value": 1
          },
          "type": "BinaryExpression"
        },
        "type": "LogicalExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
