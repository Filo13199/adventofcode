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

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2024/1/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	nums1, nums2 := []int64{}, map[int64]int64{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		numsStr := strings.Split(string(bytes), "   ")
		num1, _ := strconv.ParseInt(numsStr[0], 10, 64)
		num2, _ := strconv.ParseInt(numsStr[1], 10, 64)

		nums1 = append(nums1, num1)
		nums2[num2] += 1

	}

	score := int64(0)

	for i := range nums1 {
		score += nums1[i] * nums2[nums1[i]]
	}

	fmt.Println(score)
}
