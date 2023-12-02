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
			} else {
				start += 1
			}
		}

		if !lastNumber.isSet {
			if isNumber(text[end]) {
				lastNumber.value = text[end]
				lastNumber.isSet = true
			} else {
				end -= 1
			}
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
	var finalScore int
	for scanner.Scan() {
		result := extractNumber(scanner.Text())
		finalScore += result
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", finalScore)
}
