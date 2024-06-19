package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
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
	seeds := make([]int64, len(seedsSplit))
	for i := range seedsSplit {
		seed, _ := strconv.ParseInt(seedsSplit[i], 10, 64)
		seeds[i] = seed
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

	for i, seed := range seeds {
		dst := seed
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
		seeds[i] = dst
	}

	min := seeds[0]

	for i := 1; i < len(seeds); i++ {
		if seeds[i] <= min {
			min = seeds[i]
		}
	}
	fmt.Println(min)
}
