package main

import (
	"fmt"
	"math"
	"os"
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
	resPart1 := solve(parsedInput, 25)
	timePart1 := time.Since(startPart1)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	startPart2 := time.Now()
	resPart2 := solve(parsedInput, 75)
	timePart2 := time.Since(startPart2)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func solve(input []int, rounds int) int {
	count := 0

	// Store previously calculated values for one round step
	valuesMemo := make(map[int][]int)

	// Store all values for digits for all steps under rounds / 2
	memoLimit := rounds / 2
	digitMemo := prepareDigitMemo(memoLimit, valuesMemo)

	for _, num := range input {
		count += getCount(num, rounds, digitMemo, valuesMemo)
	}

	return count
}

func prepareDigitMemo(limit int, valuesMemo map[int][]int) map[[2]int]int {
	memo := make(map[[2]int]int)

	for num := range 10 {
		for i := range limit {
			memo[[2]int{num, i}] = getCount(num, i+1, memo, valuesMemo)
		}
	}

	return memo
}

func getCount(initialNum int, rounds int, digitsMemo map[[2]int]int, valuesMemo map[int][]int) int {
	count := 0
	prevArr := []int{initialNum}

	for i := range rounds {
		arr := []int{}

		for _, num := range prevArr {
			if val, ok := digitsMemo[[2]int{num, rounds - i - 1}]; ok {
				count += val
			} else {
				arr = append(arr, getNewValue(num, valuesMemo)...)
			}
		}

		prevArr = arr
	}

	return count + len(prevArr)
}

func getNewValue(num int, memo map[int][]int) []int {
	if val, ok := memo[num]; ok {
		return val
	}

	if num == 0 {
		value := []int{1}
		memo[num] = value
		return value
	}

	length := getLenth(num)

	if length%2 == 0 {
		value := split(num, length)
		memo[num] = value
		return value
	} else {
		value := []int{num * 2024}
		memo[num] = value
		return value
	}
}

func split(num int, length int) []int {
	halfLen := length / 2

	multiplicator := int(math.Pow(10, float64(halfLen)))
	num1 := num / multiplicator
	num2 := num - num1*multiplicator

	return []int{num1, num2}
}

func getLenth(i int) int {
	if i == 0 {
		return 1
	}

	length := 0
	for i != 0 {
		i /= 10
		length++
	}

	return length
}

func parseInput(input string) ([]int, error) {
	var arr []int
	for _, val := range strings.Split(string(input), " ") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}
	return arr, nil
}
