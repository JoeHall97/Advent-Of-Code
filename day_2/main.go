package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	numRe  = regexp.MustCompile(`\d+`)
	wordRe = regexp.MustCompile(`[a-zA-Z]+`)
)

type ColourValues struct {
	GreenValue int
	BlueValue  int
	RedValue   int
}

func main() {
	bytes := readInputFromFile("./example_input_one.txt")
	// bytes := readInputFromFile("puzzle_input.txt")
	input := byteToStringArray(bytes)
	sum := 0
	pow := 0
	maxValues := ColourValues{13, 14, 12}

	startre := regexp.MustCompile(`[:]`)

	for i, in := range input {
		idx := startre.FindStringIndex(in)
		// remove Game #: from the input before passing
		if isLineValid(in[idx[1]+1:], maxValues) {
			sum += i + 1
		}

		posGame := getPossibleGame(in[idx[1]+1:])
		pow += posGame.BlueValue * posGame.GreenValue * posGame.RedValue
	}

	fmt.Printf("SUM: %d\n", sum)
	fmt.Printf("POW: %d\n", pow)
}

func getPossibleGame(input string) ColourValues {
	n := numRe.FindAllString(input, -1)
	w := wordRe.FindAllString(input, -1)
	res := ColourValues{0, 0, 0}

	for i := range n {
		v, err := strconv.Atoi(n[i])
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch w[i] {
		case "green":
			if v > res.GreenValue {
				res.GreenValue = v
			}
		case "blue":
			if v > res.BlueValue {
				res.BlueValue = v
			}
		case "red":
			if v > res.RedValue {
				res.RedValue = v
			}
		default:
			fmt.Printf("ERROR: invalid input line, %s\n", input)
		}

	}

	return res
}

func isLineValid(input string, maxValues ColourValues) bool {
	n := numRe.FindAllString(input, -1)
	w := wordRe.FindAllString(input, -1)

	for i := range n {
		v, err := strconv.Atoi(n[i])
		if err != nil {
			fmt.Println(err)
			return false
		}

		switch w[i] {
		case "green":
			if v > maxValues.GreenValue {
				return false
			}
		case "blue":
			if v > maxValues.BlueValue {
				return false
			}
		case "red":
			if v > maxValues.RedValue {
				return false
			}
		default:
			fmt.Printf("ERROR: invalid input line, %s\n", input)
			return false
		}
	}

	return true
}

func byteToStringArray(byteArray []byte) []string {
	lines := make([]string, 0)
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
