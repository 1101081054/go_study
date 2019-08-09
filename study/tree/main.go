package main

import "www/go_study/study/tree/tree"

func main() {
	var node tree.Node
	node = tree.Node{Value:3}
	node.Left = &tree.Node{Value:0}
	node.Right = &tree.Node{5, nil, nil}


}
