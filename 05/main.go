package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Stack []rune

func (s *Stack) Pop() rune {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Push(r rune) {
	*s = append((*s), r)
}

func (s Stack) Print() {
	for _, r := range s {
		fmt.Println(string(r))
	}
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	initial := make([]string, 0)
	moves := make([]string, 0)
	target := &initial

	for scanner.Scan() {
		line := scanner.Text()
		
		if line == "" {
			target = &moves
			continue
		}

		*target = append(*target, line)
	}

	stacks := make(map[rune]*Stack)
	stackPos := make(map[rune]int)

	for i := range initial {
		l := initial[(len(initial)-1)-i]

		if i == 0 {
			for p, r := range l {
				if r != ' ' {
					stacks[r] = &Stack{}
					stackPos[r] = p
				}
			}
			continue
		}

		for r, p := range stackPos {
			if l[p] != ' ' {
				stacks[r].Push(rune(l[p]))
			}
		}
	}

	reg := regexp.MustCompile(`move (?P<count>[\d]+) from (?P<source>[\d]+) to (?P<destination>[\d]+)`)
	for _, m := range moves {
		res := reg.FindStringSubmatch(m)

		count, err := strconv.Atoi(res[1])
		if err != nil {
			log.Fatal(err)
		}

		source := rune(res[2][0])
		destination := rune(res[3][0])

		for i := 0; i < count; i++ {
			t := stacks[source].Pop()
			stacks[destination].Push(t)
		}
	}

	stackIds := make([]rune, 0)
	for k := range stacks {
		stackIds = append(stackIds, k)
	}

	sort.Slice(stackIds, func(i, j int) bool {
		return stackIds[i] < stackIds[j]
	})

	for _, sid := range stackIds {
		s := stacks[sid]
		fmt.Print(string((*s)[len(*s)-1]))
	}
	fmt.Println()
}
