package main

import (
	"fmt"
)



type Test interface {
	nihao()
}

func nihao(){
	fmt.Println("测试")
}


func main() {
	nihao()

}
