package main

import (
	"fmt"
)

//delete(map, Key)
//清空mapp的唯一方法就是重新make一个新的map。
//不用担心垃圾回收的速率，
//Go语言中的并行垃圾回收速率比写一个情况函数高效多了


func main(){
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "brazil")

	for k, v := range scene {
		fmt.Println(k, v)
	}

}