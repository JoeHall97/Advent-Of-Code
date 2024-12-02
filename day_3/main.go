package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	symbol_matchs = make([][]int, 0)
)

func main() {
	//puzzleOne()
	puzzleTwo()
}

func puzzleTwo() {
	//input := readInputFromFile("example_one_input.txt")

}

func puzzleOne() {
	input := readInputFromFile("puzzle_input.txt")
	symbol_re := regexp.MustCompile(`[^\d. \n]`)
	digit_re := regexp.MustCompile(`(\d+)`)
	digit_matchs := make([][][]int, 0)

	for _, in := range input {
		// find all of the symbols in the current row
		match := symbol_re.FindAllStringIndex(in, -1)
		line_symbol_matchs := make([]int, 0)
		for _, m := range match {
			line_symbol_matchs = append(line_symbol_matchs, m[0])
		}
		symbol_matchs = append(symbol_matchs, line_symbol_matchs)

		// find all of the digits in the current row
		match = digit_re.FindAllStringIndex(in, -1)
		line_digit_matchs := make([][]int, 0)
		line_digit_matchs = append(line_digit_matchs, match...)
		digit_matchs = append(digit_matchs, line_digit_matchs)
	}

	sum := 0
	for idx, dm := range digit_matchs {
		for _, m := range dm {
			if isValidDigit(idx, m[0], m[1]) {
				num, err := strconv.Atoi(input[idx][m[0]:m[1]])
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				sum += num
			}
		}
	}

	fmt.Printf("SUM: %d\n", sum)
}

func isValidDigit(row_index int, start_index int, end_index int) bool {
	// check above the number
	if row_index-1 >= 0 && len(symbol_matchs[row_index-1]) > 0 {
		for _, sm := range symbol_matchs[row_index-1] {
			if sm >= start_index-1 && sm <= end_index {
				return true
			}
		}
	}

	// check below the number
	if row_index+1 < len(symbol_matchs) && len(symbol_matchs[row_index+1]) > 0 {
		for _, sm := range symbol_matchs[row_index+1] {
			if sm >= start_index-1 && sm <= end_index {
				return true
			}
		}
	}

	// check beside the number
	if len(symbol_matchs[row_index]) > 0 {
		for _, sm := range symbol_matchs[row_index] {
			if sm == start_index-1 || sm == end_index {
				return true
			}
		}
	}

	return false
}

func readInputFromFile(file_name string) []string {
	lines := make([]string, 0)
	l := ""
	bs, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	for _, b := range bs {
		if b == '\n' {
			lines = append(lines, l)
			l = ""
		} else {
			l += string(b)
		}
	}

	lines = append(lines, l)

	return lines
}
