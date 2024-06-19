package main2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	dstRangeStartAccessor = iota
	srcRangeStartAccessor
	rangeLengthAccessor
)

const (
	srcRangeStart = iota
	srcRangeEnd
	dstRangeStart
	dstRangeEnd
)

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/05/sample.txt")
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
	conversions := [7][][4]int64{}
	n := -1
	for {
		line, _, e := reader.ReadLine()
		if errors.Is(e, io.EOF) {
			break
		}
		if strings.Contains(string(line), "map") {
			n++
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

		rangeTuple := [4]int64{
			rangeRule[srcRangeStartAccessor],
			rangeRule[srcRangeStartAccessor] + rangeRule[rangeLengthAccessor] - 1,
			rangeRule[dstRangeStartAccessor],
			rangeRule[dstRangeStartAccessor] + rangeRule[rangeLengthAccessor] - 1,
		}
		conversions[n] = append(conversions[n], rangeTuple)
	}

	for i := range conversions {
		slices.SortFunc(conversions[i], func(a, b [4]int64) int {
			if a[srcRangeStart] < b[srcRangeStart] {
				return -1
			} else if b[srcRangeStart] > a[srcRangeStart] {
				return 1
			}
			return 0
		})
	}

	fmt.Println(conversions)
	currentRange := conversions[len(conversions)-1][0]
	rangesSortedAscendingly := [][2]int64{}
	lastMapping := conversions[len(conversions)-1]
	maxRangeSoFar := int64(0)
	for i := 0; i < len(lastMapping); i++ {
		curr := lastMapping[i]
		if curr[srcRangeStart]-1 > int64(maxRangeSoFar) {
			rangesSortedAscendingly = append(rangesSortedAscendingly, [2]int64{int64(maxRangeSoFar), curr[srcRangeStart] - 1})
			rangesSortedAscendingly = append(rangesSortedAscendingly, [2]int64{curr[srcRangeStart], curr[srcRangeEnd]})
			maxRangeSoFar = curr[srcRangeEnd]
		} else {
			rangesSortedAscendingly = append(rangesSortedAscendingly, [2]int64{curr[srcRangeStart], curr[srcRangeEnd]})
			maxRangeSoFar = curr[srcRangeEnd]
		}
	}

	fmt.Println(conversions)

	// TODO, sort the conversions by the dst start, and backtrack from the last conversion, give priority to src numbers not in dst range that are smaller than the current dst start
}
