package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	left = iota
	right
)

func shouldEnd(currNodes []string) bool {
	for i := range currNodes {
		if currNodes[i][2] != 'Z' {
			return false
		}
	}
	return true
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/08/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	bytes, _, _ := reader.ReadLine()
	instructions := make([]int, 0, len(bytes))
	for i := range bytes {
		if bytes[i] == 'R' {
			instructions = append(instructions, right)
		} else {
			instructions = append(instructions, left)
		}
	}
	reader.ReadLine()
	nodes := make(map[string][2]string)
	currNodes := []string{}
	for {
		bytes, _, err := reader.ReadLine()

		if errors.Is(err, io.EOF) {
			break
		}
		split := strings.Split(string(bytes), " = ")
		tupleStr := strings.Trim(split[1], "()")
		tuple := strings.Split(tupleStr, ", ")
		nodes[split[0]] = [2]string{tuple[left], tuple[right]}
		if split[0][2] == 'A' {
			currNodes = append(currNodes, split[0])
		}
	}

	steps := make([]int, len(currNodes))
	for i := 0; i < len(currNodes); i++ {
		n := 0
		for {
			for j := range instructions {
				n++
				currNodes[i] = nodes[currNodes[i]][instructions[j]]
				if currNodes[i][2] == 'Z' {
					steps[i] = n
					break
				}
			}
			if currNodes[i][2] == 'Z' {
				break
			}
		}
	}

	fmt.Println(LCM(steps[0], steps[1], steps...))
	// fmt.Println(len(instructions) * n)
}
