package main

import "fmt"

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string{
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}
var CP ChipType = 9

func main(){
	fmt.Printf("%s %d", CP, CPU)
	
}