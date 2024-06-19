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
	for {
		bytes, _, err := reader.ReadLine()

		if errors.Is(err, io.EOF) {
			break
		}
		split := strings.Split(string(bytes), " = ")
		tupleStr := strings.Trim(split[1], "()")
		tuple := strings.Split(tupleStr, ", ")
		nodes[split[0]] = [2]string{tuple[left], tuple[right]}
	}

	curr := "AAA"
	i := 0
	for {
		if curr == "ZZZ" {
			break
		}
		for i := range instructions {
			curr = nodes[curr][instructions[i]]
		}
		i++
	}
	fmt.Println(len(instructions) * i)
}
