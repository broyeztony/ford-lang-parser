package main

type Rule struct {
	pattern   string
	tokenType string
}

type Tokenizer struct {
	input  string
	spec   []Rule
	cursor int
}

func NewTokenizer(input string) *Tokenizer {

	tokenizer := &Tokenizer{
		input:  input,
		cursor: 0,
	}

	// Tokenizer rules
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\s+`, tokenType: ""})              // whitespace
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\/\/.*`, tokenType: ""})           // single-line comment
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\/\*[\s\S]*?\*\/`, tokenType: ""}) // multi-line comments
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^;`, tokenType: ";"})               // delimiter
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^:`, tokenType: ":"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\{`, tokenType: "{"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\}`, tokenType: "}"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\(`, tokenType: "("})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\)`, tokenType: ")"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^,`, tokenType: ","})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\.`, tokenType: "."})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\[`, tokenType: "["})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\]`, tokenType: "]"})

	// ------------------------------------------------------------------------------------- KEYWORDS
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\blet\b`, tokenType: "let"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bif\b`, tokenType: "if"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\belse\b`, tokenType: "else"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\btrue\b`, tokenType: "true"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bfalse\b`, tokenType: "false"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bnull\b`, tokenType: "null"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bwhile\b`, tokenType: "while"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bdo\b`, tokenType: "do"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bfor\b`, tokenType: "for"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\bdef\b`, tokenType: "def"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\breturn\b`, tokenType: "return"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\brecover\b`, tokenType: "recover"})

	// ------------------------------------------------------------------------------------- NUMBERS
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)`, tokenType: "NUMBER"})

	// ------------------------------------------------------------------------------------- IDENTIFIER, OPERATORS
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\w+`, tokenType: "IDENTIFIER"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^[=!]=`, tokenType: "EQUALITY_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^->`, tokenType: "ERROR_HANDLER_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^=`, tokenType: "SIMPLE_ASSIGN"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^[\*\/\+\-]=`, tokenType: "COMPLEX_ASSIGN"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^[+\-]`, tokenType: "ADDITIVE_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^[*\/]`, tokenType: "MULTIPLICATIVE_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^[><]=?`, tokenType: "RELATIONAL_OPERATOR"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^&&`, tokenType: "LOGICAL_AND"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^\|\|`, tokenType: "LOGICAL_OR"})
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^!`, tokenType: "LOGICAL_NOT"})

	// ------------------------------------------------------------------------------------- STRINGS
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^"[^"]*"`, tokenType: "STRING"}) // single quotes
	tokenizer.spec = append(tokenizer.spec, Rule{pattern: `^'[^']*'`, tokenType: "STRING"}) // double quotes

	return tokenizer
}
