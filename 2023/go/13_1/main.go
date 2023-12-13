package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const DIFFERENCE = 0

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	data = bytes.TrimSpace(data)

	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)

	var rows int
	var columns int
	var currentPattern []string

	for s.Scan() {
		if len(s.Text()) > 0 {
			currentPattern = append(currentPattern, s.Text())
			continue
		}

		firstRow := getRowCount(currentPattern)
		if firstRow >= 0 {
			rows += firstRow
		}

		firstColumn := getColumnCount(currentPattern)
		if firstColumn >= 0 {
			columns += firstColumn
		}

		currentPattern = []string{}
	}

	firstRow := getRowCount(currentPattern)
	if firstRow >= 0 {
		rows += firstRow
	}
	firstColumn := getColumnCount(currentPattern)
	if firstColumn > 0 {
		columns += firstColumn
	}

	fmt.Println(columns + 100*rows)
}

func getRowCount(pattern []string) int {
	for i := range pattern[:len(pattern)-1] {
		if rowDifferenceCount(pattern[i], pattern[i+1]) <= DIFFERENCE && verifyRowPattern(pattern, i) {
			return i + 1
		}
	}

	return -1
}

func verifyRowPattern(pattern []string, firstRow int) bool {
	var totalDifferences int
	for i, j := firstRow, firstRow+1; i >= 0 && j < len(pattern); i, j = i-1, j+1 {
		totalDifferences += rowDifferenceCount(pattern[i], pattern[j])
	}
	return totalDifferences == DIFFERENCE
}

func rowDifferenceCount(row1, row2 string) int {
	var diff int
	for i := range row1 {
		if row1[i] != row2[i] {
			diff += 1
		}
	}
	return diff
}

func getColumnCount(pattern []string) int {
	for col := range pattern[0][:len(pattern[0])-1] {
		if columnDifferenceCount(col, col+1, pattern) <= DIFFERENCE && verifyColumnPattern(pattern, col) {
			return col + 1
		}
	}
	return -1
}

func verifyColumnPattern(pattern []string, firstColumn int) bool {
	var totalDifferences int
	for i, j := firstColumn, firstColumn+1; i >= 0 && j < len(pattern[0]); i, j = i-1, j+1 {
		totalDifferences += columnDifferenceCount(i, j, pattern)
	}
	return totalDifferences == DIFFERENCE
}

func columnDifferenceCount(col1, col2 int, pattern []string) int {
	var diff int
	for _, row := range pattern {
		if row[col1] != row[col2] {
			diff += 1
		}
	}

	return diff
}
