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
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	load := 0
	dish := [][]byte{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		dish = append(dish, slices.Clone(bytes))
	}

	for j := 0; j < len(dish[0]); j++ {
		incr := 0
		for i := 0; i < len(dish); i++ {
			if dish[i][j] == '.' {
				incr++
			} else if dish[i][j] == '#' {
				incr = 0
			} else if dish[i][j] == 'O' {
				load += (incr + (len(dish) - i))
			}
		}
	}
	fmt.Println(load)
}
