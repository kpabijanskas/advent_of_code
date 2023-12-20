package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
)

type direction uint8

const (
	up direction = iota
	down
	left
	right
)

var (
	dirsUpDown    = []direction{left, right}
	dirsLeftRight = []direction{up, down}
)

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)
	var grid [][]int

	for s.Scan() {
		t := s.Text()
		row := make([]int, 0, len(s.Text()))
		for _, r := range t {
			row = append(row, int(r-'0'))
		}

		grid = append(grid, row)
	}

	visited := make(map[state]int)
	prioQ := posHeap{
		pos{0, state{0, 0, right}},
		pos{0, state{0, 0, down}},
	}

	for len(prioQ) > 0 {
		p := heap.Pop(&prioQ).(pos)
		if p.row == len(grid)-1 && p.col == len(grid[p.row])-1 {
			fmt.Println(p.count)
			break
		}

		for i := 1; i <= 3; i++ {
			switch p.dir {
			case up:
				p.row--
			case down:
				p.row++
			case left:
				p.col--
			case right:
				p.col++
			}

			if p.row < 0 || p.row >= len(grid) || p.col < 0 || p.col >= len(grid[p.row]) {
				break
			}
			p.count += grid[p.row][p.col]

			if c, ok := visited[p.state]; ok && c < p.count {
				continue
			}
			visited[p.state] = p.count

			directions := dirsLeftRight
			if p.dir == up || p.dir == down {
				directions = dirsUpDown
			}

			for _, dir := range directions {
				heap.Push(&prioQ, pos{
					p.count,
					state{
						p.row,
						p.col,
						dir,
					},
				})
			}
		}
	}

}

type state struct {
	row, col int
	dir      direction
}

type pos struct {
	count int
	state
}

type posHeap []pos

func (p posHeap) Len() int {
	return len(p)
}

func (p posHeap) Less(i, j int) bool {
	return p[i].count < p[j].count
}

func (p posHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *posHeap) Push(x any) {
	*p = append(*p, x.(pos))
}

func (p *posHeap) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}
