package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	patterns, designs, shortestPattern, longestPattern := parseInput(string(input))

	resPart1, resPart2 := solve(patterns, designs, shortestPattern, longestPattern)
	time := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, time)
	fmt.Printf("Part 2: %v in %s\n", resPart2, time)
}

func solve(patterns map[string]bool, designs []string, shortestPattern int, longestPattern int) (int, int) {
	counter := 0
	allPossibilities := 0
	checked := make(map[string]int)

	for _, design := range designs {
		possibilities := checkDesign(patterns, design, shortestPattern, longestPattern, checked)
		allPossibilities += possibilities
		if possibilities > 0 {
			counter++
		}
	}

	return counter, allPossibilities
}

func checkDesign(patterns map[string]bool, design string, shortestPattern int, longestPattern int, checked map[string]int) int {
	if val, ok := checked[design]; ok {
		return val
	}

	possibilities := 0
	for j := shortestPattern; j <= longestPattern; j++ {
		if j <= len(design) {
			subDesign := design[0:j]

			if patterns[subDesign] {
				if j == len(design) {
					possibilities++
				} else {
					possibilities += checkDesign(patterns, design[j:], shortestPattern, longestPattern, checked)
				}
			}
		}
	}

	checked[design] = possibilities
	return possibilities
}

func parseInput(input string) (map[string]bool, []string, int, int) {
	patterns := make(map[string]bool, 0)
	designs := make([]string, 0)
	shortestPattern := 0
	longestPattern := 0

	isParsingDesigns := false
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if len(line) == 0 {
			isParsingDesigns = true
		} else if isParsingDesigns {
			designs = append(designs, line)
		} else {
			shortestPattern = len(line)
			for _, pattern := range strings.Split(line, ", ") {
				patterns[pattern] = true
				if len(pattern) < shortestPattern {
					shortestPattern = len(pattern)
				}
				if len(pattern) > longestPattern {
					longestPattern = len(pattern)
				}
			}
		}
	}

	return patterns, designs, shortestPattern, longestPattern
}
