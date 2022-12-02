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
	}

	RESULT = map[rune]int{
		'X': -1,
		'Y': 0,
		'Z': 1,
	}
)

type Game struct {
	Opposition rune
	Result rune
}

func (g Game) Score() int {

	total := SCORES[g.player()]

	if RESULT[g.Result] == 0 {
		total = total + 3
	}

	if RESULT[g.Result] == 1 {
		total = total + 6
	}

	return total
}

func (g Game) player() rune {

	if RESULT[g.Result] == -1 {
		player := g.Opposition - 1

		if player < 'A' {
			player = 'C'
		}
		return player
	}

	if RESULT[g.Result] == 1 {
		player := g.Opposition + 1

		if player > 'C' {
			player = 'A'
		}
		return player
	}

	return g.Opposition
}


func NewGame(opposition, result rune) Game {
	return Game{
		Opposition: opposition,
		Result: result,
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
