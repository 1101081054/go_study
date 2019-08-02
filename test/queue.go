package main

import (
	"fmt"
	"www/study/queue"
)

func main() {
	var q queue.Queue
	q.Push(2)
	q.Push(3)
	q.Push("nihao")
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty(), q)
	q.Pop()
	fmt.Println(q.IsEmpty(), q)
	q.Pop()
	fmt.Println(q.IsEmpty(), q)
}
