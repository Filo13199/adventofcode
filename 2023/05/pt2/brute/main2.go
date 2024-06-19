package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	dstRangeStartAccessor = iota
	srcRangeStartAccessor
	rangeLengthAccessor
)

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/05/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)
	reader.Discard(7)
	seedsBytes, _, _ := reader.ReadLine()
	seedsSplit := strings.Split(string(seedsBytes), " ")
	fmt.Println(len(seedsSplit))
	seeds := make([][2]int64, 0, int(math.Floor(float64(len(seedsSplit)/2))))
	for i := 0; i < len(seedsSplit); i += 2 {

		seedStart, _ := strconv.ParseInt(seedsSplit[i], 10, 64)
		seedEnd, _ := strconv.ParseInt(seedsSplit[i+1], 10, 64)
		seeds = append(seeds, [2]int64{seedStart, seedStart + seedEnd - 1})
	}
	reader.ReadLine()
	conversions := [7][][3]int64{}
	i := -1
	for {
		line, _, e := reader.ReadLine()
		if errors.Is(e, io.EOF) {
			break
		}
		if strings.Contains(string(line), "map") {
			i++
			continue
		}
		if len(line) == 0 {
			continue
		}
		rangeRule := [3]int64{}
		nums := strings.Split(string(line), " ")
		for j := range nums {
			num, _ := strconv.ParseInt(nums[j], 10, 64)
			rangeRule[j] = num
		}
		conversions[i] = append(conversions[i], rangeRule)
	}
	minDstSoFar := int64(math.MaxInt64)
	for i := range seeds {
		seedRange := seeds[i]
		for s := seedRange[0]; s <= seedRange[1]; s++ {
			dst := s
			for _, conversionMap := range conversions {
				for _, rangeRule := range conversionMap {
					diffStart := dst - rangeRule[srcRangeStartAccessor]
					if diffStart >= 0 && (rangeRule[srcRangeStartAccessor]+rangeRule[rangeLengthAccessor]-1) >= dst {
						offset := rangeRule[dstRangeStartAccessor] - rangeRule[srcRangeStartAccessor]
						dst += offset
						break
					}
				}
			}

			if dst < minDstSoFar {
				minDstSoFar = dst
			}
		}
	}
	fmt.Println(minDstSoFar)
}
