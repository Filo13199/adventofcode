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
		split := strings.Split(string(bytes), " | ")
		actualCardNumbers := strings.Split(strings.Split(split[0], ": ")[1], " ")
		winningCardNumbers := strings.Split(split[1], " ")
		winningArr := make([]bool, 100)
		w := 1
		val := 0
		for i := range winningCardNumbers {
			n, e := strconv.Atoi(winningCardNumbers[i])
			if e == nil {
				winningArr[n] = true
			}
		}
		for i := range actualCardNumbers {
			n, e := strconv.Atoi(actualCardNumbers[i])
			if winningArr[n] && e == nil {
				val = w
				w *= 2
			}
		}
		sum += val
	}

	fmt.Println(sum)
}
