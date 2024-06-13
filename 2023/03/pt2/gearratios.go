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

func NumbersLocations(schematic []string) [][][]int {
	locs := make([][][]int, len(schematic))
	numRegex := regexp.MustCompile("[0-9]+")
	for i := range schematic {
		rowLocs := numRegex.FindAllIndex([]byte(schematic[i]), -1)
		locs[i] = append(locs[i], rowLocs...)
	}

	return locs
}

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
	numberLocs := NumbersLocations(schematic)

	symbolRegex := regexp.MustCompile(`\*`)
	sum := 0
	for i := range schematic {

		startRow := int(math.Max(float64(i-1), 0))
		endRow := int(math.Min(float64(i+1), float64(len(schematic)-1)))
		gearLocs := symbolRegex.FindAllIndex([]byte(schematic[i]), -1)

		for j := range gearLocs {
			consecutiveNums := []int{}
			for k := startRow; k <= endRow; k++ {
				rowlookingat := schematic[k]
				rowlookingat = rowlookingat
				for z := range numberLocs[k] {
					numberLocationInRowLookingAt := numberLocs[k][z]
					numberLocationInRowLookingAt = numberLocationInRowLookingAt
					numStart := numberLocs[k][z][0]
					numEnd := numberLocs[k][z][1]

					absStart := math.Abs(float64(gearLocs[j][0] - numStart))
					if gearLocs[j][0]-1 >= numStart && gearLocs[j][0] <= numEnd {
						num, _ := strconv.Atoi(schematic[k][numStart:numEnd])
						consecutiveNums = append(consecutiveNums, num)
					} else if absStart == 1 || absStart == 0 {
						num, _ := strconv.Atoi(schematic[k][numStart:numEnd])
						consecutiveNums = append(consecutiveNums, num)
					} else if gearLocs[j][0] == numEnd {
						num, _ := strconv.Atoi(schematic[k][numStart:numEnd])
						consecutiveNums = append(consecutiveNums, num)
					}

				}
			}
			if len(consecutiveNums) == 2 {
				sum += consecutiveNums[0] * consecutiveNums[1]
			}
		}
	}
	fmt.Println(sum)
}
