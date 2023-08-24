package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoop(t *testing.T) {

	program := `
	do {
		x -= 1;
	} while (x > 10);
	`

	parser := NewParser(program)
	ast := parser.parse()

	actual := encode(ast)

	expected := `{
  "body": [
    {
      "body": {
        "body": [
          {
            "expression": {
              "left": {
                "name": "x",
                "type": "Identifier"
              },
              "operator": "-=",
              "right": {
                "type": "NumericLiteral",
                "value": "1"
              },
              "type": "AssignmentExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "test": {
        "left": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": ">",
        "right": {
          "type": "NumericLiteral",
          "value": "10"
        },
        "type": "BinaryExpression"
      },
      "type": "DoWhileStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
