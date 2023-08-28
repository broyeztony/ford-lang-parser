package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssign1(t *testing.T) {

	program := `
	x = 42;
	`

	parser := NewParser(program)
	ast := parser.parse()

	got, _ := json.MarshalIndent(ast, "", "  ")
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": "=",
        "right": {
          "type": "NumericLiteral",
          "value": 42
        },
        "type": "AssignmentExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`
	assert.Equal(t, expected, string(got))
}

func TestAssign2(t *testing.T) {

	program := `
	x = y = 42;
	`

	parser := NewParser(program)
	ast := parser.parse()

	got, _ := json.MarshalIndent(ast, "", "  ")
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "name": "x",
          "type": "Identifier"
        },
        "operator": "=",
        "right": {
          "left": {
            "name": "y",
            "type": "Identifier"
          },
          "operator": "=",
          "right": {
            "type": "NumericLiteral",
            "value": 42
          },
          "type": "AssignmentExpression"
        },
        "type": "AssignmentExpression"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`
	assert.Equal(t, expected, string(got))
}
