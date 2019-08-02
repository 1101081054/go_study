package main

import (
	"fmt"
)

/**
定义格式：
func(参数列表)(返回参数列表){
	函数体
}
*/

/**
例1：
func(data int){
	fmt.Println("hello", data)
}(100) //(100)，表示对匿名函数进行调用，传递参数为100
*/

/**
例2：
//将匿名函数体保存到f()中
f := func(data int){
	fmt.Println("hello", data)
}
//使用f()调用
f(100)
*/

//遍历切片的每个元素，通过给定函数进行元素访问
func visit(list []int, f func(int)){
	for _, v := range list{
		f(v)
	}
}

func main(){
	//使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int){
		fmt.Println(v)
	})
}