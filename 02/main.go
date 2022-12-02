package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	SCORES = map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
)

type Game struct {
	Opposition rune
	Player rune
}

func (g Game) Score() int {

	total := SCORES[g.Player]

	if g.outcome() == 0 {
		total = total + 3
	}

	if g.outcome() == 1 {
		total = total + 6
	}

	return total
}

func (g Game) outcome() int {

	if SCORES[g.Player] == SCORES[g.Opposition] {
		return 0
	}

	if SCORES[g.Player] == SCORES['Z'] && SCORES[g.Opposition] == SCORES['A'] {
		return -1
	}

	if SCORES[g.Player] == SCORES['X'] && SCORES[g.Opposition] == SCORES['C'] {
		return 1
	}

	if SCORES[g.Player] > SCORES[g.Opposition] {
		return 1
	}

	return -1
}

func NewGame(opposition, player rune) Game {
	return Game{
		Opposition: opposition,
		Player: player,
	}
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	games := make([]Game, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, NewGame(rune(line[0]), rune(line[2])))
	}

	totalScore := 0
	for _, g := range games {
		totalScore += g.Score()
	}

	fmt.Println("Total score:", totalScore)
}
