package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				//a[i]++
				fmt.Printf("Hello from goroutine %d \n", i)
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
