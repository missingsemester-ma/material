package main

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	input        string
	currentIndex int
	row          int
	col          int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, currentIndex: 0, row: 1, col: 1}
}

func (l *Lexer) NextToken() (Token, error) {
	var token Token
	insideComment := false

	for l.currentIndex < len(l.input) {
		char := l.input[l.currentIndex]

		if unicode.IsSpace(rune(char)) {
			if char == '\n' {
				l.row++
				l.col = 1
			} else {
				l.col++
			}
			l.currentIndex++
			continue
		}

		if char == '\\' {
			// Single-line comment, ignore until the end of the line
			for l.currentIndex < len(l.input) && l.input[l.currentIndex] != '\n' {
				l.currentIndex++
				l.col++
			}

			// if we are at the end of the input, we should return EOF
			if l.currentIndex == len(l.input) {
				l.col++
				return Token{Type: EOF, Value: "", Row: l.row, Col: l.col}, nil
			}

			l.currentIndex++
			l.col++

			if l.input[l.currentIndex] == '\n' {
				l.row++
			}
			continue
		}

		if char == '/' && l.currentIndex+1 < len(l.input) && l.input[l.currentIndex+1] == '*' {
			// Multi-line comment start
			insideComment = true
			l.currentIndex += 2
			l.col += 2
			continue
		}

		if insideComment {
			if char == '*' && l.currentIndex+1 < len(l.input) && l.input[l.currentIndex+1] == '/' {
				// Multi-line comment end
				insideComment = false
				l.currentIndex += 2
				l.col += 2
			} else {
				l.currentIndex++
				l.col++
			}
			continue
		}

		if l.currentIndex+1 < len(l.input) && string(char)+string(l.input[l.currentIndex+1]) == "==" {
			token = createToken("==", l.row, l.col)
			l.col += 2
			l.currentIndex += 2
			return token, nil
		}

		if char == '(' || char == ')' || char == '=' || char == '<' || char == '>' {
			token = createToken(string(char), l.row, l.col)
			l.col++
			l.currentIndex++
			return token, nil
		}

		// Handle identifiers and numbers
		if unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char)) {
			start := l.currentIndex
			for l.currentIndex < len(l.input) && (unicode.IsLetter(rune(l.input[l.currentIndex])) || unicode.IsDigit(rune(l.input[l.currentIndex]))) {
				l.currentIndex++
			}
			token = createToken(l.input[start:l.currentIndex], l.row, l.col)
			l.col += l.currentIndex - start
			return token, nil
		}

		// Handle operators
		if char == '+' || char == '-' || char == '*' || char == '/' {
			token = createToken(string(char), l.row, l.col)
			l.col++
			l.currentIndex++
			return token, nil
		}

		// if we reach here, it means we have an unexpected character
		return Token{}, fmt.Errorf("undefined token at row %d, col %d", l.row, l.col)
	}

	// if we are at the end of the input, we should return EOF
	if l.currentIndex == len(l.input) {
		return Token{Type: EOF, Value: "", Row: l.row, Col: l.col}, nil
	}

	return Token{}, fmt.Errorf("unexpected error at row %d, col %d", l.row, l.col)
}
