package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionDef(t *testing.T) {

	program := `
	def square {
    	return _.x * _.x;
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
            "argument": {
              "left": {
                "computed": false,
                "object": {
                  "name": "_",
                  "type": "Identifier"
                },
                "property": {
                  "name": "x",
                  "type": "Identifier"
                },
                "type": "MemberExpression"
              },
              "operator": "*",
              "right": {
                "computed": false,
                "object": {
                  "name": "_",
                  "type": "Identifier"
                },
                "property": {
                  "name": "x",
                  "type": "Identifier"
                },
                "type": "MemberExpression"
              },
              "type": "BinaryExpression"
            },
            "type": "ReturnStatement"
          }
        ],
        "type": "BlockStatement"
      },
      "name": {
        "name": "square",
        "type": "Identifier"
      },
      "type": "FunctionDeclaration"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
