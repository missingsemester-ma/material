package main

import "strconv"

type TokenType int

const (
	KEYWORD    TokenType = iota // A keyword (if, else, for, while, return, ...)
	IDENTIFIER                  // An identifier variable
	LITERAL                     // A LITERAL number. Would be nice to support floating numbers
	OPERATOR                    // An operator +*-/&|^
	ASSIGN                      // Assignment
	LPAR                        // Left parenthesis
	RPAR                        // Right parenthesis
	EOF                         // End of file
)

func (t TokenType) String() string {
	switch t {
	case KEYWORD:
		return "KEYWORD"
	case IDENTIFIER:
		return "IDENTIFIER"
	case LITERAL:
		return "LITERAL"
	case OPERATOR:
		return "OPERATOR"
	case ASSIGN:
		return "ASSIGN"
	case LPAR:
		return "LPAR"
	case RPAR:
		return "RPAR"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	Type  TokenType
	Value string
	Row   int
	Col   int
}

func createToken(value string, row, col int) Token {
	tokenType := IDENTIFIER
	if value == "if" || value == "then" || value == "else" || value == "return" {
		tokenType = KEYWORD
	} else if value == "(" {
		tokenType = LPAR
	} else if value == ")" {
		tokenType = RPAR
	} else if value == "=" {
		tokenType = ASSIGN
	} else if value == "==" || value == "<" || value == ">" {
		tokenType = OPERATOR
	} else if value == "+" || value == "-" || value == "*" || value == "/" {
		tokenType = OPERATOR
	} else if isNumber(value) {
		tokenType = LITERAL
	}
	return Token{Type: tokenType, Value: value, Row: row, Col: col}
}

func isNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
