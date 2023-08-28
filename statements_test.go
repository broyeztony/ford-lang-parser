package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatementList(t *testing.T) {

	program := `
	print("hello Ford!");
	42;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "arguments": [
          {
            "type": "StringLiteral",
            "value": "hello Ford!"
          }
        ],
        "callee": {
          "name": "print",
          "type": "Identifier"
        },
        "type": "CallExpression"
      },
      "type": "ExpressionStatement"
    },
    {
      "expression": {
        "type": "NumericLiteral",
        "value": 42
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
