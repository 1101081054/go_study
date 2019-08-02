package main

/**
什么时候触发宕机？
在运行依赖的必备资源缺失时主动触发宕机
*/

import (
	"fmt"
)


/**
1) func Compile(expr string)(*Regexp, error)
	编译正则表达式，发生错误时返回编译错误，Regexp为nil，
	该函数适用于在编译错误时获得编译错误进行处理，同时继续后续执行的环境
2) func MustCompile(str string) *Regexp
	当编译正则表达式发生错误时，使用panic触发宕机
	该函数适用于直接适用正则表达式而无需处理正则表达式错误的情
*/

// func MustCompile(str string) *Regexp {
// 	regexp, error := Compile(str)
// 	if error != nil {
// 		panic(`regexp: Compile(` + quote(str) + `): ` + error.Error())
// 	}
// 	return regexp
// }

func main(){

	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("crash")
}