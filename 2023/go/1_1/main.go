package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	b := bytes.NewReader(data)
	s := bufio.NewScanner(b)
	s.Split(bufio.ScanLines)

	var sum int64

	for s.Scan() {
		var digit string
		for _, r := range s.Text() {
			if unicode.IsDigit(r) {
				digit += string(r)
				break
			}
		}

		for i := len(s.Text()) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(s.Text()[i])) {
				digit += string(s.Text()[i])
				break
			}
		}

		n, err := strconv.ParseInt(digit, 10, 64)
		if err != nil {
			panic(err)
		}

		sum += n
	}

	fmt.Println(sum)
}
