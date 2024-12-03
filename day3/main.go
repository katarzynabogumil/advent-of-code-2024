package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	resPart1, err := part1(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("Error in part 1: %w", err))
	}

	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2, err := part2(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("Error in part 1: %w", err))
	}

	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(input string) (int, error) {
	r, _ := regexp.Compile(`(don't)\(\)|(do)\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	arr := r.FindAllStringSubmatch(input, -1)

	sum := 0
	isBlocked := false

	for _, group := range arr {
		switch {
		case group[1] == "don't":
			isBlocked = true
		case group[2] == "do":
			isBlocked = false
		default:
			first, err := strconv.Atoi(group[3])
			if err != nil {
				return 0, err
			}

			second, err := strconv.Atoi(group[4])
			if err != nil {
				return 0, err
			}

			if !isBlocked {
				sum += first * second
			}
		}
	}

	return sum, nil
}

func part1(input string) (int, error) {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	arr := r.FindAllStringSubmatch(input, -1)
	sum := 0

	for _, group := range arr {
		first, err := strconv.Atoi(group[1])
		if err != nil {
			return 0, err
		}

		second, err := strconv.Atoi(group[2])
		if err != nil {
			return 0, err
		}

		sum += first * second
	}

	return sum, nil
}
