package main

import (
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type StatementNode struct {
	Children []Node
}

type IfStatementNode struct {
	Condition Node
	ThenBlock *StatementNode
	ElseBlock *StatementNode
}

type AssignmentNode struct {
	Identifier string
	Expression Node
}

type BinaryExpressionNode struct {
	Operator string
	Left     Node
	Right    Node
}

type LiteralNode struct {
	Value string
}

type IdentifierNode struct {
	Value string
}

type Parser struct {
	lexer  *Lexer
	tokens []Token
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() (Node, error) {
	p.tokens = make([]Token, 0)

	for {
		token, err := p.lexer.NextToken()
		if err != nil {
			return nil, err
		}

		p.tokens = append(p.tokens, token)

		if token.Type == EOF {
			break
		}
	}

	return p.parseStatements()
}

func (p *Parser) parseStatements() (Node, error) {
	stmtNode := &StatementNode{Children: []Node{}}

	for len(p.tokens) > 0 {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		stmtNode.Children = append(stmtNode.Children, stmt)
	}

	return stmtNode, nil
}

func (p *Parser) parseStatement() (Node, error) {
	token := p.tokens[0]

	if token.Type == KEYWORD && token.Value == "if" {
		return p.parseIfStatement()
	} else if token.Type == IDENTIFIER {
		return p.parseAssignment()
	}

	return nil, fmt.Errorf("unexpected token: %+v", token)
}

func (p *Parser) parseIfStatement() (Node, error) {
	p.tokens = p.tokens[1:] // Consume "if" keyword

	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if len(p.tokens) == 0 || p.tokens[0].Type != KEYWORD || p.tokens[0].Value != "then" {
		return nil, fmt.Errorf("missing 'then' keyword")
	}
	p.tokens = p.tokens[1:] // Consume "then" keyword

	thenBlock, err := p.parseStatements()
	if err != nil {
		return nil, err
	}

	var elseBlock *StatementNode
	if len(p.tokens) > 0 && p.tokens[0].Type == KEYWORD && p.tokens[0].Value == "else" {
		p.tokens = p.tokens[1:] // Consume "else" keyword
		e, err := p.parseStatements()
		if err != nil {
			return nil, err
		}
		elseBlock = e.(*StatementNode)
	}

	return &IfStatementNode{Condition: condition, ThenBlock: thenBlock.(*StatementNode), ElseBlock: elseBlock}, nil
}

func (p *Parser) parseAssignment() (Node, error) {
	identifier := p.tokens[0].Value
	p.tokens = p.tokens[1:]

	if len(p.tokens) == 0 || p.tokens[0].Type != ASSIGN {
		return nil, fmt.Errorf("missing assignment operator")
	}
	p.tokens = p.tokens[1:] // Consume assignment operator

	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return &AssignmentNode{Identifier: identifier, Expression: expression}, nil
}

func (p *Parser) parseExpression() (Node, error) {
	term, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for len(p.tokens) > 0 && p.tokens[0].Type == OPERATOR {
		operator := p.tokens[0].Value
		p.tokens = p.tokens[1:]

		rightTerm, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		term = &BinaryExpressionNode{Operator: operator, Left: term, Right: rightTerm}
	}

	return term, nil
}

func (p *Parser) parseTerm() (Node, error) {
	factor, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	for len(p.tokens) > 0 && p.tokens[0].Type == OPERATOR {
		operator := p.tokens[0].Value
		p.tokens = p.tokens[1:]

		rightFactor, err := p.parseFactor()
		if err != nil {
			return nil, err
		}

		factor = &BinaryExpressionNode{Operator: operator, Left: factor, Right: rightFactor}
	}

	return factor, nil
}

func (p *Parser) parseFactor() (Node, error) {
	token := p.tokens[0]
	p.tokens = p.tokens[1:]

	if token.Type == LITERAL {
		return &LiteralNode{Value: token.Value}, nil
	} else if token.Type == LPAR {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if len(p.tokens) == 0 || p.tokens[0].Type != RPAR {
			return nil, fmt.Errorf("missing closing parenthesis")
		}
		p.tokens = p.tokens[1:]

		return expr, nil
	} else if token.Type == IDENTIFIER {
		return &IdentifierNode{Value: token.Value}, nil
	}

	return nil, fmt.Errorf("unexpected token: %+v", token)
}

func indent(s string, spaces int) string {
	result := ""
	for _, line := range strings.Split(s, "\n") {
		result += strings.Repeat(" ", spaces) + line + "\n"
	}
	return result
}

func main() {
	lexer := NewLexer("if x == 1\n\t then x = 5 else x = 6")
	parser := NewParser(lexer)
	ast, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ast.String())
}
