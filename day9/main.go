package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	parsedInput, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	startPart1 := time.Now()
	resPart1 := part1(slices.Clone(parsedInput))
	timePart1 := time.Since(startPart1)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	startPart2 := time.Now()
	resPart2 := part2(parsedInput)
	timePart2 := time.Since(startPart2)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(input []string) int {
	moveBlocksPart2(input)
	return checkSum(input)
}

func moveBlocksPart2(input []string) {
	for i := len(input) - 1; i > 0; i-- {
		if input[i] != "." {
			blockLength := getBlockLength(input[i], i, input)

			for j := 0; j < i; j++ {
				if input[j] == "." {
					spaceLength := getSpaceLength(j, input)

					if blockLength <= spaceLength {
						for k := range blockLength {
							input[j+k] = input[i-k]
							input[i-k] = "."
						}
						break
					}
					j += spaceLength
				}
			}
			i -= blockLength - 1
		}
	}
}

func getSpaceLength(idx int, input []string) int {
	counter := 1
	for i := idx + 1; i < len(input); i++ {
		if input[i] == "." {
			counter++
		} else {
			break
		}
	}
	return counter
}

func getBlockLength(value string, idx int, input []string) int {
	counter := 1
	for i := idx - 1; i > 0; i-- {
		if input[i] == value {
			counter++
		} else {
			break
		}
	}
	return counter
}

func part1(input []string) int {
	moveBlocksPart1(input)
	return checkSum(input)
}

func moveBlocksPart1(input []string) {
	for i, value := range input {
		res := checkForDigits(input[i:])
		if !res {
			break
		}

		if value == "." {
			for j := len(input) - 1; j > 0; j-- {
				if input[j] != "." {
					input[i] = input[j]
					input[j] = "."
					break
				}
			}
		}
	}
}

func checkForDigits(arr []string) bool {
	for _, value := range arr {
		if value != "." {
			return true
		}
	}
	return false
}

func checkSum(input []string) int {
	sum := 0
	for i, value := range input {
		if value == "." {
			continue
		}
		num, _ := strconv.Atoi(value)
		sum += i * num
	}
	return sum
}

func parseInput(input string) ([]string, error) {
	var arr []string
	index := 0

	line := strings.Split(input, "")

	for i, str := range line {
		count, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		var value string
		if i%2 == 0 {
			value = strconv.Itoa(index)
			index++
		} else {
			value = "."
		}

		for range count {
			arr = append(arr, value)
		}
	}

	return arr, nil
}
