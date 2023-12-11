package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	northSouth byte = '|'
	eastWest   byte = '-'
	northEast  byte = 'L'
	northWest  byte = 'J'
	southWest  byte = '7'
	southEast  byte = 'F'
	ground     byte = '.'
	start      byte = 'S'
)

type direction int

const (
	south direction = iota
	east
	west
	north
)

type point struct {
	x, y int
}

type position struct {
	point
	dir direction
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

	// build maze and starting pos
	pos := position{}

	maze := make([][]byte, 0)
	var i int
	for s.Scan() {
		maze = append(maze, []byte(s.Text()))
		if pos.x == 0 && pos.y == 0 {
			for j, b := range s.Bytes() {
				if b == start {
					pos.x = i
					pos.y = j
					break
				}
			}
		}
		i++
	}

	if pos.x > 0 && maze[pos.x-1][pos.y] != ground {
		pos.dir = north
	} else if pos.x < len(maze) && maze[pos.x+1][pos.y] != ground {
		pos.dir = south
	} else {
		pos.dir = east
	}

	var mazeLength int
	var mazePoints []point

LOOP:
	for {
		mazePoints = append(mazePoints, pos.point)

		switch pos.dir {
		case north:
			pos.x--
		case south:
			pos.x++
		case east:
			pos.y++
		case west:
			pos.y--
		}
		if pos.x < 0 || pos.x == len(maze) || pos.y < 0 || pos.y == len(maze[pos.x]) {
			panic("dead end")
		}

		switch maze[pos.x][pos.y] {
		case ground:
			panic("dead end")
		case start:
			// finish!
			break LOOP
		case northSouth:
			if pos.dir == east || pos.dir == west {
				panic("dead end")
			}
		case eastWest:
			if pos.dir == north || pos.dir == south {
				panic("dead end")
			}
		case northEast:
			if pos.dir == north || pos.dir == east {
				panic("dead end")
			}
			if pos.dir == south {
				pos.dir = east
			} else {
				pos.dir = north
			}
		case northWest:
			if pos.dir == west || pos.dir == north {
				panic("dead end")
			}
			if pos.dir == south {
				pos.dir = west
			} else {
				pos.dir = north
			}
		case southWest:
			if pos.dir == south || pos.dir == west {
				panic("dead end")
			}
			if pos.dir == east {
				pos.dir = south
			} else {
				pos.dir = west
			}
		case southEast:
			if pos.dir == south || pos.dir == east {
				panic("dead end")
			}
			if pos.dir == north {
				pos.dir = east
			} else {
				pos.dir = south
			}
		}

		mazeLength += 1
	}

	var area int
	for i := 0; i < len(mazePoints); i++ {
		next := mazePoints[(i+1)%len(mazePoints)]
		area += mazePoints[i].x*next.y - mazePoints[i].y*next.x
	}

	fmt.Println((area / 2) - len(mazePoints)/2 + 1)
}
