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

func isValidGame(sessions []string) bool {
	for i := range sessions {
		balls := make(map[string]int)
		balls["red"] = 12
		balls["green"] = 13
		balls["blue"] = 14
		grabs := strings.Split(sessions[i], ", ")
		for j := range grabs {
			grab := strings.Split(grabs[j], " ")
			n, _ := strconv.Atoi(grab[0])
			curr := balls[grab[1]]
			if curr-n < 0 {
				return false
			}
			balls[grab[1]] = balls[grab[1]] - n
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	sum := 0
	game := 1
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		_, str, _ := strings.Cut(string(bytes), ": ")
		sessions := strings.SplitN(str, "; ", -1)
		if isValidGame(sessions) {
			sum += game
		}
		game++
	}
	fmt.Println(sum)
}
