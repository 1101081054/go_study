package main

import (
	"fmt"
	"errors"
)

func main(){
	err := errors.New("this is an error")

	fmt.Println(err.Error())
}