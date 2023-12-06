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
	time := readFile("./input/time")
	recordDistance := readFile("./input/distance")
	var waysToBeatRecord uint64
	for i := uint64(1); i <= time; i++ {
		distance := calcDistance(i, time)
		if distance > recordDistance {
			waysToBeatRecord += 1
		}
	}

	fmt.Println(waysToBeatRecord)
}

func calcDistance(timeHeld, totalLength uint64) uint64 {
	return timeHeld * (totalLength - timeHeld)
}

func readFile(path string) uint64 {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	n, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}
