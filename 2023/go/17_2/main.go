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
	grid := map[coord]int{}

	finish := coord{0, 0}
	var i int
	for s.Scan() {
		finish.row = i
		t := s.Text()
		for j, r := range t {
			grid[coord{i, j}] = int(r - '0')
			finish.col = j
		}

		i++
	}

	visited := make(map[state]int)
	prioQ := posHeap{}
	heap.Push(&prioQ, pos{0, 1, state{dir: right}})
	heap.Push(&prioQ, pos{0, 1, state{dir: down}})

	for len(prioQ) > 0 {
		p := heap.Pop(&prioQ).(pos)

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

		val, ok := grid[p.coord]
		if !ok {
			continue
		}
		p.count += val

		if p.steps >= 4 {
			if c, ok := visited[p.state]; ok && c <= p.count {
				continue
			}
			visited[p.state] = p.count

			if p.coord == finish {
				fmt.Println(p.count)
				break
			}
			directions := dirsLeftRight
			if p.dir == up || p.dir == down {
				directions = dirsUpDown
			}

			for _, dir := range directions {
				heap.Push(&prioQ, pos{
					p.count,
					1,
					state{
						p.coord,
						dir,
					},
				})
			}
		}

		if p.steps <= 10 {
			p.steps++
			heap.Push(&prioQ, p)
		}
	}
}

type coord struct {
	row, col int
}

type state struct {
	coord
	dir direction
}

type pos struct {
	count int
	steps int
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
