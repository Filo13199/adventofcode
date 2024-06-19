package main

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

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/09/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	n := 0
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		numsStr := strings.Split(string(bytes), " ")
		nums := make([]int, len(numsStr))

		for i := range numsStr {
			num, _ := strconv.Atoi(numsStr[i])
			nums[i] = num
		}

		numsDiff := [][]int{}
		curr := slices.Clone(nums)
		x := len(nums) - 1
		for {
			currDiff := make([]int, 0, len(curr)-1)
			allZeroes := true
			for i := 0; i < x; i++ {
				allZeroes = allZeroes && (curr[i+1]-curr[i] == 0)
				currDiff = append(currDiff, curr[i+1]-curr[i])
			}
			curr = slices.Clone(currDiff)
			numsDiff = append(numsDiff, currDiff)
			x--
			if allZeroes {
				break
			}
		}
		lastDiff := 0
		for i := len(numsDiff) - 1; i >= 1; i-- {
			lastDiff = numsDiff[i-1][0] - lastDiff
		}
		n += nums[0] - lastDiff
	}
	fmt.Println(n)
}
