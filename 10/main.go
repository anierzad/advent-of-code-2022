package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	ExecTime int
	Action   func(reg int, args []string) int
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	operations := map[string]Operation{
		"addx": {
			ExecTime: 2,
			Action: func(reg int, args []string) int {
				if len(args) < 2 {
					log.Fatal("Not enough arguments passed to addx.")
				}

				val, err := strconv.Atoi(args[1])
				if err != nil {
					log.Fatal(err)
				}

				return reg + val
			},
		},
		"noop": {
			ExecTime: 1,
			Action: func(reg int, args []string) int {
				return reg
			},
		},
	}

	xReg := 1
	cycle := 0
	signals := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		op := operations[args[0]]

		for c := 0; c < op.ExecTime; c++ {
			cycle++

			if (cycle - 20) % 40 == 0 {
				signals = append(signals, xReg * cycle)
			}
		}

		xReg = op.Action(xReg, args)
	}

	total := 0
	for _, s := range signals {
		total = total + s
	}

	fmt.Println("Total signal strength:", total)
}
