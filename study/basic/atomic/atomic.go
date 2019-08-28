package main

import (
	"fmt"
	"sync"
	"time"
)

/**
互斥量
*/

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment(){
	fmt.Println("safe increment")
	a.lock.Lock()
	defer a.lock.Unlock()

	a.value++
}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
