package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericLiteral(t *testing.T) {

	program := `
	42;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
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

func TestFloatLiteral(t *testing.T) {

	program := `
	42.34;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "type": "NumericLiteral",
        "value": 42.34
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}

func TestStringLiteral(t *testing.T) {

	program := `
	"42";
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "type": "StringLiteral",
        "value": "42"
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}

func TestSingleQuoteStringLiteral(t *testing.T) {

	program := `
	'"42"';
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "type": "StringLiteral",
        "value": "\"42\""
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}

func TestBooleanLiteral(t *testing.T) {

	program := `
	true;
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "expression": {
        "type": "BooleanLiteral",
        "value": true
      },
      "type": "ExpressionStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
