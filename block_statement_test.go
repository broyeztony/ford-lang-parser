package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	{
	  "body": [
	    {
	      "body": [
	        {
	          "expression": {
	            "type": "NumericLiteral",
	            "value": 42
	          },
	          "type": "ExpressionStatement"
	        },
	        {
	          "expression": {
	            "type": "StringLiteral",
	            "value": "Hello"
	          },
	          "type": "ExpressionStatement"
	        }
	      ],
	      "type": "BlockStatement"
	    }
	  ],
	  "type": "Program"
	}
*/
func TestBlockStatement(t *testing.T) {

	program := `
	{
		42; 
		"Hello";
	}
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "body": [
        {
          "expression": {
            "type": "NumericLiteral",
            "value": 42
          },
          "type": "ExpressionStatement"
        },
        {
          "expression": {
            "type": "StringLiteral",
            "value": "Hello"
          },
          "type": "ExpressionStatement"
        }
      ],
      "type": "BlockStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
