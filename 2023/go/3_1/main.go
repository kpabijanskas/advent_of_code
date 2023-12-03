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

	var partsSum uint64

	for lineNo, line := range dataLines {
		var inNumber bool
		var num []byte
		var isPartNumber bool

		for byteNo, b := range line {
			if unicode.IsDigit(rune(b)) {
				num = append(num, b)
				inNumber = true

				if !isPartNumber {
					if (byteNo > 0 && isPartIndicator(dataLines[lineNo][byteNo-1])) ||
						(byteNo < len(line)-1 && isPartIndicator(dataLines[lineNo][byteNo+1])) ||
						(lineNo > 0 && isPartIndicator(dataLines[lineNo-1][byteNo])) ||
						(lineNo < len(dataLines)-1 && isPartIndicator(dataLines[lineNo+1][byteNo])) ||
						(byteNo > 0 && lineNo > 0 && isPartIndicator(dataLines[lineNo-1][byteNo-1])) ||
						(byteNo < len(line)-1 && lineNo < len(dataLines)-1 && isPartIndicator(dataLines[lineNo+1][byteNo+1])) ||
						(byteNo > 0 && lineNo < len(dataLines)-1 && isPartIndicator(dataLines[lineNo+1][byteNo-1])) ||
						(byteNo < len(line)-1 && lineNo > 0 && isPartIndicator(dataLines[lineNo-1][byteNo+1])) {
						isPartNumber = true
					}
				}
			}

			if (!unicode.IsDigit(rune(b)) || (byteNo == len(line)-1 && inNumber)) && inNumber {
				if isPartNumber {
					n, err := strconv.ParseUint(string(num), 10, 64)
					if err != nil {
						panic(err)
					}
					partsSum += n
				}

				//reset
				inNumber = false
				num = nil
				isPartNumber = false

			}
		}
	}
	fmt.Println(partsSum)
}

func isPartIndicator(b byte) bool {
	if !unicode.IsDigit(rune(b)) && b != '.' {
		return true
	}
	return false
}
