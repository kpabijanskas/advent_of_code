package main

import (
	"bytes"
	"fmt"
	"os"

	"slices"
)

const (
	roundRock   byte = 'O'
	cubeRock    byte = '#'
	totalCycles      = 1000000000
)

var (
	memo = map[string]int{}
)

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)

	platform := bytes.Split(data, []byte("\n"))

	var cycle int
	for {
		memo[string(data)] = cycle
		cycle++

		tiltNorth(platform)
		tiltWest(platform)
		tiltSouth(platform)
		tiltEast(platform)

		if _, ok := memo[string(data)]; ok {
			if (totalCycles-cycle)%(cycle-memo[string(data)]) == 0 {
				var totalLoad int
				for i := range platform {
					totalLoad += bytes.Count(platform[i], []byte{roundRock}) * (len(platform) - i)

				}
				fmt.Println(totalLoad)
				os.Exit(0)
			}
		}
		memo[string(data)] = cycle
	}

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

func tiltSouth(platform [][]byte) {
	for column := 0; column < len(platform[0]); column++ {
		indexes := []int{len(platform) - 1}
		for row := len(platform) - 1; row >= 0; row-- {
			if platform[row][column] == cubeRock {
				indexes = append(indexes, row)
			}
		}
		indexes = append(indexes, 0)

		for i := range indexes[:len(indexes)-1] {
			bottomRow, topRow := indexes[i], indexes[i+1]
			for topRow < bottomRow {
				switch {
				case slices.Contains([]byte{roundRock, cubeRock}, platform[bottomRow][column]):
					bottomRow--
				case platform[topRow][column] == cubeRock:
					topRow++
				case platform[topRow][column] == roundRock:
					platform[bottomRow][column], platform[topRow][column] = platform[topRow][column], platform[bottomRow][column]
				default:
					topRow++
				}
			}

		}
	}
}

func tiltWest(platform [][]byte) {
	for row := range platform {
		indexes := []int{0}
		for column := range platform[row] {
			if platform[row][column] == cubeRock {
				indexes = append(indexes, column)
			}
		}
		indexes = append(indexes, len(platform[row])-1)

		for i := range indexes[:len(indexes)-1] {
			firstColumn, lastColumn := indexes[i], indexes[i+1]
			for firstColumn < lastColumn {
				switch {
				case slices.Contains([]byte{roundRock, cubeRock}, platform[row][firstColumn]):
					firstColumn++
				case platform[row][lastColumn] == cubeRock:
					lastColumn--
				case platform[row][lastColumn] == roundRock:
					platform[row][firstColumn], platform[row][lastColumn] = platform[row][lastColumn], platform[row][firstColumn]
				default:
					lastColumn--
				}
			}
		}

	}
}
func tiltEast(platform [][]byte) {
	for row := range platform {
		indexes := []int{len(platform[row]) - 1}
		for column := len(platform[row]) - 1; column >= 0; column-- {
			if platform[row][column] == cubeRock {
				indexes = append(indexes, column)
			}
		}
		indexes = append(indexes, 0)

		for i := range indexes[:len(indexes)-1] {
			lastColumn, firstColumn := indexes[i], indexes[i+1]
			for firstColumn < lastColumn {
				switch {
				case slices.Contains([]byte{roundRock, cubeRock}, platform[row][lastColumn]):
					lastColumn--
				case platform[row][firstColumn] == cubeRock:
					firstColumn++
				case platform[row][firstColumn] == roundRock:
					platform[row][firstColumn], platform[row][lastColumn] = platform[row][lastColumn], platform[row][firstColumn]
				default:
					firstColumn++
				}
			}
		}
	}
}
