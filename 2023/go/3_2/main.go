package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace((data))

	dataLines := bytes.Split(data, []byte("\n"))

	var ratioSum uint64

	for lineNo, line := range dataLines {
		for byteNo, b := range line {
			if b == byte('*') {
				anums := adjacentNumbers(dataLines, lineNo, byteNo)
				if len(anums) == 2 {
					ratioSum += anums[0] * anums[1]
				}
			}
		}
	}

	fmt.Println(ratioSum)
}

func adjacentNumbers(data [][]byte, x, y int) []uint64 {
	var nums []uint64

	if x > 0 {
		if unicode.IsDigit(rune(data[x-1][y])) {
			nums = append(nums, getFullNum(data, x-1, y))
		} else {
			if y > 0 && unicode.IsDigit(rune(data[x-1][y-1])) {
				nums = append(nums, getFullNum(data, x-1, y-1))
			}
			if y < len(data[x])-1 && unicode.IsDigit(rune(data[x-1][y+1])) {
				nums = append(nums, getFullNum(data, x-1, y+1))
			}
		}
	}

	if x < len(data)-1 {
		if unicode.IsDigit(rune(data[x+1][y])) {
			nums = append(nums, getFullNum(data, x+1, y))
		} else {
			if y > 0 && unicode.IsDigit(rune(data[x+1][y-1])) {
				nums = append(nums, getFullNum(data, x+1, y-1))
			}
			if y < len(data[x])-1 && unicode.IsDigit(rune(data[x+1][y+1])) {
				nums = append(nums, getFullNum(data, x+1, y+1))
			}
		}
	}

	if y > 0 && unicode.IsDigit(rune(data[x][y-1])) {
		nums = append(nums, getFullNum(data, x, y-1))
	}

	if y < len(data[x])-1 && unicode.IsDigit(rune(data[x][y+1])) {
		nums = append(nums, getFullNum(data, x, y+1))
	}

	return nums
}

func getFullNum(data [][]byte, x, y int) uint64 {
	start, end := y, y+1

	for start > 0 && unicode.IsDigit(rune(data[x][start-1])) {
		start--
	}

	for end < len(data[x]) && unicode.IsDigit(rune(data[x][end])) {
		end++
	}

	n, err := strconv.ParseUint(string(data[x][start:end]), 10, 64)
	if err != nil {
		panic(err)
	}

	return n
}
