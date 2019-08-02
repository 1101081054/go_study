package main

import (
	"fmt"
)

//车轮
type Wheel struct {
	Size int
}

//车
type Car struct {
	Wheel
	Engine struct { //引擎
		Power int   //功率
		Type string //类型
	}
}

func main(){
	c := Car {
		Wheel: Wheel{ //初始化轮子
			Size: 18,
		},
		Engine: struct { //初始化引擎
			Power int
			Type string
		}{
			Type: "1.4T",
			Power: 143,
		},
	}

	fmt.Printf("%+v\n", c)
}