package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	seeds := readSeeds()
	seedToSoilMap := readMap("seed-to-soil")
	soilToFertilizerMap := readMap("soil-to-fertilizer")
	fertilizerToWaterMap := readMap("fertilizer-to-water")
	waterToLightMap := readMap("water-to-light")
	lightToTemperatureMap := readMap("light-to-temperature")
	temperatureToHumidityMap := readMap("temperature-to-humidity")
	humidityToLocationMap := readMap("humidity-to-location")

	seeds = applyMap(seeds, seedToSoilMap)
	seeds = applyMap(seeds, soilToFertilizerMap)
	seeds = applyMap(seeds, fertilizerToWaterMap)
	seeds = applyMap(seeds, waterToLightMap)
	seeds = applyMap(seeds, lightToTemperatureMap)
	seeds = applyMap(seeds, temperatureToHumidityMap)
	seeds = applyMap(seeds, humidityToLocationMap)

	fmt.Println(slices.Min(seeds))
}

func applyMap(input []uint64, sm []seedMap) []uint64 {
	var output []uint64

LOOP:
	for _, seed := range input {
		for _, m := range sm {
			if seed >= m.srcStart && seed < m.srcStart+m.size {
				offset := seed - m.srcStart
				dst := m.dstStart + offset
				output = append(output, dst)
				continue LOOP
			}
		}
		output = append(output, seed)
	}

	return output
}

func readSeeds() []uint64 {
	data, err := os.ReadFile("./input/seeds")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)

	var seeds []uint64
	for _, nb := range bytes.Split(data, []byte(" ")) {
		if len(nb) == 0 {
			continue
		}
		seeds = append(seeds, parseUint(nb))
	}

	return seeds
}

type seedMap struct {
	srcStart uint64
	dstStart uint64
	size     uint64
}

func readMap(mapName string) []seedMap {
	data, err := os.ReadFile(fmt.Sprintf("./input/%s", mapName))
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)

	br := bytes.NewReader(data)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanLines)

	var m []seedMap
	for s.Scan() {
		fields := bytes.Split(s.Bytes(), []byte(" "))

		if len(fields) != 3 {
			panic(fmt.Errorf("wrong amount of fields, expected 3, got %+v", fields))
		}

		dstRangeStart := parseUint(fields[0])
		srcRangeStart := parseUint(fields[1])
		rangeLen := parseUint(fields[2])

		m = append(m, seedMap{
			srcStart: srcRangeStart,
			dstStart: dstRangeStart,
			size:     rangeLen,
		})
	}

	return m
}

func parseUint(b []byte) uint64 {
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n

}
