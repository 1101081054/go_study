package main

import (
	"fmt"
)

func main(){
	var car []string

	car = append(car, "oldDriver")

	car = append(car, "Ice", "Sniper", "Monk")

	team := []string{"Pig", "Flyingcake", "Chicken"}

	car = append(car, team...)

	fmt.Println(car)

	var numbers []int
	for i := 0; i < 8; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d ; cap: %d; pointer: %p \n", len(numbers), cap(numbers), numbers)
	}
	numbers2 := append(numbers, 8)
	fmt.Println(numbers)
	fmt.Println(numbers2)
}