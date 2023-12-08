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
	nodefields string
	left       *Node
	right      *Node
	endsInZ    bool
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
	startDirection := &Direction{
		left: false,
	}
	insBytes := s.Bytes()
	if insBytes[0] == 'L' {
		startDirection.left = true
	}

	cur := startDirection
	for _, b := range insBytes[1:] {
		cur.next = &Direction{
			left: false,
		}
		if b == 'L' {
			cur.next.left = true
		}
		cur = cur.next
	}

	cur.next = startDirection

	// Prep nodes
	nodes := map[string]*Node{}
	positions := []*Node{}
	for s.Scan() {
		fields := strings.Split(s.Text(), " = ")

		fields[1] = strings.TrimLeft(fields[1], "(")
		fields[1] = strings.TrimRight(fields[1], ")")

		nodes[fields[0]] = &Node{
			nodefields: fields[1],
			endsInZ:    fields[0][2] == 'Z',
		}

		if fields[0][2] == 'A' {
			positions = append(positions, nodes[fields[0]])
		}
	}

	// connect nodes
	for _, node := range nodes {
		next := strings.Split(node.nodefields, ", ")
		node.left = nodes[next[0]]
		node.right = nodes[next[1]]
	}

	// LCM
	counts := make([]uint64, 0, len(positions))
	for _, pos := range positions {
		direction := startDirection
		var steps uint64
		for !pos.endsInZ {
			if direction.left {
				pos = pos.left
			} else {
				pos = pos.right
			}

			direction = direction.next
			steps++
		}
		counts = append(counts, steps)
	}

	result := counts[0]
	for _, count := range counts[1:] {
		result = result * count / gcd(result, count)
	}

	fmt.Println(result)
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
