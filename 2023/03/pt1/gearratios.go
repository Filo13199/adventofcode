package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)
	schematic := []string{}

	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		schematic = append(schematic, string(bytes))
	}

	sum := 0
	numRegex := regexp.MustCompile("[0-9]+")

	for i := range schematic {
		numLocs := numRegex.FindAllIndex([]byte(schematic[i]), -1)
		for k := range numLocs {
			found := false
			j := int(math.Max(float64(i-1), 0))
			startAtEachRow := int(math.Max(float64(numLocs[k][0]-1), 0))
			endAtEachRow := int(math.Min(float64(numLocs[k][1]), float64(len(schematic[i])-1)))

			for {

				if j > i+1 || (j == i+1 && i == len(schematic)-1) {
					break
				}

				for z := startAtEachRow; z <= endAtEachRow; z++ {

					if j == i && z >= numLocs[k][0] && z < numLocs[k][1] {
						continue
					}

					if schematic[j][z] != '.' && !(schematic[j][z] >= '0' && schematic[j][z] <= '9') {
						found = true
					}
				}

				if found {
					val, _ := strconv.Atoi(schematic[i][numLocs[k][0]:numLocs[k][1]])
					sum += val
					break
				}

				j++
			}

		}
	}

	fmt.Println(sum)

	// fmt.Println(schematic)
}
