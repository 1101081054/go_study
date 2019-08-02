package main

import(
	"fmt"
)

/**
结构体的定义格式
type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    字段3, 字段4, 字段5 类型 //同类型的变量可以写在同一行
    …
}
*/

/**
实例化：
1、基本实例化形式(var ins T)
例：
	type Point struct {
		X, Y int
	}

	var p Point
	p.X = 10
	p.Y = 20

*/
/**
2、创建指针类型的结构体(ins := new(T))
例：
	type Player struct {
		Name string
		HealthPoint, MagicPoint int
	}

	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300
*/
//3、取结构体的地址实例化(ins := &T{})
type Command struct {
	Name string //指令名称
	Var *int //指令绑定的变量
	Comment string //指令的注释
}

// var version int = 1
// cmd := &Command{}
// cmd.Name = "version"
// cmd.Var = &version
// cmd.Comment = "show version"

//可以使用函数封装初始化过程
func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name: name,
		Var: varref,
		Comment: comment,
	}
}


/**初始化匿名结构体
ins := struct {
	//匿名结构体字段定义
	字段1 类型1
	字段2 类型2
	...
}{
	//字段值初始化
	字段1: 值1,
	字段2: 值2
	...
}
*/

//匿名结构体例子
func printMsgType(msg *struct {
	id int
	data string
}){
	//使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}

func main(){
	version := 1
	cmd := newCommand("version", &version, "show version")
	fmt.Println(cmd.Name)

	//实例化一个匿名结构体
	msg := &struct {
		id int
		data string
	}{
		1024,
		"hello",
	}

	printMsgType(msg)
}
