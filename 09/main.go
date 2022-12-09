package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

const (
	ROPE_LENGTH = 10
)

type Point struct {
	X, Y int
}

func (p Point) Touches(op Point) bool {
	xDiff := int(math.Abs(float64(p.X - op.X)))
	yDiff := int(math.Abs(float64(p.Y - op.Y)))
	return xDiff < 2 && yDiff < 2
}

func NewPoint(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

type Node struct {
	Symbol   rune
	Position Point
	child    *Node
	Visited  map[Point]bool
}

func (n *Node) Attach(child *Node) {
	n.child = child
}

func (n *Node) Move(direction rune) {

	switch direction {
	case 'U':
		n.Position.Y++
	case 'D':
		n.Position.Y--
	case 'R':
		n.Position.X++
	case 'L':
		n.Position.X--
	}

	n.Visited[n.Position] = true

	if n.child != nil {
		n.child.Follow(n)
	}
}

func (n *Node) Follow(parent *Node) {
	if n.doFollow(parent) {
		n.Visited[n.Position] = true

		if n.child != nil {
			n.child.Follow(n)
		}
	}
}

func (n *Node) doFollow(parent *Node) bool {

	if n.Position.Touches(parent.Position) {
		return false
	}

	xDiff := parent.Position.X - n.Position.X
	yDiff := parent.Position.Y - n.Position.Y

	if xDiff != 0 {
		if xDiff > 0 {
			xDiff = 1
		} else {
			xDiff = -1
		}
	}

	if yDiff != 0 {
		if yDiff > 0 {
			yDiff = 1
		} else {
			yDiff = -1
		}
	}

	n.Position.X = n.Position.X + xDiff
	n.Position.Y = n.Position.Y + yDiff
	return true
}

func NewNode(symbol rune) *Node {
	return &Node{
		Symbol: symbol,
		Visited: map[Point]bool{
			NewPoint(0, 0): true,
		},
	}
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	knots := make([]*Node, ROPE_LENGTH)

	for i := 0; i < ROPE_LENGTH; i++ {
		if i == 0 {
			knots[i] = NewNode('H')
			continue
		}
		knots[i] = NewNode(rune(strconv.Itoa(i)[0]))
		knots[i-1].Attach(knots[i])
	}

	head := knots[0]
	tail := knots[len(knots)-1]

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		moveParts := strings.Split(scanner.Text(), " ")
		count, err := strconv.Atoi(moveParts[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < count; i++ {
			head.Move(rune(moveParts[0][0]))
		}
	}

	total := 0
	for _, v := range tail.Visited {
		if v {
			total++
		}
	}

	fmt.Println("Positions tail visited:", total)
}
