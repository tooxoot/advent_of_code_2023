package main

import (
	"os"
	"regexp"
	"strings"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func convertToInt(s string) int {
	return digitMap[s]
}

var digitRegex = regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)

func getDigits(s string, c chan int) {
	current := s
	idxs := digitRegex.FindStringIndex(s)
	first := convertToInt(s[idxs[0]:idxs[1]])
	last := 0

	for idxs != nil {

		last = convertToInt(current[idxs[0]:idxs[1]])
		startIdx := idxs[1]
		if idxs[1]-idxs[0] > 1 {
			startIdx -= 1
		}
		current = current[startIdx:]
		idxs = digitRegex.FindStringIndex(current)
	}

	c <- 10*first + last
}

func main() {
	dat, err := os.ReadFile("../input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	results := make(chan int)
	for _, line := range lines {
		go getDigits(line, results)
	}

	lineCount := len(lines)
	i := 0
	sum := 0
	for i < lineCount {
		sum += <-results
		i += 1
	}

	println(sum)
}
