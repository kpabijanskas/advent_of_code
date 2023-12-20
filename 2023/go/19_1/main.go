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

	parts := make([]map[string]int64, 0)
	for s.Scan() {
		d := bytes.TrimRight(bytes.TrimLeft(s.Bytes(), "{"), "}")

		fields := bytes.Split(d, []byte(","))

		part := make(map[string]int64)
		for _, field := range fields {
			f := bytes.Split(field, []byte("="))

			n, err := strconv.ParseInt(string(f[1]), 10, 64)
			if err != nil {
				panic(err)
			}

			part[string(f[0])] = n
		}

		parts = append(parts, part)
	}

	var ans int64

	for _, p := range parts {
		next := "in"

	LOOP:
		for next != "R" && next != "A" {
			for _, w := range workflows[next] {
				if w.cond == "" {
					next = w.target
				}
				if w.cond == "<" {
					if p[w.part] < w.val {
						next = w.target
						continue LOOP
					}
				} else {
					if p[w.part] > w.val {
						next = w.target
						continue LOOP
					}
				}
			}
		}

		if next == "A" {
			for _, c := range p {
				ans += c
			}
		}
	}

	fmt.Println(ans)
}
