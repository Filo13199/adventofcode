package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
)

const (
	y = iota
	x
)

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/11/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	space := [][]byte{}
	galaxies := [][2]int{}
	for {
		b, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		bytes := slices.Clone(b)
		ind := slices.Index(bytes, '#')
		if ind == -1 {
			space = append(space, bytes)
		}
		space = append(space, bytes)
	}

	for col := 0; col < len(space[0]); col++ {
		empty := true
		for row := range space {
			if space[row][col] == '#' {
				empty = false
				break
			}
		}
		if empty {
			for i := 0; i < len(space); i++ {
				tmpRow := append([]byte{}, space[i][:col]...)
				tmpRow = append(tmpRow, '.')
				tmpRow = append(tmpRow, space[i][col:]...)
				space[i] = tmpRow
			}
			col += 1
		}
	}

	for i := 0; i < len(space); i++ {
		for j := range space[i] {
			if space[i][j] == '#' {
				tuple := [2]int{}
				tuple[x] = j
				tuple[y] = i
				galaxies = append(galaxies, tuple)
			}
		}
	}

	pairs := [][2][2]int{}
	summation := int64(0)
	for i := 1; i < len(galaxies); i++ {
		for j := 0; j < i; j++ {
			pairs = append(pairs, [2][2]int{galaxies[i], galaxies[j]})
		}
	}
	for i := range pairs {
		dist := int(math.Abs(float64(pairs[i][0][x]-pairs[i][1][x]))) + int(math.Abs(float64(pairs[i][0][y]-pairs[i][1][y])))
		summation += int64(dist)
	}
	fmt.Println(summation)
}
