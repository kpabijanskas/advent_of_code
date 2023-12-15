package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type lense struct {
	label    string
	focalLen uint64
}

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	steps := bytes.Split(bytes.Split(data, []byte("\n"))[0], []byte(","))

	var boxes [256][]lense
	for _, step := range steps {
		switch {
		case bytes.Index(step, []byte("-")) > 0:
			box := hash(step[:len(step)-1])
			label := string(step[:len(step)-1])
			boxes[box] = slices.DeleteFunc(boxes[box], func(l lense) bool {
				return l.label == label
			})

		case bytes.Index(step, []byte("=")) > 0:
			fields := bytes.FieldsFunc(step, func(r rune) bool {
				return r == '='
			})
			box := hash(fields[0])
			label := string(fields[0])
			focalLen, err := strconv.ParseUint(string(fields[1]), 10, 64)
			if err != nil {
				panic(err)
			}

			idx := slices.IndexFunc(boxes[box], func(l lense) bool {
				return l.label == label
			})

			if idx == -1 {
				boxes[box] = append(boxes[box], lense{
					label:    label,
					focalLen: focalLen,
				})
				continue
			}
			boxes[box][idx].focalLen = focalLen

		default:
			panic(fmt.Sprintf("can't find action in %s", string(step)))
		}
	}

	var result uint64
	for i, box := range boxes {
		for j, l := range box {
			score := uint64(i+1) * uint64(j+1) * l.focalLen
			result += score
		}
	}

	fmt.Println(result)
}

func hash(step []byte) int {
	var s int

	for _, b := range step {
		s += int(b)
		s *= 17
		s %= 256
	}

	return s
}
