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
	}
*/
func TestEquality(t *testing.T) {

	program := `
	x > 0 == true;
	`

	parser := NewParser(program)
	ast := parser.parse()

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(ast)
	jsonString := buffer.String()
	actual := strings.TrimRight(jsonString, "\n")

	expected := "{\"body\":[{\"expression\":{\"left\":{\"left\":{\"name\":\"x\",\"type\":\"Identifier\"},\"operator\":\">\",\"right\":{\"type\":\"NumericLiteral\",\"value\":\"0\"},\"type\":\"BinaryExpression\"},\"operator\":\"==\",\"right\":{\"type\":\"BooleanLiteral\",\"value\":true},\"type\":\"BinaryExpression\"},\"type\":\"ExpressionStatement\"}],\"type\":\"Program\"}"
	t.Log("@ expected: ", expected)

	assert.Equal(t, expected, actual)
}
