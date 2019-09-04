package main

import (
	"fmt"
	"os"
)

func realMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {

			fmt.Fscan(file, &maze[i][j])
		}
	}

	return maze
}

type Point struct {
	i, j int
}

var dirs = [4]Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p Point) add(dir Point) (Point, bool) {
	i := p.i + dir.i
	j := p.j + dir.j
	if i < 0 || i >= len(maze){
		return Point{0, 0}, false
	}

	if j < 0 || j >= len(maze[i]) {
		return Point{0, 0}, false
	}

	return Point{i, j}, true
}

func (p Point) at(grid [][]int) int {
	return grid[p.i][p.j]
}

func walk(maze [][]int, start, end Point) ([][]int, []Point, bool) {
	steps := make([][]int, len(maze))
	for s := range steps {
		steps[s] = make([]int, len(maze[s]))
	}

	Queue := []Point{start}
	var ifEnd = false
	endfor:
		for len(Queue) > 0 {
			cur := Queue[0]
			Queue = Queue[1:]

			for _, dir := range dirs {
				next, ok := cur.add(dir)
				if !ok {
					continue
				}

				val := next.at(maze)
				if val == 1 {
					continue
				}

				val = next.at(steps)
				if val != 0 {
					continue
				}

				if next == start {
					continue
				}

				steps[next.i][next.j] = cur.at(steps) + 1
				Queue = append(Queue, next)

				if next == end {
					ifEnd = true
					break endfor
				}
			}
		}

	route := []Point{}
	if !ifEnd {
		return steps, route, ifEnd
	}

	route = append(route, end)
	endRoute:
	for {
		cur := route[len(route) - 1]
		for _, dir := range dirs {
			next, ok := cur.add(dir)
			if !ok {
				continue
			}

			if steps[cur.i][cur.j] - 1 != steps[next.i][next.j] {
				continue
			}
			route = append(route, next)

			if next == start {
				break endRoute
			}
		}
	}
	return steps, route, ifEnd
}


var maze [][]int
func main() {
	maze = realMaze("go_study/study/maze/maze.in")
	steps, route, ok := walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})

	if !ok {
		fmt.Println("走不到终点")
	}

	fmt.Println(route)

	for _, row := range steps{
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
