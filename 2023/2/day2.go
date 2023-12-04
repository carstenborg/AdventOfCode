package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	maxRedCubes := 12
	maxGreenCubes := 13
	maxBlueCubes := 14

	allowedGameIdSums := 0
	leastNeededCubesProductSum := 0

	// Regex go get the game input and the game id
	gameInputRegexPattern := regexp.MustCompile(`Game\s*(?P<gameid>\d*):\s(?P<input>.*)`)
	pickInputRegexPattern := regexp.MustCompile(`^\s*(?P<number>\d*)\s*(?P<colour>green|red|blue)\s*$`)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error occured while opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		gameIdMatch := gameInputRegexPattern.FindStringSubmatch(line)
		gameId := gameIdMatch[gameInputRegexPattern.SubexpIndex("gameid")]

		gameIdNumber, err := strconv.Atoi(gameId)
		if err != nil {
			fmt.Print("Error while parsing game id int: ", gameId, " - ", err)
			return
		}

		gameInput := gameIdMatch[gameInputRegexPattern.SubexpIndex("input")]

		gameAllowed := true
		leastNeededGreenCubes := 0
		leastNeededBlueCubes := 0
		leastNeededRedCubes := 0

		// Get the input for each individual round
		rounds := strings.Split(gameInput, ";")
		for _, round := range rounds {

			roundGreenCubes := 0
			roundBlueCubes := 0
			roundRedCubes := 0

			// fmt.Println("Round: ", round)
			picks := strings.Split(round, ",")
			for _, pick := range picks {

				pickMatch := pickInputRegexPattern.FindStringSubmatch(pick)
				number := strings.TrimSpace(pickMatch[pickInputRegexPattern.SubexpIndex("number")])
				colour := strings.TrimSpace(pickMatch[pickInputRegexPattern.SubexpIndex("colour")])

				intNumber, err := strconv.Atoi(number)
				if err != nil {
					fmt.Print("Error while parsing int: ", number, " - ", err)
					return
				}

				switch colour {
				case "red":
					roundRedCubes += intNumber

					if intNumber > leastNeededRedCubes {
						leastNeededRedCubes = intNumber
					}

				case "green":
					roundGreenCubes += intNumber

					if intNumber > leastNeededGreenCubes {
						leastNeededGreenCubes = intNumber
					}

				case "blue":
					roundBlueCubes += intNumber

					if intNumber > leastNeededBlueCubes {
						leastNeededBlueCubes = intNumber
					}
				}
			}

			if roundRedCubes > maxRedCubes || roundBlueCubes > maxBlueCubes || roundGreenCubes > maxGreenCubes {
				gameAllowed = false

				// Due to part 2, we cannot break out of each game, even if it's not allowed,
				// as we need to locate the least possible amount of cubes neccessary
				// break
			}

		}

		if gameAllowed {
			allowedGameIdSums += gameIdNumber
		}

		gameCubeProduct := leastNeededBlueCubes * leastNeededGreenCubes * leastNeededRedCubes
		leastNeededCubesProductSum += gameCubeProduct

	}

	fmt.Println("allowedGameIdSums:", allowedGameIdSums)
	fmt.Println("leastNeededCubesProductSum:", leastNeededCubesProductSum)

}
