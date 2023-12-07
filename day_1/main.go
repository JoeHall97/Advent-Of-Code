package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// partOne("puzzle_input.txt")
	partTwo("puzzle_input.txt")
}

const (
	ZERO_CHAR_BYTE_VALUE = 48
	NINE_CHAR_BYTE_VALUE = 57
)

var numberStrings = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// for each line, find the first and last digit contained within the text
// append these two digits together to get a two digit number
// sum all numbers together to get the answer
func partOne(file_name string) {
	bs := readInputFromFile(file_name)

	first_num, last_num, sum := 0, 0, 0

	for _, s := range bs {
		if s == '\n' {
			sum += first_num + last_num
			first_num, last_num = 0, 0
			continue
		}

		if s > ZERO_CHAR_BYTE_VALUE && s <= NINE_CHAR_BYTE_VALUE {
			last_num = int(s) - ZERO_CHAR_BYTE_VALUE
		}
		if first_num == 0 {
			first_num = last_num * 10
		}
	}
	sum += first_num + last_num
	fmt.Println("SUM: ", sum)
}

// same as part one, but now numbers can be spelled out (e.g. eight, two, etc.)
func partTwo(file_name string) {
	bs := readInputFromFile(file_name)
	s := byteToStringArray(bs)

	sum := 0

	for _, line := range s {
		sum += findNumbers(line)
	}
	fmt.Println("SUM: ", sum)
}

func findNumbers(s string) int {
	fNum, lNum := 0, 0

	// not sure why this doesn't work, but it doesn't
	// for key, value := range numberStrings {
	// 	tempF := strings.Index(s, key)
	// 	tempL := strings.LastIndex(s, key)

	// 	if tempF != -1 && (fi == -1 || tempF < fi) {
	// 		fi = tempF
	// 		fNum = value
	// 	}
	// 	if tempL != -1 && (li != -1 || tempL > li) {
	// 		li = tempL
	// 		lNum = value
	// 	}
	// }

	// check first number
	for i := 0; i < len(s); i++ {
		if s[i] > ZERO_CHAR_BYTE_VALUE && s[i] <= NINE_CHAR_BYTE_VALUE {
			fNum = int(s[i]) - ZERO_CHAR_BYTE_VALUE
			break
		}
		for key, value := range numberStrings {
			if key[0] == s[i] && i+len(key) < len(s) {
				if s[i:i+len(key)] == key {
					fNum = value
					break
				}
			}
		}
		if fNum != 0 {
			break
		}
	}

	// check last number
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > ZERO_CHAR_BYTE_VALUE && s[i] <= NINE_CHAR_BYTE_VALUE {
			lNum = int(s[i]) - ZERO_CHAR_BYTE_VALUE
			break
		}
		for key, value := range numberStrings {
			if key[len(key)-1] == s[i] && i-len(key)+1 > 0 {
				if s[i-len(key)+1:i+1] == key {
					lNum = value
					break
				}
			}
		}
		if lNum != 0 {
			break
		}
	}

	return (fNum * 10) + lNum
}

func byteToStringArray(byteArray []byte) []string {
	lines := make([]string, 0, 0)
	var out bytes.Buffer

	for _, b := range byteArray {
		if b != '\n' {
			out.WriteByte(b)
		} else {
			lines = append(lines, out.String())
			out.Reset()
		}
	}
	lines = append(lines, out.String())

	return lines
}

func readInputFromFile(file_name string) []byte {
	bs, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return bs
}
