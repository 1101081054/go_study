package main

import (
	"fmt"
)

var eventByName = make(map[string][]func(interface{}))

/**
事件注册
事件系统需要为外部提供一个注册入口。
这个注册入口传入注册的事件名称和对应事件名称的响应函数
时间注册的过程就是将事件名称和响应函数关联并保存起来
*/
//注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	//通过名字查找事件列表
	list := eventByName[name]

	//在列表切片中添加函数
	list = append(list, callback)

	//将修改的事件列表切片保存回去
	eventByName[name] = list
}

//删除事件
func DeleteEvent(name string) {
	delete(eventByName, name)
}

//调用事件
func CallEvent(name string, param interface{}) {
	if eventByName[name] == nil {
		fmt.Println("未查找到该事件")
		return
	}
	//通过名字找到事件列表
	list := eventByName[name]

	//遍历这个事件的所有回调
	for _, callback := range list {
		//传入参数调用回调
		callback(param)
	}
}


/**使用事件系统*/
type Actor struct {

}

//为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event: ", param)
}

//全局事件
func GlobelEvent(param interface{}) {
	fmt.Println("global event: ", param)
}

func main(){
	//实例化一个角色
	a := new(Actor)

	//注册名为OnSkill的回调
	RegisterEvent("OnSkill", a.OnEvent)

	//再次在OnSkill上注册全局事件
	RegisterEvent("OnSkill", GlobelEvent)

	//调用事件，所有注册的同名函数都会被调用
	CallEvent("OnSkill", 100)
	DeleteEvent("OnSkill")
	CallEvent("OnSkill", 100)
}