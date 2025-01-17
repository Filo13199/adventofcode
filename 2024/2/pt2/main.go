package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkSlice(nums []int64) bool {
	positive := nums[0]-nums[1] >= 0
	for i := 0; i < len(nums)-1; i++ {
		diff := int(math.Abs(float64(nums[i] - nums[i+1])))
		currPossitive := (nums[i] - nums[i+1]) >= 0

		if diff == 0 || diff > 3 || currPossitive != positive {
			return false
		}

	}

	return true
}

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2024/2/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	safeReports := 0
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		numsStr := strings.Split(string(bytes), " ")
		nums := make([]int64, 0, len(numsStr))
		for i := range numsStr {
			num, _ := strconv.ParseInt(numsStr[i], 10, 64)
			nums = append(nums, num)
		}
		safe := checkSlice(nums)
		if !safe {
			for i := 0; i < len(nums); i++ {
				copy := slices.Clone(nums)
				tempNums := append(copy[0:i], copy[i+1:]...)

				safe = checkSlice(tempNums)
				if safe {
					break
				}
			}
		}
		if safe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
