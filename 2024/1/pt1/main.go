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

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2024/1/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	nums1, nums2 := []int64{}, []int64{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		numsStr := strings.Split(string(bytes), "   ")
		fmt.Println(numsStr)
		num1, _ := strconv.ParseInt(numsStr[0], 10, 64)
		num2, _ := strconv.ParseInt(numsStr[1], 10, 64)

		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)

	}

	slices.Sort(nums1)
	slices.Sort(nums2)

	diff := int64(0)

	for i := range nums1 {
		diff += int64(math.Abs(float64(nums1[i] - nums2[i])))
	}

	fmt.Println(diff)
}
