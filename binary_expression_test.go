package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleBinaryExpression(t *testing.T) {

	program := `
	2 + 2;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "+",
        "right": {
          "type": "NumericLiteral",
          "value": 2
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

func TestMulBinaryExpression(t *testing.T) {

	program := `
	2 + 2 * 2;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "+",
        "right": {
          "left": {
            "type": "NumericLiteral",
            "value": 2
          },
          "operator": "*",
          "right": {
            "type": "NumericLiteral",
            "value": 2
          },
          "type": "BinaryExpression"
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

func TestParenthesizedBinaryExpression(t *testing.T) {

	program := `
	2 * (3 + 6);
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "left": {
          "type": "NumericLiteral",
          "value": 2
        },
        "operator": "*",
        "right": {
          "left": {
            "type": "NumericLiteral",
            "value": 3
          },
          "operator": "+",
          "right": {
            "type": "NumericLiteral",
            "value": 6
          },
          "type": "BinaryExpression"
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
