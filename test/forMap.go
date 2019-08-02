package main

import (
	"fmt"
	"sort"
)

func main(){
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for k, v := range scene {
		fmt.Println(k, v)
	}
	
	/**
	如只遍历值，可以使用下面的形式：
	for _, v := range scene {}
	*/

	/**
	如只遍历键时，使用下面的形式
	for k := range scene {}
	*/

	//排序
	//声明一个切片保存map数据
	var sceneList []string

	//将map数据便利复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}

	//对切片进行排序
	sort.Strings(sceneList)

	//输出
	fmt.Println(sceneList)
}