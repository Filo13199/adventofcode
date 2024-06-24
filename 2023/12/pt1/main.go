package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkIfValidCombination(springs string, RL []int) int {
	springs = strings.TrimLeft(springs, ".")
	for {
		if len(RL) == 0 && strings.Contains(springs, "#") {
			return 0
		}
		if (!strings.Contains(springs, "#") && len(RL) == 0) {
			return 1
		}

		count := 0
		for i := range springs {
			if springs[i] == '#' {
				count++
			} else {
				break
			}
		}

		if count != RL[0] {
			return 0
		}

		springs = strings.TrimLeft(springs, "#")
		springs = strings.TrimLeft(springs, ".")
		RL = RL[1:]
	}
}

func getAllPossibleCombinations(springs []byte, RL []int, stringsMap map[string]int) int {
	if len(springs) == 0 || len(RL) == 0 {
		return 0
	}

	if _, ok := stringsMap[string(springs)]; ok {
		return 0
	}

	firstOccurence := -1
	for i, c := range springs {
		if c == '?' {
			if firstOccurence == -1 {
				firstOccurence = i
			}
		}
	}

	if firstOccurence == -1 {
		n := checkIfValidCombination(string(springs),RL)
		stringsMap[string(springs)] = n
		return n
	}

	s1 := slices.Clone(springs)
	s2 := slices.Clone(springs)

	s1[firstOccurence] = '.'
	s2[firstOccurence] = '#'
	return getAllPossibleCombinations(s1, RL,stringsMap) + getAllPossibleCombinations(s2, RL,stringsMap)
}

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/12/sample.txt")
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

		ind := slices.Index(bytes, ' ')
		springs := bytes[:ind]
		RLString := strings.Split(string(bytes[ind+1:]), ",")
		RL := make([]int, len(RLString))

		for i := range RLString {
			n, _ := strconv.Atoi(RLString[i])
			RL[i] = n
		}
		x :=getAllPossibleCombinations(springs,RL,map[string]int{})
		sum +=x
	}
	fmt.Println(sum)
}
