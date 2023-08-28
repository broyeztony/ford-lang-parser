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

func TestObjectLiteral(t *testing.T) {

	program := `
	let b = { x: a };
	`

	parser := NewParser(program)
	ast := parser.parse()
	actual := encode(ast)
	expected := `{
  "body": [
    {
      "declarations": [
        {
          "id": {
            "name": "b",
            "type": "Identifier"
          },
          "initializer": {
            "type": "ObjectLiteral",
            "values": [
              {
                "name": "x",
                "value": {
                  "name": "a",
                  "type": "Identifier"
                }
              }
            ]
          },
          "type": "VariableDeclaration"
        }
      ],
      "type": "VariableStatement"
    }
  ],
  "type": "Program"
}`

	assert.Equal(t, expected, actual)
}
