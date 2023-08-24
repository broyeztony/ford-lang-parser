package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
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
	            "value": "42"
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

	expected := "{\"body\":[{\"body\":[{\"expression\":{\"type\":\"NumericLiteral\",\"value\":\"42\"},\"type\":\"ExpressionStatement\"},{\"expression\":{\"type\":\"StringLiteral\",\"value\":\"Hello\"},\"type\":\"ExpressionStatement\"}],\"type\":\"BlockStatement\"}],\"type\":\"Program\"}"

	assert.Equal(t, expected, actual)
}

func encode(ast interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(ast)
	jsonString := buffer.String()
	return strings.TrimRight(jsonString, "\n")
}
