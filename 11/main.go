package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []int
	operation   func(old int) int
	divisible   int
	pass        int
	fail        int
	Inspections int
}

func (m *Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) InspectItems(monkeys map[int]*Monkey) {
	for _, item := range m.items {
		m.Inspections++

		worry := m.operation(item)
		worry = int(math.Floor(float64(worry) / 3))

		if worry % m.divisible == 0 {
			monkeys[m.pass].AddItem(worry)
			continue
		}

		monkeys[m.fail].AddItem(worry)
	}

	m.items = make([]int, 0)
}

func NewMonkey(op func(old int) int, divisible, pass, fail int) *Monkey {
	return &Monkey{
		items:     []int{},
		operation: op,
		divisible: divisible,
		pass:      pass,
		fail:      fail,
	}
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	monkeys := make(map[int]*Monkey)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Monkey") {

			line = strings.ReplaceAll(line, ":", "")
			parts := strings.Split(line, " ")
			monkeyNo, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal(err)
			}

			scanner.Scan()
			line = scanner.Text()
			itemsStr := strings.TrimSpace(strings.ReplaceAll(line, "Starting items:", ""))
			startItems := strings.Split(itemsStr, ", ")

			scanner.Scan()
			line = scanner.Text()
			opStr := strings.TrimSpace(strings.ReplaceAll(line, "Operation:", ""))
			opParts := strings.Split(opStr, " ")

			var operation func(old int) int
			if opParts[3] == "*" {

				operation = func(old int) int {

					if opParts[4] == "old" {
						return old * old
					}

					arg, err := strconv.Atoi(opParts[4])
					if err != nil {
						log.Fatal(err)
					}

					return old * arg
				}
			}
			if opParts[3] == "+" {

				operation = func(old int) int {

					if opParts[4] == "old" {
						return old + old
					}

					arg, err := strconv.Atoi(opParts[4])
					if err != nil {
						log.Fatal(err)
					}

					return old + arg
				}
			}

			scanner.Scan()
			line = scanner.Text()
			divStr := strings.TrimSpace(strings.ReplaceAll(line, "Test: divisible by", ""))
			divisible, err := strconv.Atoi(divStr)
			if err != nil {
				log.Fatal(err)
			}

			scanner.Scan()
			line = scanner.Text()
			passStr := strings.TrimSpace(strings.ReplaceAll(line, "If true: throw to monkey", ""))
			pass, err := strconv.Atoi(passStr)
			if err != nil {
				log.Fatal(err)
			}

			scanner.Scan()
			line = scanner.Text()
			failStr := strings.TrimSpace(strings.ReplaceAll(line, "If false: throw to monkey", ""))
			fail, err := strconv.Atoi(failStr)
			if err != nil {
				log.Fatal(err)
			}

			monkey := NewMonkey(operation, divisible, pass, fail)

			for _, itemStr := range startItems {
				item, err := strconv.Atoi(itemStr)
				if err != nil {
					log.Fatal(err)
				}
				monkey.AddItem(item)
			}

			monkeys[monkeyNo] = monkey
		}
	}

	monkeyIds := make([]int, 0)
	for mid := range monkeys {
		monkeyIds = append(monkeyIds, mid)
	}
	sort.Ints(monkeyIds)

	for round := 0; round < 20; round++ {
		for _, mid := range monkeyIds {
			monkeys[mid].InspectItems(monkeys)
		}
	}

	inspections := make([]int, 0)
	for _, mid := range monkeyIds {
		fmt.Println("Monkey ", mid, "inspected items", monkeys[mid].Inspections, "times.")
		inspections = append(inspections, monkeys[mid].Inspections)
	}
	sort.Ints(inspections)

	fmt.Println("Monkey business:", inspections[len(inspections)-1] * inspections[len(inspections)-2])
}
