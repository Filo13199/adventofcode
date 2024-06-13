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

func powerset(sessions []string) int {
	max := make(map[string]int)
	for i := range sessions {
		balls := make(map[string]int)
		grabs := strings.Split(sessions[i], ", ")
		for j := range grabs {
			grab := strings.Split(grabs[j], " ")
			n, _ := strconv.Atoi(grab[0])
			balls[grab[1]] += n
		}
		for color, v := range balls {
			if val := max[color]; v > val {
				max[color] = v
			}
		}
	}

	powerset := max["red"] * max["green"] * max["blue"]

	return powerset
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
		sum += powerset(sessions)
		game++
	}
	fmt.Println(sum)
}
