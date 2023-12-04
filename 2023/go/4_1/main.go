package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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

	var scoreSum uint64

	for s.Scan() {
		var gameScore uint64
		cardData := bytes.Split(s.Bytes(), []byte(": "))[1]
		winningNumbers := map[uint64]bool{}

		numData := bytes.Split(cardData, []byte(" | "))
		for _, b := range bytes.Split(bytes.TrimSpace(numData[0]), []byte(" ")) {
			if len(b) == 0 {
				continue
			}

			n, err := strconv.ParseUint(string(b), 10, 64)
			if err != nil {
				panic(err)
			}
			winningNumbers[n] = true
		}

		for _, b := range bytes.Split(bytes.TrimSpace(numData[1]), []byte(" ")) {
			if len(b) == 0 {
				continue
			}

			n, err := strconv.ParseUint(string(b), 10, 64)
			if err != nil {
				panic(err)
			}

			if winningNumbers[n] {
				if gameScore == 0 {
					gameScore = 1
				} else {
					gameScore *= 2
				}
			}
		}
		scoreSum += gameScore
	}

	fmt.Println(scoreSum)
}
