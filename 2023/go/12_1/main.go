package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	springDamaged = '#'
	springOK      = '.'
	springUnknown = '?'
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

	var totalCount uint
	for s.Scan() {
		fields := strings.Fields(s.Text())
		totalCount += findCounts(fields)
	}
	fmt.Println(totalCount)

}

func findCounts(maps []string) uint {
	var arrangements []uint64
	for _, no := range strings.Split(maps[1], ",") {
		n, err := strconv.ParseUint(no, 10, 64)
		if err != nil {
			panic(err)
		}
		arrangements = append(arrangements, n)
	}

	return findRecursive(maps[0], arrangements)
}

func skip(springMap string) string {
	for len(springMap) > 0 {
		if springMap[0] == springDamaged {
			return springMap
		}
		springMap = springMap[1:]
	}

	return springMap
}

func findRecursive(springMap string, remainingCounts []uint64) uint {
	if len(remainingCounts) == 0 {
		springMap = skip(springMap)
		if len(springMap) > 0 {
			return 0
		}
	}
	if len(springMap) == 0 && len(remainingCounts) == 0 {
		return 1
	}

	if len(springMap) == 0 {
		return 0
	}

	switch springMap[0] {
	case springDamaged:
		springMap, ok := apply(springMap, remainingCounts[0])
		if ok {
			return findRecursive(springMap, remainingCounts[1:])
		}

		return 0
	case springUnknown:
		count := findRecursive(springMap[1:], remainingCounts)

		springMap, ok := apply(springMap, remainingCounts[0])
		if ok {
			count += findRecursive(springMap, remainingCounts[1:])
		}

		return count
	case springOK:
		count := findRecursive(springMap[1:], remainingCounts)
		return count
	}
	return 0
}

func apply(springMap string, count uint64) (string, bool) {
	for i := uint64(0); i < count; i++ {
		if len(springMap) == 0 {
			return springMap, false
		}

		if springMap[0] == springOK {
			return springMap, false
		}

		springMap = springMap[1:]
	}

	if len(springMap) == 0 {
		return springMap, true
	}

	if springMap[0] == springDamaged {
		return springMap, false
	}

	return springMap[1:], true
}
