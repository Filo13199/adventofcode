package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type gridElement struct {
	char              string
	visited           bool
	visited2          bool
	distanceFromStart int
}

// func traverseHelper(grid [][]gridElement, i, j, steps int) int {
// 	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 {
// 		return 0
// 	}

// 	if grid[i][j].char != "." && grid[i][j].char != "S" {
// 		return 0
// 	}

// 	if steps > 64 {
// 		return 0
// 	}

// 	// goDown := ((i<len(grid)-1) && grid[i+1][j].char=="." && !grid[i+1][j].visited)
// 	// goUp :=((i>0) && grid[i-1][j].char=="." && !grid[i-1][j].visited)
// 	// goLeft := ((j>0) && grid[i][j-1].char=="." && !grid[i][j-1].visited)
// 	// goRight := ((j<len(grid[0])-1)&& grid[i][j+1].char == "." && !grid[i][j+1].visited)
// 	incr := 0

// 	if grid[i][j].distanceFromStart && !grid[i][j].visited2 {
// 		grid[i][j].visited2 = true
// 		incr = 1
// 	}

// 	steps++

// 	return incr + traverseHelper(grid, i+1, j, steps) + traverseHelper(grid, i-1, j, steps) + traverseHelper(grid, i, j+1, steps) + traverseHelper(grid, i, j-1, steps)
// }

func traverseAll(grid [][]gridElement, i, j, stepsSoFar int) {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 {
		return
	}

	if grid[i][j].char == "#" {
		return
	}

	if grid[i][j].visited && stepsSoFar >= grid[i][j].distanceFromStart {
		return
	}

	grid[i][j].distanceFromStart = stepsSoFar
	grid[i][j].visited = true
	stepsSoFar++
	traverseAll(grid, i+1, j, stepsSoFar)
	traverseAll(grid, i-1, j, stepsSoFar)
	traverseAll(grid, i, j+1, stepsSoFar)
	traverseAll(grid, i, j-1, stepsSoFar)

	return
}

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2023/21/sample.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)
	grid := [][]gridElement{}
	rows := 0
	x, y := -1, -1
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		row := make([]gridElement, 0, len(bytes))
		for i := range bytes {
			row = append(row, gridElement{char: string(bytes[i])})
			ind := slices.Index(bytes, 'S')
			if ind > -1 {
				y = rows
				x = ind
			}
		}
		grid = append(grid, row)
		rows++

	}

	traverseAll(grid, y, x, 0)

	count := 0
	steps := 11
	rem := steps % 2
	for i := range grid {
		for j := range grid[i] {

			if grid[i][j].distanceFromStart <= steps && grid[i][j].distanceFromStart%2 == rem && grid[i][j].char != "#" && !grid[i][j].visited2 && grid[i][j].visited {
				count++
				grid[i][j].visited2 = true
			}
			char := grid[i][j].char
			if char == "." && grid[i][j].visited2 {
				char = "O"
			}
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println(count)
}
