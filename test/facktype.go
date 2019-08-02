package main

import (
	"fmt"
	"reflect"
)

//定义商标结构
type Brand struct{

}

//为商标结构添加Show()方法
func (t Brand) Show() {

}

//为Brand定义一个别名FakeBrand
type FakeBrand = Brand

//定义车辆结构
type Vehicle struct {
	//嵌入两个结构
	FakeBrand
	Brand
}

func main(){
	var a Vehicle

	a.FakeBrand.Show()

	ta := reflect.TypeOf(a)
	
	for i := 0; i < ta.NumField(); i++ {
		//a的成员信息
		f := ta.Field(i)

		//打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}

}