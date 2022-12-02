package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Meals []int
}

func (e *Elf) AddMeal(meal int) {
	e.Meals = append(e.Meals, meal)
}

func (e Elf) Calories() int {
	total := 0

	for _, m := range e.Meals {
		total = total + m
	}

	return total
}

func NewElf() *Elf {
	return &Elf{
		Meals: []int{},
	}
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	elves := make([]*Elf, 0)
	var elf *Elf
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "" {
			elf = nil
			continue
		}

		if elf == nil {
			elf = NewElf()
			elves = append(elves, elf)
		}

		mealCals, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		elf.AddMeal(mealCals)
	}

	elvesMost := make([]*Elf, 0)

	for _, elf := range elves {

		// Add if slice not full.
		if len(elvesMost) < 3 {
			elvesMost = append(elvesMost, elf)
			continue
		}

		// Compare.
		for _, elfMost := range elvesMost {
			if elfMost.Calories() < elf.Calories() {

				elvesMost = append(elvesMost, elf)
				sort.Slice(elvesMost, func(i, j int) bool {
					return elvesMost[i].Calories() < elvesMost[j].Calories()
				})
				elvesMost = elvesMost[1:]
				break
			}
		}
	}

	total := 0
	for _, e := range elvesMost {
		total = total + e.Calories()
	}

	fmt.Printf("Total of 3 elves with most calories: %d\n", total)
}
