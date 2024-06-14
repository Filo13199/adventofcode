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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)
	cards := []string{}
	sum := 0
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		cards = append(cards, string(bytes))
	}
	originalAndcopies := make([]int, len(cards))

	for i := range originalAndcopies {
		originalAndcopies[i] = 1
	}

	for c := range cards {
		split := strings.Split(cards[c], " | ")
		actualCardNumbers := strings.Split(strings.Split(split[0], ": ")[1], " ")
		winningCardNumbers := strings.Split(split[1], " ")
		winningArr := make([]bool, 100)
		n := 0
		for i := range winningCardNumbers {
			num, e := strconv.Atoi(winningCardNumbers[i])
			if e == nil {
				winningArr[num] = true
			}
		}
		for i := range actualCardNumbers {
			num, e := strconv.Atoi(actualCardNumbers[i])
			if winningArr[num] && e == nil {
				n++
			}
		}

		for ; n > 0; n-- {
			originalAndcopies[c+n] += originalAndcopies[c]
		}
		sum += originalAndcopies[c]
	}

	fmt.Println(sum)
}
