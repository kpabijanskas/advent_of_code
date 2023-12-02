package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color uint8

const (
	RED color = iota
	GREEN
	BLUE
)

const (
	DEFAULT_RED   = 12
	DEFAULT_GREEN = 13
	DEFAULT_BLUE  = 13
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	b := bytes.NewReader(data)
	s := bufio.NewScanner(b)
	s.Split(bufio.ScanLines)

	availableCubes := make(map[color]int64, 3)

	var idSum int64
LOOP:
	for s.Scan() {

		gameData := strings.Split(s.Text(), ":")

		gameID, err := strconv.ParseInt(strings.Split(gameData[0], " ")[1], 10, 64)
		if err != nil {
			panic(err)
		}

		for _, revealed := range strings.Split(gameData[1], ";") {
			availableCubes[RED] = DEFAULT_RED
			availableCubes[GREEN] = DEFAULT_GREEN
			availableCubes[BLUE] = DEFAULT_BLUE

			for _, cubes := range strings.Split(revealed, ",") {
				cubes = strings.TrimSpace(cubes)
				var clr color
				switch {
				case strings.Contains(cubes, "red"):
					clr = RED
				case strings.Contains(cubes, "green"):
					clr = GREEN
				case strings.Contains(cubes, "blue"):
					clr = BLUE
				default:
					panic(fmt.Sprintf("no known color in %s", cubes))
				}

				count, err := strconv.ParseInt(strings.Split(cubes, " ")[0], 10, 64)
				if err != nil {
					panic(err)
				}

				availableCubes[clr] -= count
				if availableCubes[clr] < 0 {
					continue LOOP
				}
			}
		}
		idSum += gameID

	}
	fmt.Println(idSum)

}
