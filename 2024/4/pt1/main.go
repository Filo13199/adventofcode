package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2024/4/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	words := []string{}
	grid := []string{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		grid = append(grid, string(slices.Clone(bytes)))

	}

	words = append(words, grid...)

	colGrid := make([]string, len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			colGrid[j] += string(grid[i][j])
		}
	}

	words = append(words, colGrid...)

	diagonalGrid := make([]string, 4*len(grid)-2)

	i, j := 0, 0
	dCount := 0
	for iStart := 0; iStart < len(grid); iStart++ {
		i = iStart
		j = 0
		for i < len(grid) && j < len(grid) {
			diagonalGrid[dCount] += string(grid[i][j])
			i++
			j++
		}
		dCount++
	}

	for jStart := 1; jStart < len(grid); jStart++ {
		j = jStart
		i = 0
		for i < len(grid) && j < len(grid) {
			diagonalGrid[dCount] += string(grid[i][j])
			i++
			j++
		}
		dCount++
	}

	for iStart := 0; iStart < len(grid); iStart++ {
		i = iStart
		j = len(grid) - 1
		for i < len(grid) && j > -1 {
			diagonalGrid[dCount] += string(grid[i][j])
			i++
			j--
		}
		dCount++
	}

	for jStart := len(grid[0]) - 2; jStart > -1; jStart-- {
		i = 0
		j = jStart
		for i < len(grid) && j > -1 {
			diagonalGrid[dCount] += string(grid[i][j])
			i++
			j--
		}
		dCount++
	}

	count := 0
	words = append(words, diagonalGrid...)

	for i := range words {
		count += strings.Count(words[i], "XMAS")
		count += strings.Count(words[i], "SAMX")
	}

	fmt.Println(count)
}
