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
	for {
		b, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		bytes := slices.Clone(b)
		space = append(space, bytes)
	}

	summationF1 := solve(space, 1)
	summationF2 := solve(space, 2)
	diff := summationF2 - summationF1
	factor := 1000000
	summation := diff * int64(factor-1)
	fmt.Println(diff)
	fmt.Println(solve(space, 1))
	fmt.Println(solve(space, 2), summationF1+diff)
	fmt.Println(solve(space, 3), summationF1+(diff*int64(3-1)))
	fmt.Println(summationF1 + summation)
}

func solve(space [][]byte, factor int) int64 {
	galaxies := [][2]int{}
	for i := 0; i < len(space); i++ {
		if !slices.Contains(space[i], '#') {
			tmpSpace := [][]byte{}
			tmpSpace = append(tmpSpace, space[:i]...)
			for j := 1; j < factor; j++ {
				tmpSpace = append(tmpSpace, space[i])
			}
			tmpSpace = append(tmpSpace, space[i:]...)
			space = tmpSpace
			i += (factor - 1)
		}
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
				extra := make([]byte, factor-1)
				for t := 0; t < len(extra); t++ {
					extra[t] = '.'
				}
				tmpRow = append(tmpRow, extra...)
				tmpRow = append(tmpRow, space[i][col:]...)
				space[i] = tmpRow
			}
			col += (factor - 1)
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
		// fmt.Println(string(space[i]))
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
	return summation
}
