package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	CRT_WIDTH = 40
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		op := operations[args[0]]

		for c := 0; c < op.ExecTime; c++ {
			cycle++
			position := (cycle - 1) % CRT_WIDTH
			symbol := '.'

			// fmt.Println(cycle, position)

			if position == 0 {
				fmt.Println()
				fmt.Printf("Cycle %3d -> ", cycle)
			}

			if position >= xReg - 1 && position <= xReg + 1 {
				symbol = '#'
			}

			fmt.Print(string(symbol))
		}

		xReg = op.Action(xReg, args)
	}
	fmt.Println()
}
