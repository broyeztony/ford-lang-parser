package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFor(t *testing.T) {

	program := `
	for (let i = 0 ; i < 10 ; i += 1) {
    	x += i;
	}
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
              "operator": "+=",
              "right": {
                "name": "i",
                "type": "Identifier"
              },
              "type": "AssignmentExpression"
            },
            "type": "ExpressionStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "init": {
        "declarations": [
          {
            "id": {
              "name": "i",
              "type": "Identifier"
            },
            "initializer": {
              "type": "NumericLiteral",
              "value": 0
            },
            "type": "VariableDeclaration"
          }
        ],
        "type": "VariableStatement"
      },
      "test": {
        "left": {
          "name": "i",
          "type": "Identifier"
        },
        "operator": "<",
        "right": {
          "type": "NumericLiteral",
          "value": 10
        },
        "type": "BinaryExpression"
      },
      "type": "ForStatement",
      "update": {
        "left": {
          "name": "i",
          "type": "Identifier"
        },
        "operator": "+=",
        "right": {
          "type": "NumericLiteral",
          "value": 1
        },
        "type": "AssignmentExpression"
      }
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
