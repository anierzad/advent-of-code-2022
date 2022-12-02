package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var elfMost *Elf

	for _, elf := range elves {
		if elfMost == nil || elfMost.Calories() < elf.Calories() {
			elfMost = elf
		}
	}

	fmt.Printf("Most calories: %d\n", elfMost.Calories())
}
