package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var cases [][]int64
	for s.Scan() {
		fields := strings.Fields(s.Text())

		nums := make([]int64, 0, len(fields))
		for _, field := range fields {
			num, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		cases = append(cases, nums)
	}

	var predSum int64
	for _, c := range cases {
		predSum += getPrediction(c)
	}

	fmt.Println(predSum)
}

func getPrediction(c []int64) int64 {
	allZeroes := true
	next := make([]int64, 0, len(c)-1)
	for i := 0; i < len(c)-1; i++ {
		if c[i] == 0 && c[i+1] == 0 {
			next = append(next, 0)
			continue
		}
		allZeroes = false

		next = append(next, c[i+1]-c[i])
	}

	if allZeroes {
		return 0
	}

	predicted := getPrediction(next)
	return c[0] - predicted

}
