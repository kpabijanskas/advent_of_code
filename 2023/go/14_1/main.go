package main

import (
	"bytes"
	"fmt"
	"os"

	"slices"
)

const (
	roundRock byte = 'O'
	cubeRock  byte = '#'
)

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)

	platform := bytes.Split(data, []byte("\n"))

	tiltNorth(platform)

	var totalLoad int
	for i := range platform {
		totalLoad += bytes.Count(platform[i], []byte{roundRock}) * (len(platform) - i)
	}

	fmt.Println(totalLoad)
}

func tiltNorth(platform [][]byte) {
	for column := 0; column < len(platform[0]); column++ {
		indexes := []int{0}
		for row := range platform {
			if platform[row][column] == '#' {
				indexes = append(indexes, row)
			}
		}
		indexes = append(indexes, len(platform)-1)

		for i := range indexes[:len(indexes)-1] {
			topRow, bottomRow := indexes[i], indexes[i+1]
			for topRow < bottomRow {
				switch {
				case slices.Contains([]byte{roundRock, cubeRock}, platform[topRow][column]):
					topRow++
				case platform[bottomRow][column] == cubeRock:
					bottomRow--
				case platform[bottomRow][column] == roundRock:
					platform[bottomRow][column], platform[topRow][column] = platform[topRow][column], platform[bottomRow][column]
				default:
					bottomRow--
				}
			}
		}
	}
}
