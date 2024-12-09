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
	moveBlocksPart2(&input)
	return checkSum(&input)
}

func moveBlocksPart2(input *[]string) {
	for j := len(*input) - 1; j > 0; j-- {
		if (*input)[j] != "." {
			blockLength := getBlockLength((*input)[j], j, input)

			for i := 0; i < j; i++ {
				if (*input)[i] == "." {
					spaceLength := getSpaceLength(i, input)

					if blockLength <= spaceLength {
						for k := 0; k < blockLength; k++ {
							(*input)[i+k] = (*input)[j-k]
							(*input)[j-k] = "."
						}
						break
					}
					i += spaceLength
				}
			}
			j -= blockLength - 1
		}
	}
}

func getSpaceLength(idx int, input *[]string) int {
	counter := 1
	for i := idx + 1; i < len(*input); i++ {
		if (*input)[i] == "." {
			counter++
		} else {
			break
		}
	}
	return counter
}

func getBlockLength(value string, idx int, input *[]string) int {
	counter := 1
	for i := idx - 1; i > 0; i-- {
		if (*input)[i] == value {
			counter++
		} else {
			break
		}
	}
	return counter
}

func part1(input []string) int {
	moveBlocksPart1(&input)
	return checkSum(&input)
}

func moveBlocksPart1(input *[]string) {
	for i := 0; i < len(*input); i++ {
		res := checkForDigits((*input)[i:])
		if !res {
			break
		}

		if (*input)[i] == "." {
			for j := len(*input) - 1; j > 0; j-- {
				if (*input)[j] != "." {
					(*input)[i] = (*input)[j]
					(*input)[j] = "."
					break
				}
			}
		}
	}
}

func checkForDigits(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != "." {
			return true
		}
	}
	return false
}

func checkSum(input *[]string) int {
	sum := 0
	for i := 0; i < len(*input); i++ {
		value := (*input)[i]
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

		for j := 0; j < count; j++ {
			arr = append(arr, value)
		}
	}

	return arr, nil
}
