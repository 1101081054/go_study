package main

import (
	"fmt"
	"container/list"
)

//list的初始化有两种方法：New 和 声明（效果一致）
//变量名 := list.New()
//var 变量名 list.List

//列表的元素可以是任意类型
//但，给一个列表放入了非期望类型的值，在取出值

func main(){
	l := list.New()

	//尾部添加
	l.PushBack("canon")

	//头部添加
	l.PushFront(67)

	//尾部添加后保存元素句柄
	element := l.PushBack("fist")

	//在fist之后添加high
	l.InsertAfter("high", element)

	//在fist之前添加noon
	l.InsertBefore("noon", element)

	l2 := list.New()
	l2.PushBack("canon2")
	l2.PushFront(672)
	//添加l2列表到l的头部
	l.PushFrontList(l2)


	l3 := list.New()
	l3.PushBack("canon3")
	l3.PushFront(673)
	//添加l3列表到l的尾部
	l.PushBackList(l3)

	//移除element对应的元素
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}