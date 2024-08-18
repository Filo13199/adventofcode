package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

type myDish [][]byte

var display bool = false

func Display(d [][]byte) {
	if !display {
		return
	}
	for i := range d {
		fmt.Println(string(d[i]))
	}
	return
}

func tiltNorth(dish [][]byte) {
	for j := 0; j < len(dish[0]); j++ {
		incr := 0
		for i := 0; i < len(dish); i++ {
			if dish[i][j] == '.' {
				incr++
			} else if dish[i][j] == '#' {
				incr = 0
			} else if dish[i][j] == 'O' && incr > 0 {
				dish[i-incr][j] = 'O'
				dish[i][j] = '.'
			}
		}
	}
}

func tiltWest(dish [][]byte) {
	for i := 0; i < len(dish); i++ {
		incr := 0
		for j := 0; j < len(dish[i]); j++ {
			if dish[i][j] == '.' {
				incr++
			} else if dish[i][j] == '#' {
				incr = 0
			} else if dish[i][j] == 'O' && incr > 0 {
				dish[i][j-incr] = 'O'
				dish[i][j] = '.'
			}
		}
	}
}

func tiltSouth(dish [][]byte) {
	for j := 0; j < len(dish[0]); j++ {
		incr := 0
		for i := len(dish) - 1; i > -1; i-- {
			if dish[i][j] == '.' {
				incr++
			} else if dish[i][j] == '#' {
				incr = 0
			} else if dish[i][j] == 'O' && incr > 0 {
				dish[i+incr][j] = 'O'
				dish[i][j] = '.'
			}
		}
	}
}

func tiltEast(dish [][]byte) {
	for i := 0; i < len(dish); i++ {
		incr := 0
		for j := len(dish[i]) - 1; j > -1; j-- {
			if dish[i][j] == '.' {
				incr++
			} else if dish[i][j] == '#' {
				incr = 0
			} else if dish[i][j] == 'O' && incr > 0 {
				dish[i][j+incr] = 'O'
				dish[i][j] = '.'
			}
		}
	}
}

func getHash(dish [][]byte) string {
	hash := ""
	for i := range dish {
		hash += string(dish[i])
	}
	return hash
}

func deepCopy(dish [][]byte) [][]byte {
	originalDish := make([][]byte, len(dish))
	for i := range dish {
		originalDish[i] = make([]byte, len(dish[i]))
		for j := range dish[i] {
			originalDish[i][j] = dish[i][j]
		}
	}
	return originalDish
}

func solve(dish [][]byte) {
	vals := map[string]int{}
	keys := []string{}
	originalDish := make([][]byte, len(dish))

	for i := range dish {
		originalDish[i] = make([]byte, len(dish[i]))
		for j := range dish[i] {
			originalDish[i][j] = dish[i][j]
		}
	}

	lastUniqueCycle := 1
	for i := 1; i < 113; i++ {
		Display(dish)
		tiltNorth(dish)
		Display(dish)
		tiltWest(dish)
		Display(dish)
		tiltSouth(dish)
		Display(dish)
		tiltEast(dish)
		Display(dish)

		x := calcWeight(dish)
		hash := getHash(dish)
		// firstRept := true
		if _, ok := vals[hash]; ok {
			fmt.Println(i, vals[hash])
		} else if !ok {
			vals[hash] = x
			lastUniqueCycle = i
			keys = append(keys, hash)
			fmt.Println(i, vals[hash])
			// fmt.Printf("Cycle #%d\n", i)
		}
		// fmt.Printf("CYCLE #%d done ! \n", i)
	}
	fmt.Println(lastUniqueCycle)
	// fmt.Println(len(vals))
	// t := 224 % (lastUniqueCycle - 1)

	// fmt.Println(vals[keys[t]])
	// fmt.Println(vals[keys[109]])
	// fmt.Println(vals[keys[110]])
	// fmt.Println(vals[keys[111]])
}

func calcWeight(dish [][]byte) int {
	weight := 0
	for i := range dish {
		c := 0
		for j := range dish[i] {
			if dish[i][j] == 'O' {
				c++
			}
		}
		weight += (c * (len(dish) - i))
	}
	return weight
}

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/14/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	load := 0
	dish := [][]byte{}
	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		dish = append(dish, slices.Clone(bytes))
	}
	// tilt north
	display = false
	solve(dish)
	// for j := 0; j < len(dish[0]); j++ {
	// 	incr := 0
	// 	for i := 0; i < len(dish); i++ {
	// 		if dish[i][j] == '.' {
	// 			incr++
	// 		} else if dish[i][j] == '#' {
	// 			incr = 0
	// 		} else if dish[i][j] == 'O' {
	// 			load += (incr + (len(dish) - i))
	// 		}
	// 	}
	// }
	fmt.Println()
	fmt.Println(load)
	fmt.Println(1000000000 % 14)
}
