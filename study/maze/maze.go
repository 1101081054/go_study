package main

import (
	"fmt"
	"os"
)

func realMaze(filename string) [][]int {
	file, err  := os.Open(filename)
	if err != nil {
		panic(nil)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func main() {
	maze := realMaze("go_study/study/maze/maze.in")
	for _, row := range maze{
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
