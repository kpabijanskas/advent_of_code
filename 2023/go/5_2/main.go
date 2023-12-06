package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].start < seeds[j].start
	})

	fmt.Println(seeds[0].start)
}

func applyMap(input []seedRange, sm []seedMap) []seedRange {
	var output []seedRange

LOOP:
	for len(input) > 0 {
		seed := input[0]
		input = input[1:]

		for _, m := range sm {
			overlaping, ok := overlap(seed, m)

			if ok {
				nonoverlaping := nonOverlap(seed, m)
				input = append(input, nonoverlaping...)

				offset := overlaping.start - m.start
				overlaping.start = m.dst + offset
				output = append(output, overlaping)

				continue LOOP
			}
		}
		output = append(output, seed)
	}

	return output
}

func overlap(s seedRange, m seedMap) (seedRange, bool) {
	a := []uint64{s.start, s.start + s.size}
	b := []uint64{m.start, m.start + m.size}
	start := max(a[0], b[0])
	end := min(a[1], b[1])
	if end <= start {
		return seedRange{}, false
	}

	r := seedRange{start, end - start}
	return r, true
}

func nonOverlap(s seedRange, m seedMap) []seedRange {

	a := []uint64{s.start, s.start + s.size}
	b := []uint64{m.start, m.start + m.size}

	ranges := make([]seedRange, 0, 2)

	if a[0] < b[0] {
		start := a[0]
		end := min(a[1], b[0])

		ranges = append(ranges, seedRange{start, end - start})
	}

	if a[1] > b[1] {
		start := max(b[1], a[0])
		end := a[1]
		ranges = append(ranges, seedRange{start, end - start})
	}

	return ranges
}

func readSeeds() []seedRange {
	data, err := os.ReadFile("./input/seeds")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)

	fields := bytes.Split(data, []byte(" "))

	var seeds []seedRange
	for i := 0; i < len(fields); i += 2 {
		start := parseUint(fields[i])
		size := parseUint(fields[i+1])

		seeds = append(seeds, seedRange{
			start: start,
			size:  size,
		})
	}

	return seeds
}

type seedRange struct {
	start uint64
	size  uint64
}

type seedMap struct {
	seedRange
	dst uint64
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

		dst := parseUint(fields[0])
		src := parseUint(fields[1])
		size := parseUint(fields[2])

		m = append(m, seedMap{
			seedRange: seedRange{
				start: src,
				size:  size,
			},
			dst: dst,
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
