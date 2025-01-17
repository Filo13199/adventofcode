package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/filo/Documents/adventofcode/2024/3/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	str := ""
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		str += string(bytes)
	}

	r, _ := regexp.Compile(`mul\([0-9][0-9]*[0-9]*,[0-9][0-9]*[0-9]*\)`)
	locs := r.FindAllStringIndex(str, -1)
	sum := int64(0)
	for i := range locs {
		substr := str[locs[i][0]:locs[i][1]]
		substr = strings.TrimLeft(substr, "mul(")
		substr = strings.TrimRight(substr, ")")
		numsstr := strings.Split(substr, ",")

		num1, _ := strconv.ParseInt(numsstr[0], 10, 64)
		num2, _ := strconv.ParseInt(numsstr[1], 10, 64)
		sum += num1 * num2
	}
	fmt.Println(sum)
}
