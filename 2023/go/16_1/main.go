package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type direction uint

const (
	up direction = iota
	down
	left
	right
)

const (
	emptySpace        = '.'
	leftMirror        = '\\'
	rightMirror       = '/'
	upDownSplitter    = '|'
	leftRightSplitter = '-'
)

type pos struct {
	row, col int
	dir      direction
}

type tile struct {
	c         byte
	energized bool
	// for loop detection
	dirsTravelled map[direction]bool
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

	grid := [][]tile{}
	for s.Scan() {
		row := make([]tile, 0, len(s.Bytes()))

		for _, b := range s.Bytes() {
			row = append(row, tile{c: b, dirsTravelled: make(map[direction]bool)})
		}

		grid = append(grid, row)
	}

	queue := []pos{{0, -1, right}}

	var energizedCount uint
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		switch p.dir {
		case up:
			p.row--
			if p.row < 0 {
				continue
			}
		case down:
			p.row++
			if p.row == len(grid) {
				continue
			}
		case left:
			p.col--
			if p.col < 0 {
				continue
			}
		case right:
			p.col++
			if p.col == len(grid[p.row]) {
				continue
			}
		}

		if !grid[p.row][p.col].energized {
			grid[p.row][p.col].energized = true
			energizedCount++
		}
		if grid[p.row][p.col].dirsTravelled[p.dir] {
			continue
		}
		grid[p.row][p.col].dirsTravelled[p.dir] = true

		switch grid[p.row][p.col].c {
		case emptySpace:
			queue = append(queue, p)
			continue
		case upDownSplitter:
			if p.dir == up || p.dir == down {
				queue = append(queue, p)
				continue
			} else {
				p2 := p
				p.dir = up
				p2.dir = down
				queue = append(queue, p)
				queue = append(queue, p2)
			}
		case leftRightSplitter:
			if p.dir == left || p.dir == right {
				queue = append(queue, p)
				continue
			} else {
				p2 := p
				p.dir = left
				p2.dir = right
				queue = append(queue, p)
				queue = append(queue, p2)
			}
		case leftMirror:
			switch p.dir {
			case up:
				p.dir = left
			case down:
				p.dir = right
			case left:
				p.dir = up
			case right:
				p.dir = down
			}
			queue = append(queue, p)
		case rightMirror:
			switch p.dir {
			case up:
				p.dir = right
			case down:
				p.dir = left
			case left:
				p.dir = down
			case right:
				p.dir = up
			}
			queue = append(queue, p)
		}
	}

	fmt.Println(energizedCount)
}
