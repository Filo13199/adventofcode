package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
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
	doR, _ := regexp.Compile(`do\(\)`)
	dontR, _ := regexp.Compile(`don\'t\(\)`)
	locs := r.FindAllStringIndex(str, -1)
	sum := int64(0)
	dosLocs := doR.FindAllStringIndex(str, -1)
	dontLocs := dontR.FindAllStringIndex(str, -1)

	fmt.Println(dosLocs)
	fmt.Println(dontLocs)
	slices.SortFunc(dosLocs, func(d1, d2 []int) int {
		if d1[1] >= d2[1] {
			return 1
		} else {
			return -1
		}
	})

	slices.SortFunc(dontLocs, func(dn1, dn2 []int) int {
		if dn1[1] >= dn2[1] {
			return 1
		} else {
			return -1
		}
	})

	do := true

	for i := range locs {
		latestDont := []int{}
		for j := range dontLocs {
			if dontLocs[j][1] > locs[i][0] {
				break
			}
			latestDont = dontLocs[j]
		}

		latestDo := []int{}
		for j := range dosLocs {
			if dosLocs[j][1] > locs[i][0] {
				break
			}
			latestDo = dosLocs[j]
		}

		do = (do && len(latestDont) == 0) || len(latestDont) == 0 && len(latestDo) > 0 || (len(latestDo) > 0 && latestDo[0] > latestDont[0])
		if !do {
			continue
		}
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
