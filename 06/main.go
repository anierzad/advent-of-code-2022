package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	MARKER_LENGTH = 4
)

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	var readErr error
	offset := 0
	for readErr == nil {
		var next []byte
		next, readErr = reader.Peek(MARKER_LENGTH)
		if allUnique(next) {
			break
		}
		reader.Discard(1)
		offset++
	}

	fmt.Println("Stopped:", offset + MARKER_LENGTH)
}

func allUnique(dat []byte) bool {

	tally := make(map[byte]int)

	for _, b := range dat {
		val := tally[b]
		if val > 0 {
			return false
		}
		tally[b] = val + 1
	}

	return true
}
