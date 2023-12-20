package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	row, col int
}

func (p *pos) goLeft(n int) {
	p.col -= n
}

func (p *pos) goRight(n int) {
	p.col += n
}

func (p *pos) goUp(n int) {
	p.row -= n
}

func (p *pos) goDown(n int) {
	p.row += n
}

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)

	vertices := []pos{}

	cur := pos{}

	r := strings.NewReplacer("(", "", ")", "", "#", "")
	var boundary int
	for s.Scan() {
		fields := strings.Fields(s.Text())

		ds := r.Replace(fields[2])

		n, err := strconv.ParseInt(ds[0:5], 16, 64)
		if err != nil {
			panic(err)
		}

		boundary += int(n)

		switch ds[5] {
		case '2':
			cur.goLeft(int(n))
		case '0':
			cur.goRight(int(n))
		case '3':
			cur.goUp(int(n))
		case '1':
			cur.goDown(int(n))
		default:
			panic("UNKNOWN DIR")
		}

		vertices = append(vertices, cur)
	}
	boundary /= 2
	vertices = append(vertices, vertices[0])

	var inside int
	for i := range vertices[:len(vertices)-1] {
		a, b := vertices[i], vertices[i+1]
		inside += a.col*b.row - a.row*b.col
	}
	inside = abs(inside) / 2

	fmt.Println(boundary + inside + 1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
