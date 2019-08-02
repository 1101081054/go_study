package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/**函數式編程*/
type iAdder func(int) (int, iAdder)
func adder2(base int) iAdder {
	return func(v int) (int, iAdder){
		return base + v, adder2(base + v)
	}
}

func main() {
	//a := adder()
	//for i := 0; i <= 10; i++ {
	//	fmt.Println(a(i))
	//}

	a := adder2(0)
	var s int
	for i := 0; i <= 10; i++ {
		s, a = a(i)
		fmt.Println(s)
	}
}
