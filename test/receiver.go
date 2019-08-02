package main

import(
	"fmt"
)

/**
接收器 ----- 方法作用的目标
接收器的格式如下：
	func(接收器变量 接收器类型) 方法名 (参数列表)(返回参数){
		函数体
	}
	对各部分的说明：
	1、接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母
	2、接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型
	3、方法名、参数列表、返回参数：格式与函数定义一致
*/

/**指针类型的接收器*/
type Property struct {
	value int 
}
func (p *Property) setValue(v int){
	p.value = v
}
func (p *Property) Value() int {
	return p.value
}


/**非指针类型的接收器*/
type Point struct {
	X int
	Y int
}
func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func main(){
	//实例化指针类型的接收器
	p := new(Property)
	p.setValue(100)
	fmt.Println(p.Value())

	//实例化非指针类型的接收器
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	result := p1.Add(p2)
	fmt.Println(result)
}