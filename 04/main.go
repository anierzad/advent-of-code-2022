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

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		elves := strings.Split(line, ",")

		aMin, aMax, err := limits(elves[0])
		if err != nil {
			log.Fatal(err)
		}

		bMin, bMax, err := limits(elves[1])
		if err != nil {
			log.Fatal(err)
		}

		if (bMin >= aMin && bMax <= aMax) || (aMin >= bMin && aMax <= bMax) {
			total++
		}
	}

	fmt.Println("Total:", total)
}

func limits(rangeStr string) (min, max int, err error) {

	parts := strings.Split(rangeStr, "-")

	min, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	max, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}

	return
}
