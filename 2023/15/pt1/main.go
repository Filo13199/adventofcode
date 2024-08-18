package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/Users/filo/Documents/GitHub/adventofcode/2023/15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	sum := int32(0)
	for {
		cr := int32(0)
		bytes, err := reader.ReadSlice(',')
		if errors.Is(err, io.EOF) {
			for i := range bytes {
				if bytes[i]==','{break}
				if bytes[i]=='\n'{continue}
				b :=bytes[i]
				cr += int32(b)
				cr *= 17
				cr =cr%256
			}
			sum+=cr
			break
		}
		for i := range bytes {
			if bytes[i]==','{break}
			if bytes[i]=='\n'{continue}
			b :=bytes[i]
			cr += int32(b)
			cr *= 17
			cr =cr%256
		}
		sum += cr
	}

	fmt.Println(sum)
}
