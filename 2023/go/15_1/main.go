package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	steps := bytes.Split(bytes.Split(data, []byte("\n"))[0], []byte(","))

	var result uint64
	for _, step := range steps {
		var s uint64

		for _, b := range step {
			s += uint64(b)
			s *= 17
			s %= 256
		}

		result += s
	}

	fmt.Println(result)
}
