package main

import (
	"fmt"
)

//提供一个值，每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {

	//返回一个闭包
	return func() int {
		value++
		return value
	}
}

func main(){
	value := 1
	//创建一个累加器，初始值为1
	accumulator := Accumulate(value)

	//累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	value = 10
	accumulator2 := Accumulate(value)
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())

	fmt.Println(accumulator())

}