package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rucksack struct {
	first  map[rune]int
	second map[rune]int
}

func (r Rucksack) InBoth() []rune {
	both := make([]rune, 0)

	for k := range r.first {
		_, e := r.second[k]
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
		first:  first,
		second: second,
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
		rucksack := NewRucksack(line)

		for _, r := range rucksack.InBoth() {
			total = total + itemPriority(r)
		}
	}

	fmt.Println("Total:", total)
}
