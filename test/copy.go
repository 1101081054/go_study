package main

import (
	"fmt"
)

func main(){
	//设置元素数量为1000
	const elementCount = 1000

	//预分配足够多的元素切片
	srcData := make([]int, elementCount)

	//将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	//引用切片数据
	refData := srcData

	//预分配足够多的元素切片
	copyData := make([]int, elementCount)
	//将数据复制到新的切片空间中
	copy(copyData, srcData)

	//修改原始数据的第一个元素
	srcData[0] = 999

	//打印因用工切片的第一个元素
	fmt.Println(refData[0])

	//打印复制切片的第一个元素和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])

	//复制原始数据从4到6（不包含）
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

	copy(copyData, srcData[10:20])
	fmt.Println(copyData)

}