package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rucksack struct {
	First    map[rune]int
	Second   map[rune]int
}

func (r Rucksack) InBoth() []rune {
	both := make([]rune, 0)

	for k := range r.First {
		_, e := r.Second[k]
		if e {
			both = append(both, k)
		}
	}

	return both
}

func NewRucksack(contents string) Rucksack {
	half := len(contents) / 2

	first := make(map[rune]int)
	for _, r := range contents[:half] {
		val, exist := first[r]
		if !exist {
			first[r] = 0
		}
		first[r] = val + 1
	}
	
	second := make(map[rune]int)
	for _, r := range contents[half:] {
		val, exist := second[r]
		if !exist {
			second[r] = 0
		}
		second[r] = val + 1
	}

	return Rucksack{
		First:    first,
		Second:   second,
	}
}

func itemPriority(item rune) int {

	var priority int

	if item >= 65 && item <= 90 {
		priority = (int(item) - 64) + 26
	}

	if item >= 97 && item <= 122 {
		priority = int(item) - 96
	}

	return priority
}

func findCommon(rucksacks []Rucksack) rune {
	rucksackTally := make(map[rune]int)

	for _, r := range rucksacks {
		contents := make(map[rune]bool)
		for k := range r.First {
			contents[k] = true
		}
		for k := range r.Second {
			contents[k] = true
		}

		for k := range contents {
			val, exists := rucksackTally[k]
			if !exists {
				rucksackTally[k] = 0
			}
			rucksackTally[k] = val + 1

			if rucksackTally[k] == 3 {
				return k
			}
		}
	}

	return '!'
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	group := make([]Rucksack, 0)
	for scanner.Scan() {
		line := scanner.Text()
		group = append(group, NewRucksack(line))

		if len(group) < 3 {
			continue
		}

		badge := findCommon(group)
		total = total + itemPriority(badge)

		group = make([]Rucksack, 0)
	}

	fmt.Println("Total:", total)
}
