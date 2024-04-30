package main

import ()

type TokenType int

const (
	KEYWORD    = iota // A keyword (if, else, for, while, return, ...)
	IDENTIFIER        // An identifier variable
	LITERAL           // A LITERAL number. Would be nice to support floating numbers
	OPERATOR          // An operator +*-/&|^
	ASSIGN            // Assignment
	LPAR              // Left parenthesis
	RPAR              // Right parenthesis
)

type Token struct {
	Type  TokenType // The Token type
	Value string    // The lexem/value of this token
	Row   int       // The row in which this token appears
	Col   int       // The column in which this token appears
}

func tokenize(input string) ([]Token, error) {
	// Implement me
	return nil, nil
}
