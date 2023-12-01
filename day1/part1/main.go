package main

import (
	"os"
	"strconv"
	"strings"
)

func getDigits(s string, c chan int) {
	digits := [2]int{}

	for _, char := range strings.Split(s, "") {
		current, err := strconv.Atoi(char)
		if err != nil {
			continue
		}
		if digits[0] == 0 {
			digits[0] = current
			digits[1] = current
		} else {
			digits[1] = current
		}
	}

	c <- 10*digits[0] + digits[1]
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
