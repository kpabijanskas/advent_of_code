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

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	b := bytes.NewReader(data)
	s := bufio.NewScanner(b)
	s.Split(bufio.ScanLines)

	var pwrSum int64
	for s.Scan() {
		availableCubes := make(map[color]int64, 3)

		for _, revealed := range strings.Split(strings.Split(s.Text(), ":")[1], ";") {
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

				availableCubes[clr] = max(availableCubes[clr], count)
			}

		}
		pwrSum += availableCubes[RED] * availableCubes[GREEN] * availableCubes[BLUE]

	}
	fmt.Println(pwrSum)

}
