package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			"if x > 5 then y = x+2 else y = 0",
			[]Token{
				{Type: KEYWORD, Value: "if", Row: 1, Col: 1},
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 4},
				{Type: OPERATOR, Value: "<", Row: 1, Col: 6},
				{Type: LITERAL, Value: "5", Row: 1, Col: 8},
				{Type: KEYWORD, Value: "then", Row: 1, Col: 10},
				{Type: IDENTIFIER, Value: "y", Row: 1, Col: 15},
				{Type: ASSIGN, Value: "=", Row: 1, Col: 17},
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 19},
				{Type: OPERATOR, Value: "+", Row: 1, Col: 20},
				{Type: LITERAL, Value: "2", Row: 1, Col: 21},
				{Type: KEYWORD, Value: "else", Row: 1, Col: 23},
				{Type: IDENTIFIER, Value: "y", Row: 1, Col: 25},
				{Type: ASSIGN, Value: "=", Row: 1, Col: 27},
				{Type: LITERAL, Value: "0", Row: 1, Col: 29},
			},
		},
		// Test with white spaces and multilines
		{
			"if (x==1)\n\treturn var1\nelse return var2     \t",
			[]Token{
				{Type: KEYWORD, Value: "if", Row: 1, Col: 1},
				{Type: LPAR, Value: "(", Row: 1, Col: 4},
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 5},
				{Type: OPERATOR, Value: "==", Row: 1, Col: 6},
				{Type: LITERAL, Value: "1", Row: 1, Col: 7},
				{Type: RPAR, Value: ")", Row: 1, Col: 8},
				{Type: KEYWORD, Value: "return", Row: 2, Col: 2}, // return is in the column 2 after \t and 2nd line
				{Type: IDENTIFIER, Value: "var1", Row: 2, Col: 9},
				{Type: KEYWORD, Value: "else", Row: 3, Col: 1},
				{Type: KEYWORD, Value: "return", Row: 3, Col: 6},
				{Type: IDENTIFIER, Value: "var2", Row: 3, Col: 8},
			},
		},
		// Same test as above but with single line comment
		{
			"if (x==1)\n\treturn var1 \\ This will be ignored\nelse return var2     \t",
			[]Token{
				{Type: KEYWORD, Value: "if", Row: 1, Col: 1},
				{Type: LPAR, Value: "(", Row: 1, Col: 4},
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 5},
				{Type: OPERATOR, Value: "==", Row: 1, Col: 6},
				{Type: LITERAL, Value: "1", Row: 1, Col: 7},
				{Type: RPAR, Value: ")", Row: 1, Col: 8},
				{Type: KEYWORD, Value: "return", Row: 2, Col: 2}, // return is in the column 2 after \t and 2nd line
				{Type: IDENTIFIER, Value: "var1", Row: 2, Col: 9},
				{Type: KEYWORD, Value: "else", Row: 3, Col: 1},
				{Type: KEYWORD, Value: "return", Row: 3, Col: 6},
				{Type: IDENTIFIER, Value: "var2", Row: 3, Col: 8},
			},
		},
		// Same test as above but with multi line comment
		{
			"if (x==1)\n\treturn var1 /* This will \n be ignored */\nelse return var2     \t",
			[]Token{
				{Type: KEYWORD, Value: "if", Row: 1, Col: 1},
				{Type: LPAR, Value: "(", Row: 1, Col: 4},
				{Type: IDENTIFIER, Value: "x", Row: 1, Col: 5},
				{Type: OPERATOR, Value: "==", Row: 1, Col: 6},
				{Type: LITERAL, Value: "1", Row: 1, Col: 7},
				{Type: RPAR, Value: ")", Row: 1, Col: 8},
				{Type: KEYWORD, Value: "return", Row: 2, Col: 2}, // return is in the column 2 after \t and 2nd line
				{Type: IDENTIFIER, Value: "var1", Row: 2, Col: 9},
				{Type: KEYWORD, Value: "else", Row: 4, Col: 1},
				{Type: KEYWORD, Value: "return", Row: 4, Col: 6},
				{Type: IDENTIFIER, Value: "var2", Row: 4, Col: 8},
			},
		},
	}

	for _, test := range tests {
		tokens, err := tokenize(test.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(tokens) != len(test.expected) {
			t.Errorf("Input: %q - Expected %d tokens, got %d",
				test.input, len(tokens),
				len(test.expected))
		}

		for i, token := range tokens {
			if token != test.expected[i] {
				t.Errorf("Input %q - Token %d mismatch: Expected %v, got %v",
					test.input, i, test.expected[i], token)
			}
		}
	}
}

func TestLexerHasError(t *testing.T) {
	tests := []string{
		"if x > .5 then y = 2  else then y = 3",
	}

	for _, test := range tests {
		_, err := tokenize(test)

		if err != nil {
			t.Errorf("Expected error while tokenizing %v", test)
		}
	}
}
