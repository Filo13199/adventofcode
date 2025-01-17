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

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2024/4/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	grid := []string{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		grid = append(grid, string(slices.Clone(bytes)))

	}

	count := 0

	for i := range grid {
		for j := range grid[i] {
			ch := grid[i][j]
			if ch == 'A' && i > 0 && i < len(grid)-1 && j > 0 && j < len(grid)-1 {
				w1 := fmt.Sprintf("%s", []byte{grid[i-1][j-1], grid[i][j], grid[i+1][j+1]})
				w2 := fmt.Sprintf("%s", []byte{grid[i-1][j+1], grid[i][j], grid[i+1][j-1]})

				if (w1 == "MAS" || w1 == "SAM") && (w2 == "MAS" || w2 == "SAM") {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
