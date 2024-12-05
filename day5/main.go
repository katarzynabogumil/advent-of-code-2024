package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't open file"))
	}

	scanner := bufio.NewScanner(file)

	lines, rulesAfter, rulesBefore, err := parseInput(scanner)
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1, incorrectLines := part1(lines, rulesAfter, rulesBefore)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(incorrectLines, rulesAfter, rulesBefore)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(lines [][]int, rulesAfter *map[int][]int, rulesBefore *map[int][]int) int {
	sum := 0

	for _, line := range lines {
		sortedLine := sort(line, rulesAfter, rulesBefore)
		sum += sortedLine[len(sortedLine)/2]
	}

	return sum
}

func sort(line []int, rulesAfter *map[int][]int, rulesBefore *map[int][]int) []int {
	isCorrect := true

	for i, currNum := range line {
		if i == 0 {
			continue
		}

		numbersBefore := line[:i]
		numbersAfter := line[i+1:]

		ruleBefore := (*rulesBefore)[currNum]
		ruleAfter := (*rulesAfter)[currNum]

		for j, n := range numbersBefore {
			if slices.Contains(ruleAfter, n) {
				isCorrect = false
				temp := line[j]
				line[j] = line[i]
				line[i] = temp
			}
		}

		for j, n := range numbersAfter {
			if slices.Contains(ruleBefore, n) {
				isCorrect = false
				temp := line[j]
				line[j] = line[i]
				line[i] = temp
			}
		}

	}

	if isCorrect {
		return line
	} else {
		return sort(line, rulesAfter, rulesBefore)
	}
}

func part1(lines [][]int, rulesAfter *map[int][]int, rulesBefore *map[int][]int) (int, [][]int) {
	sum := 0
	var incorrectLines [][]int

	for _, line := range lines {
	out:
		for i, currNum := range line {
			if i == 0 {
				continue
			}
			numbersBefore := line[:i]
			numbersAfter := line[i+1:]

			ruleBefore := (*rulesBefore)[currNum]
			ruleAfter := (*rulesAfter)[currNum]

			for _, n := range numbersBefore {
				if slices.Contains(ruleAfter, n) {
					incorrectLines = append(incorrectLines, line)
					break out
				}
			}

			for _, n := range numbersAfter {
				if slices.Contains(ruleBefore, n) {
					incorrectLines = append(incorrectLines, line)
					break out
				}
			}

			if i == len(line)-1 {
				sum += line[len(line)/2]
			}
		}
	}

	return sum, incorrectLines
}

func parseInput(scanner *bufio.Scanner) ([][]int, *map[int][]int, *map[int][]int, error) {
	var lines [][]int
	secondPart := false
	rulesBefore := make(map[int][]int)
	rulesAfter := make(map[int][]int)

	for scanner.Scan() {
		if scanner.Text() == "" {
			secondPart = true
		} else if !secondPart {
			line := strings.Split(scanner.Text(), "|")
			i, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, nil, nil, err
			}
			j, err := strconv.Atoi(line[1])
			if err != nil {
				return nil, nil, nil, err
			}
			rulesAfter[i] = append(rulesAfter[i], j)
			rulesBefore[j] = append(rulesBefore[j], i)

		} else {
			var line []int
			arr := strings.Split(scanner.Text(), ",")
			for _, num := range arr {
				i, err := strconv.Atoi(num)
				if err != nil {
					return nil, nil, nil, err
				}
				line = append(line, i)
			}
			lines = append(lines, line)
		}
	}

	return lines, &rulesAfter, &rulesBefore, nil
}
