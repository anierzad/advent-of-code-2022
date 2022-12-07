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
	MAX_SPACE = 70000000
	UPDATE_SPACE = 30000000
)

type Node struct {
	Name     string
	Contents map[string]*Node
	Parent   *Node
	size     int
}

func (n *Node) Size() int {
	total := n.size

	for _, v := range n.Contents {
		total = total + v.Size()
	}

	return total
}

func (n *Node) CreateDir(name string) {
	n.Contents[name] = &Node{
		Name:     name,
		Contents: make(map[string]*Node),
		Parent:   n,
	}
}

func (n *Node) CreateFile(name string, size int) {
	n.Contents[name] = &Node{
		Name: name,
		size: size,
	}
}

func (n *Node) Ls() {

	for _, v := range n.Contents {

		if v.Contents != nil {
			fmt.Printf("dir %s", v.Name)
			continue
		}

		fmt.Printf("%d %s", v.Size(), v.Name)
	}
}

func (n *Node) AllDirectories() []*Node {
	allDirs := make([]*Node, 0)

	for _, v := range n.Contents {
		if v.Contents != nil {
			allDirs = append(allDirs, v)
			allDirs = append(allDirs, v.AllDirectories()...)
		}
	}

	return allDirs
}

func main() {

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	root := Node{
		Contents: make(map[string]*Node),
	}
	pwd := &root

	scanner := bufio.NewScanner(f)

	var last string
	for scanner.Scan() {
		line := scanner.Text()

		// New command?
		if line[0] == '$' {
			args := strings.Split(line, " ")

			last = args[1]

			switch args[1] {
			case "cd":

				// Special case.
				if args[2] == ".." {
					pwd = pwd.Parent
					continue
				}

				pwd.CreateDir(args[2])
				pwd = pwd.Contents[args[2]]
			}
			continue
		}

		switch last {
		case "ls":
			args := strings.Split(line, " ")

			if args[0] == "dir" {
				pwd.CreateDir(args[1])
			} else {
				size, err := strconv.Atoi(args[0])
				if err != nil {
					log.Fatal(err)
				}
				pwd.CreateFile(args[1], size)
			}
		}
	}

	allDirs := root.AllDirectories()

	usedSpace := root.Size()
	available := MAX_SPACE - usedSpace
	required := UPDATE_SPACE - available

	best := MAX_SPACE

	for _, dir := range allDirs {
		size := dir.Size()
		if size > required && size < best {
			best = size
		}
	}

	fmt.Println("Smallest dir that's big enough:", best)
}
