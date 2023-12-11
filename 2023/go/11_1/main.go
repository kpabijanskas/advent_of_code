package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const sizeFactor = 2

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	data = bytes.TrimSpace(data)
	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)

	emptyRows := map[int]bool{}
	emptyColumns := map[int]bool{}
	galaxies := [][]int{}

	var row int
	for s.Scan() {
		// prep columns on first loop
		if row == 0 {
			for i := 0; i < len(s.Bytes()); i++ {
				emptyColumns[i] = true
			}
		}

		emptyRow := true
		for column, b := range s.Bytes() {
			if b == '#' {
				emptyRow = false
				emptyColumns[column] = false
				galaxies = append(galaxies, []int{row, column})
			}
		}
		if emptyRow {
			emptyRows[row] = true
		}
		row++
	}

	var spSum int
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[i+1:] {
			var sp int
			row1, row2 := galaxy1[0], galaxy2[0]
			if row2 < row1 {
				row1, row2 = row2, row1
			}

			for k := row1; k < row2; k++ {
				if emptyRows[k] {
					sp += sizeFactor
				} else {
					sp++
				}
			}

			column1, column2 := galaxy1[1], galaxy2[1]
			if column2 < column1 {
				column1, column2 = column2, column1
			}

			for k := column1; k < column2; k++ {
				if emptyColumns[k] {
					sp += sizeFactor
				} else {
					sp++
				}
			}
			spSum += sp
		}
	}

	fmt.Println(spSum)
}
