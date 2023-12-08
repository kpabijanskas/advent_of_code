package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	left bool
	next *Direction
}

type Node struct {
	left  string
	right string
}

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	data = bytes.TrimSpace(data)

	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)

	s.Scan()

	// prep directions
	nextDirection := &Direction{
		left: false,
	}
	insBytes := s.Bytes()
	if insBytes[0] == 'L' {
		nextDirection.left = true
	}

	cur := nextDirection
	for _, b := range insBytes[1:] {
		cur.next = &Direction{
			left: false,
		}
		if b == 'L' {
			cur.next.left = true
		}
		cur = cur.next
	}

	cur.next = nextDirection

	graph := map[string]Node{}
	for s.Scan() {
		fields := strings.Split(s.Text(), " = ")

		fields[1] = strings.TrimLeft(fields[1], "(")
		fields[1] = strings.TrimRight(fields[1], ")")

		next := strings.Split(fields[1], ", ")

		graph[fields[0]] = Node{
			left:  next[0],
			right: next[1],
		}
	}

	curNode := "AAA"
	var steps uint
	for curNode != "ZZZ" {
		if nextDirection.left {
			curNode = graph[curNode].left
		} else {
			curNode = graph[curNode].right
		}

		nextDirection = nextDirection.next
		steps++
	}

	fmt.Println(steps)
}
