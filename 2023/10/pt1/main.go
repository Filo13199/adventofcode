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

const (
	y = iota
	x
)

const (
	up = iota
	down
	right
	left
)

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/10/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	currentPos := [2]int{-1, -1}
	pipesMap := [][]rune{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if slices.Contains(bytes, 'S') {
			currentPos[x] = slices.Index(bytes, 'S')
			currentPos[y] = len(pipesMap)
		}
		runes := make([]rune, len(bytes))
		for i := range bytes {
			runes[i] = rune(bytes[i])
		}
		pipesMap = append(pipesMap, runes)
	}
	animalPos := [2]int{}
	animalPos[x] = currentPos[x]
	animalPos[y] = currentPos[y]
	direction := up
	xDir, yDir := 1, 1
	switch direction {
	case up:
		yDir = -1
		currentPos[y] += yDir
	case down:
		yDir = 1
		currentPos[y] += yDir
	case right:
		xDir = 1
		currentPos[x] += xDir
	case left:
		xDir = -1
		currentPos[x] += xDir
	}
	i := 1
	for currentPos != animalPos {
		currChar := pipesMap[currentPos[y]][currentPos[x]]
		switch currChar {
		case '|':
			currentPos[y] += yDir
		case '-':
			currentPos[x] += xDir

		case 'L':
			if direction == left {
				direction = up
				yDir = -1
				currentPos[y] += yDir
			} else if direction == down {
				direction = right
				xDir = 1
				currentPos[x] += xDir
			}
		case 'J':
			if direction == right {
				direction = up
				yDir = -1
				currentPos[y] += yDir
			} else if direction == down {
				direction = left
				xDir = -1
				currentPos[x] += xDir
			}
		case '7':
			if direction == up {
				direction = left
				xDir = -1
				currentPos[x] += xDir
			} else if direction == right {
				direction = down
				yDir = 1
				currentPos[y] += yDir
			}
		case 'F':
			if direction == up {
				direction = right
				xDir = 1
				currentPos[x] += xDir
			} else if direction == left {
				direction = down
				yDir = 1
				currentPos[y] += yDir
			}
		}
		i++
	}
	fmt.Println(i / 2)
}
