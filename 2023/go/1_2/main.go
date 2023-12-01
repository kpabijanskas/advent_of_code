package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		for i := 0; i < len(s.Text()); i++ {
			d, ok := getDigitAsString(s.Text()[i:])
			if ok {
				digit += d
				break
			}
		}

		for i := len(s.Text()) - 1; i >= 0; i-- {
			d, ok := getDigitAsString(s.Text()[i:])
			if ok {
				digit += d
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

func getDigitAsString(s string) (string, bool) {
	switch {
	case unicode.IsDigit(rune(s[0])):
		return string(s[0]), true
	case strings.Index(s, "one") == 0:
		return "1", true
	case strings.Index(s, "two") == 0:
		return "2", true
	case strings.Index(s, "three") == 0:
		return "3", true
	case strings.Index(s, "four") == 0:
		return "4", true
	case strings.Index(s, "five") == 0:
		return "5", true
	case strings.Index(s, "six") == 0:
		return "6", true
	case strings.Index(s, "seven") == 0:
		return "7", true
	case strings.Index(s, "eight") == 0:
		return "8", true
	case strings.Index(s, "nine") == 0:
		return "9", true
	}

	return "", false
}
