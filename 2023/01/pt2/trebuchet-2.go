package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getLineCalibrationValue(line string) int {
	newStr := strings.TrimFunc(string(line), func(r rune) bool {
		return !(r >= '0' && r <= '9')
	})
	if len(newStr) == 0 {
		return -1
	}
	digits := "-1"
	if len(newStr) == 1 {
		digits = newStr + newStr
	} else if len(newStr) > 1 {
		digits = string(newStr[0]) + string(newStr[len(newStr)-1])
	}

	val, err := strconv.Atoi(digits)
	if err != nil {
		return -1
	}

	return val
}

func main() {
	file, err := os.Open("./calibrationdoc.txt")
	if err != nil {
		log.Fatal("error openiing file", err)
	}

	reader := bufio.NewReader(file)

	// one
	// two
	// three
	// four
	// five
	// six
	// seven
	// eight
	// nine
	// zero

	sum := 0
	regex := regexp.MustCompile(`\zero|eight|one|two|three|four|five|six|seven|nine`)
	replMap := map[string]byte{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for {
		bytes, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		loc := regex.FindIndex(bytes)
		if loc != nil && len(loc) == 2 && len(bytes) > 2 {
			tmpBytes := []byte{}
			tmpBytes = append(tmpBytes, bytes[0:loc[0]]...)
			tmpBytes = append(tmpBytes, replMap[string(bytes[loc[0]:loc[1]])])
			tmpBytes = append(tmpBytes, bytes[loc[1]:]...)
			bytes = tmpBytes
		} else {
		}
		str := string(bytes)
		for i := len(bytes) - 1; i >= 0; i-- {
			// strsofar := string(bytes[i:])
			tmpB := bytes[i:]
			strB := string(tmpB)
			if loc := regex.FindIndex(tmpB); loc != nil && len(loc) == 2 {
				tmpBytes := []byte{}
				tmpBytes = append(tmpBytes, replMap[string(tmpB[loc[0]:loc[1]])])
				tmpBytes = append(tmpBytes, tmpB[loc[1]:]...)
				tstr := string(tmpBytes)
				str = strings.Replace(str, strB, tstr, 1)
				break
			}
		}
		v := getLineCalibrationValue(str)
		if v == -1 {
			log.Fatal("whatt ?")
		}
		sum += v
	}
	fmt.Println(sum)
}
