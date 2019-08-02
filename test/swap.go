package main

import(
	"fmt"
)

//交换函数
func swap(a, b *int){
	*a, *b = *b, *a
}

func main(){
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}