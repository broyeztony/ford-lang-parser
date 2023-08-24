package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquality(t *testing.T) {

	program := `
	x > 0 == true;
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
            "value": "0"
          },
          "type": "BinaryExpression"
        },
        "operator": "==",
        "right": {
          "type": "BooleanLiteral",
          "value": true
        },
        "type": "BinaryExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}