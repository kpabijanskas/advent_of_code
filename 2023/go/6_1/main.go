package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type race struct {
	time           uint64
	recordDistance uint64
}

func main() {
	times := readFile("./input/time")
	distances := readFile("./input/distance")
	var races []race

	for _, time := range times {
		races = append(races, race{time: time})
	}
	for i, distance := range distances {
		races[i].recordDistance = distance
	}

	m := uint64(1)
	for _, race := range races {
		var waysToBeatRecord uint64
		for i := uint64(1); i <= race.time; i++ {
			distance := calcDistance(i, race.time)
			if distance > race.recordDistance {
				waysToBeatRecord += 1
			}
		}
		m *= waysToBeatRecord
	}

	fmt.Println(m)
}

func calcDistance(timeHeld, totalLength uint64) uint64 {
	return timeHeld * (totalLength - timeHeld)
}

func readFile(path string) []uint64 {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)

	var nums []uint64
	for _, num := range bytes.Fields(data) {
		n, err := strconv.ParseUint(string(num), 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	return nums
}
