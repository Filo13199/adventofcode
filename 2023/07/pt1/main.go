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

const (
	highCard = iota
	onePair
	twoPair
	threeOfKind
	FullHouse
	FourKind
	FiveKind
)

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
var cardsMap = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var handsMap = map[string]int{
	"11111": 0,
	"1112":  1,
	"122":   2,
	"113":   3,
	"23":    4,
	"14":    5,
	"5":     6,
}

type handWithBid struct {
	hand string
	bid  int
}

func (h handWithBid) getHandType() int {
	charsMap := map[byte]int{}
	for i := range h.hand {
		charsMap[h.hand[i]]++
	}
	vals := make([]int, 0, len(charsMap))
	for _, v := range charsMap {
		vals = append(vals, v)
	}
	slices.Sort(vals)
	s := ""
	for i := range vals {
		s += strconv.Itoa(vals[i])
	}
	return handsMap[s]
}

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/07/input.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}
	reader := bufio.NewReader(file)
	handsWithBids := []handWithBid{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		split := strings.Split(string(bytes), " ")
		bid, e := strconv.Atoi(split[1])
		if e != nil {
			fmt.Println(e)
		}
		h := handWithBid{hand: split[0], bid: bid}
		handsWithBids = append(handsWithBids, h)
	}

	slices.SortFunc(handsWithBids, func(a, b handWithBid) int {
		atype, btype := a.getHandType(), b.getHandType()
		if atype > btype {
			return 1
		}

		if btype > atype {
			return -1
		}

		for i := range a.hand {
			if cardsMap[a.hand[i]] > cardsMap[b.hand[i]] {
				return 1
			}
			if cardsMap[a.hand[i]] < cardsMap[b.hand[i]] {
				return -1
			}
		}
		return 0
	})

	sum := int64(0)
	for i := range handsWithBids {
		sum += int64(handsWithBids[i].bid * (i + 1))
	}
	fmt.Println(sum)
}