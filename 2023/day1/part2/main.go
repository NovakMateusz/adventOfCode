package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Number struct {
	value uint8
	isSet bool
}

var LOOKUP_MAP = map[string]uint8{
	"one":   49,
	"two":   50,
	"three": 51,
	"four":  52,
	"five":  53,
	"six":   54,
	"seven": 55,
	"eight": 56,
	"nine":  57,
}

func isNumber(val uint8) bool {
	// Function checks if passed ASCII value corresponds to the number from 1-9
	if val >= 49 && val <= 57 {
		return true
	}
	return false
}

func parseToInt(firstNumber, lastNumber Number) int {
	intVar, err := strconv.Atoi(fmt.Sprintf("%c%c", firstNumber.value, lastNumber.value))
	if err != nil {
		log.Fatal(err)
	}
	return intVar
}

func extractNumber(text string) int {
	start, end := 0, len(text)-1
	firstNumber, lastNumber := Number{isSet: false}, Number{isSet: false}
	for start <= end {
		if !firstNumber.isSet {

			if isNumber(text[start]) {
				firstNumber.value = text[start]
				firstNumber.isSet = true
				continue
			}

			if start+3 < len(text) {

				value, ok := LOOKUP_MAP[text[start:start+3]]
				if ok {
					firstNumber.value = value
					firstNumber.isSet = true
					continue
				}
			}
			if start+4 < len(text) {
				value, ok := LOOKUP_MAP[text[start:start+4]]
				if ok {
					firstNumber.value = value
					firstNumber.isSet = true
					continue
				}
			}
			if start+5 < len(text) {
				value, ok := LOOKUP_MAP[text[start:start+5]]
				if ok {
					firstNumber.value = value
					firstNumber.isSet = true
					continue
				}
			}
			start += 1

		}

		if !lastNumber.isSet {
			if isNumber(text[end]) {
				lastNumber.value = text[end]
				lastNumber.isSet = true
				continue
			}

			if end-3 >= 0 {

				value, ok := LOOKUP_MAP[text[end-2:end+1]]
				if ok {
					lastNumber.value = value
					lastNumber.isSet = true
					continue
				}
			}

			if end-4 >= 0 {

				value, ok := LOOKUP_MAP[text[end-3:end+1]]
				if ok {
					lastNumber.value = value
					lastNumber.isSet = true
					continue
				}
			}

			if end-5 >= 0 {

				value, ok := LOOKUP_MAP[text[end-4:end+1]]
				if ok {
					lastNumber.value = value
					lastNumber.isSet = true
					continue
				}
			}
			end -= 1
		}

		if firstNumber.isSet && lastNumber.isSet {
			break
		}

	}

	return parseToInt(firstNumber, lastNumber)
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var finalScore int = 0
	for scanner.Scan() {
		result := extractNumber(scanner.Text())
		finalScore += result
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", finalScore)
}
