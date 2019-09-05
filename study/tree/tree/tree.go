package tree

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func CreateNode(value int) *Node{
	return &Node{Value:value}
}

func (n *Node) SetValue(v int){
	n.Value = v
}

func (n *Node) print(){
	fmt.Println(n.Value)
}

func (n *Node) Traverse(){
	if n == nil {
		return
	}

	n.Left.Traverse()
	n.print()
	n.Right.Traverse()
}


func (n *Node) TraverseFunc(f func(node *Node)){
	if n == nil{
		return
	}

	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func (n *Node) TraverseWithChannel() chan *Node {
	c := make(chan *Node, 1)
	go func() {
		n.TraverseFunc(func(node *Node) {
			c <- node
		})
		close(c)
	}()
	return c
}