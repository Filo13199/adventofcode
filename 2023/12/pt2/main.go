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
		if !strings.Contains(springs, "#") && len(RL) == 0 {
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

func getAllPossibleCombinations(springs []byte, i int, pipes int, RL []int, stringsMap map[string]int) int {
	if i == len(springs)-1 && len(RL)==1 && (i-pipes)+1==RL[0]{
		return 1
	}
	if len(springs) == 0 && len(RL) == 0 {
		return 1
	}
	if len(springs) == 0 || len(RL) == 0 || i >= len(springs) {
		return 0
	}

	if _, ok := stringsMap[string(springs)]; ok {
		return 0
	}

	if springs[i] == '#' {
		if pipes == -1 {
			pipes = i
		}
		return getAllPossibleCombinations(springs, i+1, pipes, RL, stringsMap)
	} else if springs[i] == '.' && pipes > -1 && (i-pipes) != RL[0] {
		return 0
	} else if springs[i] == '.' && pipes > -1 {
		RL = RL[1:]
		x := getAllPossibleCombinations(springs[i:], 0, -1, RL, stringsMap)
		return x
	}else if springs[i]=='.' && pipes==-1{
		return getAllPossibleCombinations(springs,i+1,pipes,RL,stringsMap)
	}
	s1 := slices.Clone(springs)
	s1[i]='.'
	s2 := slices.Clone(springs)
	s2[i]='#'
	
	p1 := getAllPossibleCombinations(s1, i, pipes, RL, stringsMap)
	p2 := getAllPossibleCombinations(s2, i, pipes, RL, stringsMap)
	stringsMap[string(springs)] = p1 +p2
	return p1 + p2
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
		rl := string(bytes[ind+1:])
		//RLString := strings.Split(fmt.Sprintf("%s,%s,%s,%s,%s", rl, rl, rl, rl, rl), ",")
		RLString:=strings.Split(rl,",")
		RL := make([]int, len(RLString))

		for i := range RLString {
			n, _ := strconv.Atoi(RLString[i])
			RL[i] = n
		}
		duplicatedSprings := make([]byte, 0, len(springs)*5+5)
		for i := 1; i <= 1; i++ {
			duplicatedSprings = append(duplicatedSprings, springs...)
			// if i < 5 {
			// 	duplicatedSprings = append(duplicatedSprings, '?')
			// }
		}
		x := getAllPossibleCombinations(duplicatedSprings, 0, -1, RL, map[string]int{})
		sum += x
	}
	fmt.Println(sum)
}
