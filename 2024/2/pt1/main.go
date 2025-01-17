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
		positive := nums[0]-nums[1] >= 0
		isValid := true
		for i := 0; i < len(nums)-1; i++ {
			diff := int(math.Abs(float64(nums[i] - nums[i+1])))
			currPossitive := (nums[i] - nums[i+1]) >= 0
			if diff == 0 || diff > 3 || currPossitive != positive {
				isValid = false
				break
			}
		}
		if isValid {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
