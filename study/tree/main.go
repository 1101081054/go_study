package main

import (
	"fmt"
	"www/go_study/study/tree/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value:3}
	root.Left = &tree.Node{Value:0}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count: ", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for max := range c {
		if max.Value > maxNode {
			maxNode = max.Value
		}
	}

	fmt.Println("max node value", maxNode)

}
