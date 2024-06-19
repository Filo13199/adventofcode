package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/06/sample.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	reader.Discard(5)
	bytes, _, _ := reader.ReadLine()
	timesstr := strings.Split(string(bytes), " ")
	times := []int{}
	for i := range timesstr {
		t, e := strconv.Atoi(timesstr[i])
		if e == nil {
			times = append(times, t)
		}
	}

	reader.Discard(9)

	bytes, _, _ = reader.ReadLine()

	distancesstr := strings.Split(string(bytes), " ")

	distances := []int{}

	for i := range distancesstr {
		d, e := strconv.Atoi(distancesstr[i])
		if e == nil {
			distances = append(distances, d)
		}
	}
	total := 1
	for i := range times {
		n := 0
		for timeHeld := 0; timeHeld <= times[i]; timeHeld++ {
			distanceCovered := timeHeld * (times[i] - timeHeld)
			if distanceCovered > distances[i] {
				n++
			}
		}
		if n > 0 {
			total *= n
		}
	}
	fmt.Println(total)
}
