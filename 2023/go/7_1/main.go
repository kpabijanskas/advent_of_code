package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type cardType uint

const (
	highCardCardType cardType = iota
	onePairCardType
	twoPairCardType
	threeOfAKIndCardType
	fullHouseCardType
	fourOfAKindCardType
	fiveOfAKindCardType
)

type cardLabels uint8

const (
	cardLabel2 cardLabels = iota
	cardLabel3
	cardLabel4
	cardLabel5
	cardLabel6
	cardLabel7
	cardLabel8
	cardLabel9
	cardLabelT
	cardLabelJ
	cardLabelQ
	cardLabelK
	cardLabelA
)

type card struct {
	score uint64
	hand  []cardLabels
	ct    cardType
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

	var cards []card
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) != 2 {
			panic(err)
		}

		score, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			panic(err)
		}

		c := card{
			score: score,
			hand:  genHand(fields[0]),
			ct:    genCardType(fields[0]),
		}

		cards = append(cards, c)
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].ct < cards[j].ct {
			return true
		}
		if cards[i].ct == cards[j].ct {
			for k := range cards[i].hand {
				if cards[i].hand[k] < cards[j].hand[k] {
					return true
				}
				if cards[i].hand[k] > cards[j].hand[k] {
					return false
				}
			}

		}
		return false
	})

	var score uint64

	for i, card := range cards {
		score += (uint64(i) + 1) * card.score
	}

	fmt.Println(score)
}

func genCardType(hand string) cardType {
	cards := make(map[rune]int)
	for _, card := range hand {
		cards[card] += 1
	}

	frequencies := make([]int, 0, 5)
	for _, cardFreq := range cards {
		frequencies = append(frequencies, cardFreq)
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i] > frequencies[j]
	})

	switch frequencies[0] {
	case 5:
		return fiveOfAKindCardType
	case 4:
		return fourOfAKindCardType
	case 3:
		if frequencies[1] == 2 {
			return fullHouseCardType
		}
		return threeOfAKIndCardType
	case 2:
		if frequencies[1] == 2 {
			return twoPairCardType
		}
		return onePairCardType
	case 1:
		return highCardCardType
	default:
		panic("unknonw card type")
	}

}

func genHand(hand string) []cardLabels {
	labels := make([]cardLabels, 0, 5)
	for _, card := range hand {
		switch card {
		case '2':
			labels = append(labels, cardLabel2)
		case '3':
			labels = append(labels, cardLabel3)
		case '4':
			labels = append(labels, cardLabel4)
		case '5':
			labels = append(labels, cardLabel5)
		case '6':
			labels = append(labels, cardLabel6)
		case '7':
			labels = append(labels, cardLabel7)
		case '8':
			labels = append(labels, cardLabel8)
		case '9':
			labels = append(labels, cardLabel9)
		case 'T':
			labels = append(labels, cardLabelT)
		case 'J':
			labels = append(labels, cardLabelJ)
		case 'Q':
			labels = append(labels, cardLabelQ)
		case 'K':
			labels = append(labels, cardLabelK)
		case 'A':
			labels = append(labels, cardLabelA)
		default:
			panic("unknown label")
		}
	}
	return labels
}
