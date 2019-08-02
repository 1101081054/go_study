package main

import "fmt"

/**
接口声明的格式
type 接口类型名 interface {
	方法名1(参数列表1) 返回值列表1
	方法名2(参数列表2) 返回值列表2
	...
}
*/

/**
接口被实现的条件
1、接口的方法与实现接口的类型方法格式一致
	在类型中添加与接口签名一致的方法就可以实现该方法。
	签名包括方法中的名称、参数列表、返回参数列表
2、接口中所有方法均被实现
*/

//例子

//一个服务需要满足能够开启和写日志的功能
type Service interface{
	Start()		//开启服务
	Log(string) //日志输出
}

//日志器
type Logger struct {

}

//实现Service的Log()方法
func (s *Logger) Log(l string){
	fmt.Println(l)
}

//游戏服务
type GameService struct {
	Logger //嵌入日志器
}

//实现Service的Start()方法
func (g *GameService) Start() {
	fmt.Println("开启服务")
}

func main(){
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
}