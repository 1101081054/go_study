package main

import (
	"fmt"
)

func main(){
	//map的定义
	//map[KeyType]ValueType
	scene := make(map[string]int)

	scene["route"] = 66

	fmt.Println(scene["route"])
	// v, ok := scene["route"]
	// fmt.Println(v, ok)
	v := scene["route2"]
	fmt.Println(v)

	//声明时填充内容
	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}

	fmt.Println(m)
}