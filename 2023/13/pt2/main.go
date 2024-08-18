package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"slices"
)

func score(grid [][]byte) int {
	trnsp := make([][]byte, len(grid[0]))
	rows := make([]uint32, 0, len(grid))
	for i := range grid {
		for j := range grid[i] {
			trnsp[j] = append(trnsp[j], grid[i][j])
		}
		rows = append(rows, transformBytesToBinary(grid[i]))
	}
	cols := make([]uint32, 0, len(trnsp))
	for i := range trnsp {
		cols = append(cols, transformBytesToBinary(trnsp[i]))
	}

	for i := 0; i < len(rows)-1; i++ {
		start, end := i, i+1
		c := 0
		for start > -1 && end < len(rows) && c <= 1 {
			diff := bits.OnesCount32(rows[start] ^ rows[end])
			c += diff
			start--
			end++
		}

		if c == 1 {
			return (i + 1) * 100
		}

	}

	for i := 0; i < len(cols)-1; i++ {
		start, end := i, i+1
		c := 0
		for start > -1 && end < len(cols) && c <= 1 {
			diff := bits.OnesCount32(cols[start] ^ cols[end])
			c += diff
			start--
			end++
		}

		if c == 1 {
			return (i + 1)
		}
	}

	return 0
}

func transformBytesToBinary(row []byte) uint32 {
	rowUint := uint32(0)
	for i := len(row) - 1; i >= 0; i-- {
		if row[i] == '#' {
			x := (len(row) - i) - 1
			rowUint += 1 << x
		}
	}
	return rowUint
}

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	summary := 0
	grid := [][]byte{}
	i := 0
	for {
		bytes, _, err := reader.ReadLine()

		if errors.Is(err, io.EOF) {
			i++
			x := score(grid)
			if x == -1 {
				fmt.Printf(" %d is zero\n", x)
			}
			summary += x
			break
		}

		if len(bytes) == 0 {
			i++
			x := score(grid)
			if x == -1 {
				fmt.Printf(" %d is zero \n", x)
			}
			fmt.Printf(" row #%d =>  score = %d\n", i, x)
			summary += x
			grid = [][]byte{}
		} else {
			grid = append(grid, slices.Clone(bytes))
		}

	}
	fmt.Println(summary)
}
