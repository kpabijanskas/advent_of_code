package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	cardID             int
	winningNumberCount uint
}

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)

	allCardData := []card{}
	queue := []card{}

	for i := 0; s.Scan(); i++ {
		cardData := bytes.Split(s.Bytes(), []byte(": "))
		numData := bytes.Split(cardData[1], []byte(" | "))

		var winningNumberCount uint
		winningNumbers := map[uint64]bool{}
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
				winningNumberCount++
			}
		}

		allCardData = append(allCardData, card{cardID: i, winningNumberCount: winningNumberCount})
		queue = append(queue, card{cardID: i, winningNumberCount: winningNumberCount})
	}

	var totalCards uint
	for len(queue) > 0 {
		card := queue[0]
		queue = queue[1:]
		totalCards++

		if card.winningNumberCount > 0 {
			for i := card.cardID + 1; i < card.cardID+1+int(card.winningNumberCount); i++ {
				queue = append(queue, allCardData[i])
			}
		}
	}

	fmt.Println(totalCards)
}
