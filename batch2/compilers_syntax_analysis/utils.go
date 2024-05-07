package main

import "fmt"

func (n *StatementNode) String() string {
	var result string
	for _, child := range n.Children {
		result += child.String() + "\n"
	}
	return result
}

func (n *IfStatementNode) String() string {
	result := "If\n"
	result += "  Condition: " + n.Condition.String() + "\n"
	result += "  Then:\n" + indent(n.ThenBlock.String(), 4)
	if n.ElseBlock != nil {
		result += "  Else:\n" + indent(n.ElseBlock.String(), 4)
	}
	return result
}

func (n *AssignmentNode) String() string {
	return fmt.Sprintf("Assignment: %s = %s", n.Identifier, n.Expression.String())
}

func (n *BinaryExpressionNode) String() string {
	return fmt.Sprintf("(%s %s %s)", n.Left.String(), n.Operator, n.Right.String())
}

func (n *LiteralNode) String() string {
	return n.Value
}

func (n *IdentifierNode) String() string {
	return n.Value
}
