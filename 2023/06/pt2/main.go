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
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/06/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	reader.Discard(5)
	bytes, _, _ := reader.ReadLine()
	timesstr := strings.Split(string(bytes), " ")
	timestr := ""
	for i := range timesstr {
		_, e := strconv.Atoi(timesstr[i])
		if e == nil {
			timestr += timesstr[i]
		}
	}

	reader.Discard(9)

	bytes, _, _ = reader.ReadLine()

	distancesstr := strings.Split(string(bytes), " ")
	distancestr := ""
	for i := range distancesstr {
		_, e := strconv.Atoi(distancesstr[i])
		if e == nil {
			distancestr += distancesstr[i]
		}
	}

	time, _ := strconv.ParseInt(timestr, 10, 64)
	distance, _ := strconv.ParseInt(distancestr, 10, 64)
	total := 0

	for timeHeld := int64(0); timeHeld <= time; timeHeld++ {
		distanceCovered := timeHeld * (time - timeHeld)
		if distanceCovered > distance {
			total++
		}
	}
	fmt.Println(total)
}
