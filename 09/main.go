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
	Previous Point
	child    *Node
	Visited  map[Point]bool
}

func (n *Node) Attach(child *Node) {
	n.child = child
}

func (n *Node) Move(direction rune) {

	n.Previous = n.Position

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
	if !n.Position.Touches(parent.Position) {
		n.Position = parent.Previous
		n.Visited[n.Position] = true
	}
}

func NewNode(symbol rune) Node {
	return Node{
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

	head := NewNode('H')
	tail := NewNode('T')
	head.Attach(&tail)

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
