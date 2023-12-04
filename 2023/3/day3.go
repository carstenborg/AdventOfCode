package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Start by reading 3 lines, and then have a sliding window of 1 until EOF
	// Store the indexs of any char, that is not a . or a digit for each line
	// For each number on each line, get the start and end index
	// If the start index (-1) or the end index (+1) matches any symbol indexes on either of the 3 lines, it's valid!

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error occured while opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for index, runeValue := range line {
			fmt.Printf("Index: %d, Rune: %c\n", index, runeValue)
		}
	}

}
