package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func findFirstValueStringOccurence(input string, mapToLookup map[string]int) (int, int) {
	currentFirstIndex := math.MaxInt
	currentLastIndex := math.MinInt

	currentFirstDigit := 0
	currentLastDigit := 0

	// Iterate through all maps
	for key, value := range mapToLookup {
		firstIndex := strings.Index(input, key)
		if firstIndex != -1 && firstIndex < currentFirstIndex {
			currentFirstIndex = firstIndex
			currentFirstDigit = value
		}

		lastIndex := strings.LastIndex(input, key)
		if lastIndex != -1 && lastIndex > currentLastIndex {
			currentLastIndex = lastIndex
			currentLastDigit = value
		}
	}

	return currentFirstDigit, currentLastDigit

}

func main() {

	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error occured while opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line to a rune slice
		runes := []rune(line)

		// Set the first digit
		var firstDigit rune
		for _, char := range runes {
			if unicode.IsDigit(char) {
				firstDigit = char
				break
			}
		}

		// Set the second digit by iterating backwards
		var secondDigit rune
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				secondDigit = runes[i]
				break
			}
		}

		// Multiply the first digit with 10, as we need to combined literal number
		first := int(firstDigit-'0') * 10
		second := int(secondDigit - '0')
		sum += first + second
	}

	fmt.Println("Part 1: Total sum: ", fmt.Sprint(sum))

	// Part 2: Create a dictionary (map) of all the possible string occurences
	valueMap := make(map[string]int)
	valueMap["one"] = 1
	valueMap["two"] = 2
	valueMap["three"] = 3
	valueMap["four"] = 4
	valueMap["five"] = 5
	valueMap["six"] = 6
	valueMap["seven"] = 7
	valueMap["eight"] = 8
	valueMap["nine"] = 9

	valueMap["1"] = 1
	valueMap["2"] = 2
	valueMap["3"] = 3
	valueMap["4"] = 4
	valueMap["5"] = 5
	valueMap["6"] = 6
	valueMap["7"] = 7
	valueMap["8"] = 8
	valueMap["9"] = 9

	newSum := 0

	newFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error occured while opening file", err)
		return
	}
	defer newFile.Close()

	newScanner := bufio.NewScanner(newFile)
	for newScanner.Scan() {

		firstVal, secondVal := findFirstValueStringOccurence(newScanner.Text(), valueMap)
		sumToAdd := (firstVal * 10) + secondVal
		newSum += sumToAdd
	}

	fmt.Println("Part 2: Total sum: ", fmt.Sprint(newSum))
}
