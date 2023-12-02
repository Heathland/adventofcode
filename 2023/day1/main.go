package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var digits = map[string]int{
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

// We read the lines from beginning to end and check if it's a digit by either int or the word of a digit
func getFirstDigit(s string, translate bool) (d int) {
	var err error
	for i := 0; i < len(s); i++ {
		d, err = strconv.Atoi(fmt.Sprintf("%c", s[i]))
		if err == nil {
			return d
		}
		if !translate {
			continue
		}
		for name, digit := range digits {
			if strings.HasPrefix(s[i:], name) {
				return digit
			}
		}
	}
	return
}

// We read the lines from end to beginning and check if it's a digit by either int or the word of a digit
// If a word ends with for example 'oneight', it will just read the 8, and not the 1. Cause we read it with
// back to front priority
func getLastDigit(s string, translate bool) (d int) {
	var err error
	for i := len(s) - 1; i >= 0; i-- {
		d, err = strconv.Atoi(fmt.Sprintf("%c", s[i]))
		if err == nil {
			return d
		}
		if !translate {
			continue
		}
		for name, digit := range digits {
			if strings.HasSuffix(s[:i+1], name) {
				return digit
			}
		}
	}
	return
}

func readDigits(input []string, translate bool) (totalCount int) {
	for _, line := range input {
		firstDigit := getFirstDigit(line, translate)
		lastDigit := getLastDigit(line, translate)
		totalCount += (firstDigit * 10) + lastDigit
	}
	return
}

func main() {
	lines := strings.Split(input, "\n")

	log.Printf("Total ammount part1: %d", readDigits(lines, false))
	log.Printf("Total ammount part2: %d", readDigits(lines, true))
}
