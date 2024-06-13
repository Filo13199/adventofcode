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

func getLineCalibrationValue(line []byte) int {
	newStr := strings.TrimFunc(string(line), func(r rune) bool {
		return !(r >= '0' && r <= '9')
	})
	if len(newStr) == 0 {
		return -1
	}
	digits := "-1"
	if len(newStr) == 1 {
		digits = newStr + newStr
	} else if len(newStr) > 1 {
		digits = string(newStr[0]) + string(newStr[len(newStr)-1])
	}

	val, err := strconv.Atoi(digits)
	if err != nil {
		return -1
	}

	return val
}

func main() {
	file, err := os.Open("./calibrationdoc.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)
	sum := 0
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		v := getLineCalibrationValue(bytes)
		if v == -1 {
			log.Fatal("whatt ?")
		}
		sum += v
	}
	fmt.Println(sum)
}
