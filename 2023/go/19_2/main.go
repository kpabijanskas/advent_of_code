package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type cond struct {
	part   string
	cond   string
	val    int64
	target string
}

type rating struct {
	x_start, x_end int64
	m_start, m_end int64
	a_start, a_end int64
	s_start, s_end int64
	next           string
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

	workflows := make(map[string][]cond)

	for s.Scan() {
		if len(s.Bytes()) == 0 {
			break
		}

		fields := bytes.Split(s.Bytes(), []byte("{"))
		fields[1] = bytes.TrimRight(fields[1], "}")

		conditions := bytes.Split(fields[1], []byte(","))

		workflowName := string(fields[0])
		workflows[workflowName] = make([]cond, 0)

		for _, condition := range conditions {
			f2 := bytes.Split(condition, []byte(":"))
			if len(f2) == 1 {
				workflows[workflowName] = append(workflows[workflowName], cond{
					target: string(f2[0]),
				})
				continue
			}

			var f3 [][]byte
			var c string
			if bytes.Index(f2[0], []byte("<")) > 0 {
				f3 = bytes.Split(f2[0], []byte("<"))
				c = "<"
			} else {
				f3 = bytes.Split(f2[0], []byte(">"))
				c = ">"
			}

			n, err := strconv.ParseInt(string(f3[1]), 10, 64)
			if err != nil {
				panic(err)
			}

			workflows[workflowName] = append(workflows[workflowName], cond{
				part:   string(f3[0]),
				cond:   c,
				val:    n,
				target: string(f2[1]),
			})
		}
	}

	startPart := rating{
		1, 4000, 1, 4000, 1, 4000, 1, 4000, "in",
	}

	queue := []rating{startPart}
	ansQueue := []rating{}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.next == "A" {
			ansQueue = append(ansQueue, p)
		} else if p.next == "R" {
			continue
		} else {
			for _, w := range workflows[p.next] {
				if w.cond == "" {
					p.next = w.target
					queue = append(queue, p)
				} else if w.cond == "<" {
					switch w.part {
					case "x":
						p2 := p
						p2.x_end = w.val - 1
						p2.next = w.target
						queue = append(queue, p2)
						p.x_start = w.val
					case "m":
						p2 := p
						p2.m_end = w.val - 1
						p2.next = w.target
						queue = append(queue, p2)
						p.m_start = w.val
					case "a":
						p2 := p
						p2.a_end = w.val - 1
						p2.next = w.target
						queue = append(queue, p2)
						p.a_start = w.val
					case "s":
						p2 := p
						p2.s_end = w.val - 1
						p2.next = w.target
						queue = append(queue, p2)
						p.s_start = w.val
					}
				} else { // >
					switch w.part {
					case "x":
						p2 := p
						p2.x_start = w.val + 1
						p2.next = w.target
						queue = append(queue, p2)
						p.x_end = w.val
					case "m":
						p2 := p
						p2.m_start = w.val + 1
						p2.next = w.target
						queue = append(queue, p2)
						p.m_end = w.val
					case "a":
						p2 := p
						p2.a_start = w.val + 1
						p2.next = w.target
						queue = append(queue, p2)
						p.a_end = w.val
					case "s":
						p2 := p
						p2.s_start = w.val + 1
						p2.next = w.target
						queue = append(queue, p2)
						p.s_end = w.val
					}
				}

			}

		}

	}

	var ans int64

	for _, p := range ansQueue {
		ans += (p.x_end - p.x_start + 1) * (p.m_end - p.m_start + 1) * (p.a_end - p.a_start + 1) * (p.s_end - p.s_start + 1)
	}

	fmt.Println(ans)
}
