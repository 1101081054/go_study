package main

import (
	"bufio"
	"fmt"
	"os"
	"www/study/functional/fib"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err.(*os.PathError))
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}
