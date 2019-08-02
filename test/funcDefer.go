package main

import (
	"fmt"
	"sync"
	"os"
)

func test() string {
	fmt.Println("defer begin")

	//将defer放入延迟调用栈
	defer fmt.Println(1)

	defer fmt.Println(2)

	//最后一个放入，位于栈顶，最先调用
	defer fmt.Println(3)

	return "defer end"
}

/**使用延迟并发解锁*/
var (
	//一个演示用的映射
	valueByKey = make(map[string]int)
	//保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

//根据键读取值
func readValue(key string) int {
	//对共享资源加锁
	valueByKeyGuard.Lock()

	//defer延迟解锁
	defer valueByKeyGuard.Unlock()

	return valueByKey[key]
}


/**使用延迟释放文件句柄*/
func fileSize(filename string) int64 {

	//根据文件名打开文件，返回文件句柄和错误
	f, err := os.Open(filename)

	//如果打开时发生错误，返回文件大小为0
	if err != nil {
		return 0
	}

	//延迟关闭文件
	defer f.Close()

	//取文件状态信息
	info, err := f.Stat()

	//如果获取信息时发生错误，关闭文件并返回文件大小为0
	if err != nil {
		return 0
	}

	//取文件大小
	size := info.Size()

	return size
}



func main(){
	fmt.Println(test())

	fmt.Println(fileSize("sin.png"))
}