package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

type Tree struct {
	Position Point
	Height   int
}

func (t *Tree) viewNorth(forest Forest) (distance int, blocked bool) {
	for y := t.Position.Y-1;; y-- {
		p := NewPoint(t.Position.X, y)

		ct, exists := forest[p]
		if !exists {
			break
		}

		distance++
		if ct.Height >= t.Height {
			return distance, true
		}
	}

	return distance, false
}

func (t *Tree) viewEast(forest Forest) (distance int, blocked bool) {
	for x := t.Position.X+1;; x++ {
		p := NewPoint(x, t.Position.Y)

		ct, exists := forest[p]
		if !exists {
			break
		}

		distance++
		if ct.Height >= t.Height {
			return distance, true
		}
	}

	return distance, false
}

func (t *Tree) viewSouth(forest Forest) (distance int, blocked bool) {
	for y := t.Position.Y+1;; y++ {
		p := NewPoint(t.Position.X, y)

		ct, exists := forest[p]
		if !exists {
			break
		}

		distance++
		if ct.Height >= t.Height {
			return distance, true
		}
	}

	return distance, false
}

func (t *Tree) viewWest(forest Forest) (distance int, blocked bool) {
	for x := t.Position.X-1;; x-- {
		p := NewPoint(x, t.Position.Y)

		ct, exists := forest[p]
		if !exists {
			break
		}

		distance++
		if ct.Height >= t.Height {
			return distance, true
		}
	}

	return distance, false
}

func (t *Tree) Visible(forest Forest) bool {
	_, visN := t.viewNorth(forest)
	_, visE := t.viewEast(forest)
	_, visS := t.viewSouth(forest)
	_, visW := t.viewWest(forest)

	return !visN || !visE || !visS || !visW
}

func (t *Tree) ScenicScore(forest Forest) int {
	visN, _ := t.viewNorth(forest)
	visE, _ := t.viewEast(forest)
	visS, _ := t.viewSouth(forest)
	visW, _ := t.viewWest(forest)

	return visN * visE * visS * visW
}

type Forest map[Point]Tree

func (f *Forest) AddTree(x, y, height int) {
	p := NewPoint(x, y)
	(*f)[p] = Tree{
		Position: p,
		Height: height,
	}
}

func (f Forest) PrintVis() {
	var maxX, maxY int
	for p := range f {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	fmt.Println()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := NewPoint(x, y)
			t := f[p]
			v := t.Visible(f)
			r := 'O'
			if v {
				r = '*'
			}
			fmt.Print(string(r))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	forest := Forest{}

	scanner := bufio.NewScanner(f)
	yPos := 0
	for scanner.Scan() {
		line := scanner.Text()
		for xPos, h := range line {
			height, err := strconv.Atoi(string(h))
			if err != nil {
				log.Fatal(err)
			}
			forest.AddTree(xPos, yPos, height)
		}
		yPos++
	}

	total := 0
	for _, t := range forest {
		if t.Visible(forest) {
			total++
		}
	}

	forest.PrintVis()

	fmt.Println("Visible trees:", total)

	best := 0
	for _, t := range forest {
		score := t.ScenicScore(forest)
		if score > best {
			best = score
		}
	}

	fmt.Println("Best scenic score:", best)
}
