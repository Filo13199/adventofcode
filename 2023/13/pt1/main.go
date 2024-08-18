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

func score(grid [][]byte) int {
	trnsp := make([][]byte, len(grid[0]))
	rows := make([]string, 0, len(grid))
	for i := range grid {
		for j := range grid[i] {
			trnsp[j] = append(trnsp[j], grid[i][j])
		}
		rows = append(rows, string(grid[i]))
	}

	cols := make([]string, 0, len(trnsp))
	for i := range trnsp {
		cols = append(cols, string(trnsp[i]))
	}
	if cols[0] == cols[len(cols)-1] {
		// fmt.Println(cols[0], "     =>  ", cols[len(cols)-1])
	}
	horizontalReflPoints := []int{}
	for i := 0; i < len(grid)-1; i++ {
		if string(grid[i]) == string(grid[i+1]) {
			horizontalReflPoints = append(horizontalReflPoints, i+1)
			if i == 0 {
				return 100
			}
			if i == len(grid)-2 {
				return (i + 1) * 100
			}
		}
	}

	verticalReflPoints := []int{}
	for i := 0; i < len(trnsp)-1; i++ {
		if string(trnsp[i]) == string(trnsp[i+1]) {
			verticalReflPoints = append(verticalReflPoints, i+1)
			if i == 0 {
				return 1
			}
			if i == len(trnsp)-2 {
				return (i + 1)
			}
		}
	}
	score := 0

	for _, v := range horizontalReflPoints {
		valid := true
		for i := 1; (v+i) < len(grid) && (v-i-1) >= 0; i++ {
			if string(grid[v+i]) != string(grid[v-i-1]) {
				valid = false
				break
			}
		}
		if valid {
			return v * 100
		}
	}
	if score == 0 {
		for _, v := range verticalReflPoints {
			valid := true
			for i := 1; (v+i) < len(trnsp) && (v-i-1) >= 0; i++ {
				if string(trnsp[v+i]) != string(trnsp[v-i-1]) {
					valid = false
					break
				}
			}
			if valid {
				return v
			}
		}
	}

	// fmt.Printf("%+v",grid)
	// fmt.Printf("transpose : %+v", trnsp)
	return 0
}

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	summary := 0
	grid := [][]byte{}
	for {
		bytes, _, err := reader.ReadLine()

		if errors.Is(err, io.EOF) {
			summary += score(grid)
			break
		}

		if len(bytes) == 0 {
			summary += score(grid)
			grid = [][]byte{}
		} else {
			grid = append(grid, slices.Clone(bytes))
		}
	}
	fmt.Println(summary)
}
