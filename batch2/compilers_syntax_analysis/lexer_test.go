package main

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "if x == 5 then y = 10 else y = 20",
			expected: []Token{
				{Type: KEYWORD, Value: "if", Row: 1, Col: 1},
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 4},
				{Type: OPERATOR, Value: "==", Row: 1, Col: 6},
				{Type: LITERAL, Value: "5", Row: 1, Col: 9},
				{Type: KEYWORD, Value: "then", Row: 1, Col: 11},
				{Type: IDENTIFIER, Value: "y", Row: 1, Col: 16},
				{Type: ASSIGN, Value: "=", Row: 1, Col: 18},
				{Type: LITERAL, Value: "10", Row: 1, Col: 20},
				{Type: KEYWORD, Value: "else", Row: 1, Col: 23},
				{Type: IDENTIFIER, Value: "y", Row: 1, Col: 28},
				{Type: ASSIGN, Value: "=", Row: 1, Col: 30},
				{Type: LITERAL, Value: "20", Row: 1, Col: 32},
				{Type: EOF, Value: "", Row: 1, Col: 34},
			},
		},
		{
			input: "x = (y + 5) * 2",
			expected: []Token{
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 1},
				{Type: ASSIGN, Value: "=", Row: 1, Col: 3},
				{Type: LPAR, Value: "(", Row: 1, Col: 5},
				{Type: IDENTIFIER, Value: "y", Row: 1, Col: 6},
				{Type: OPERATOR, Value: "+", Row: 1, Col: 8},
				{Type: LITERAL, Value: "5", Row: 1, Col: 10},
				{Type: RPAR, Value: ")", Row: 1, Col: 11},
				{Type: OPERATOR, Value: "*", Row: 1, Col: 13},
				{Type: LITERAL, Value: "2", Row: 1, Col: 15},
				{Type: EOF, Value: "", Row: 1, Col: 16},
			},
		},
		{
			input: "x = 10 \\ This is a comment",
			expected: []Token{
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 1},
				{Type: ASSIGN, Value: "=", Row: 1, Col: 3},
				{Type: LITERAL, Value: "10", Row: 1, Col: 5},
				{Type: EOF, Value: "", Row: 1, Col: 28},
			},
		},
		{
			input: "/* Multi-line\ncomment */\nx = 5",
			expected: []Token{
				{Type: IDENTIFIER, Value: "x", Row: 3, Col: 1},
				{Type: ASSIGN, Value: "=", Row: 3, Col: 3},
				{Type: LITERAL, Value: "5", Row: 3, Col: 5},
				{Type: EOF, Value: "", Row: 3, Col: 6},
			},
		},
	}

	for _, tc := range testCases {
		lexer := NewLexer(tc.input)
		for _, expectedToken := range tc.expected {
			token, err := lexer.NextToken()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if token != expectedToken {
				t.Errorf("Input: %s, Expected token: %+v, got: %+v", tc.input, expectedToken, token)
			}
		}
	}
}

// Test error (undefined token)
func TestNextTokenError(t *testing.T) {
	input := "if x > .5 then y = 2 else then y = 3"
	lexer := NewLexer(input)
	var err error
	for i := 0; i < len(input); i++ {
		_, err = lexer.NextToken()
		if err != nil {
			break
		}
	}
	if err == nil {
		t.Errorf("Expected error, got none")
	} else {
		// Check if the error message is correct
		expected := "undefined token at row 1, col 8"
		if err.Error() != expected {
			t.Errorf("Expected error %q, got %q", expected, err.Error())
		} else {
			t.Logf("test case passed")
		}
	}
}
