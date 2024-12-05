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

	matrix := parseInput(string(input))

	resPart1 := part1(matrix)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(matrix)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(matrix [][]string) int {
	counter := 0
	for j, row := range matrix {
		for i, letter := range row {
			if letter != "A" {
				continue
			}

			if j < 1 || j > len(matrix)-2 {
				continue
			}

			if i < 1 || i > len(row)-2 {
				continue
			}

			if checkXMAS(matrix, i, j) {
				counter += 1
			}

		}
	}
	return counter
}

func checkXMAS(matrix [][]string, i int, j int) bool {
	if matrix[j-1][i-1] == "M" && matrix[j+1][i+1] == "S" &&
		matrix[j-1][i+1] == "M" && matrix[j+1][i-1] == "S" {
		return true
	}
	if matrix[j-1][i-1] == "M" && matrix[j+1][i+1] == "S" &&
		matrix[j-1][i+1] == "S" && matrix[j+1][i-1] == "M" {
		return true
	}
	if matrix[j-1][i-1] == "S" && matrix[j+1][i+1] == "M" &&
		matrix[j-1][i+1] == "M" && matrix[j+1][i-1] == "S" {
		return true
	}
	if matrix[j-1][i-1] == "S" && matrix[j+1][i+1] == "M" &&
		matrix[j-1][i+1] == "S" && matrix[j+1][i-1] == "M" {
		return true
	}
	return false
}

func part1(matrix [][]string) int {
	counter := 0
	for j, row := range matrix {
		for i, letter := range row {
			if letter != "X" {
				continue
			}

			if j > 2 && checkUp(matrix, i, j) {
				counter += 1
			}
			if j < len(matrix)-3 && checkDown(matrix, i, j) {
				counter += 1
			}
			if i > 2 && checkLeft(matrix, i, j) {
				counter += 1
			}
			if i < len(row)-3 && checkRight(matrix, i, j) {
				counter += 1
			}
			if i < len(row)-3 && j > 2 && checkUpRight(matrix, i, j) {
				counter += 1
			}
			if i > 2 && j > 2 && checkUpLeft(matrix, i, j) {
				counter += 1
			}
			if i < len(row)-3 && j < len(matrix)-3 && checkDownRight(matrix, i, j) {
				counter += 1
			}
			if i > 2 && j < len(matrix)-3 && checkDownLeft(matrix, i, j) {
				counter += 1
			}

		}
	}
	return counter
}

func checkLeft(matrix [][]string, i int, j int) bool {
	if matrix[j][i-1] != "M" || matrix[j][i-2] != "A" || matrix[j][i-3] != "S" {
		return false
	}
	return true
}

func checkRight(matrix [][]string, i int, j int) bool {
	if matrix[j][i+1] != "M" || matrix[j][i+2] != "A" || matrix[j][i+3] != "S" {
		return false
	}
	return true
}

func checkDown(matrix [][]string, i int, j int) bool {
	if matrix[j+1][i] != "M" || matrix[j+2][i] != "A" || matrix[j+3][i] != "S" {
		return false
	}
	return true
}

func checkUp(matrix [][]string, i int, j int) bool {
	if matrix[j-1][i] != "M" || matrix[j-2][i] != "A" || matrix[j-3][i] != "S" {
		return false
	}
	return true
}

func checkDownLeft(matrix [][]string, i int, j int) bool {
	if matrix[j+1][i-1] != "M" || matrix[j+2][i-2] != "A" || matrix[j+3][i-3] != "S" {
		return false
	}
	return true
}

func checkDownRight(matrix [][]string, i int, j int) bool {
	if matrix[j+1][i+1] != "M" || matrix[j+2][i+2] != "A" || matrix[j+3][i+3] != "S" {
		return false
	}
	return true
}

func checkUpLeft(matrix [][]string, i int, j int) bool {
	if matrix[j-1][i-1] != "M" || matrix[j-2][i-2] != "A" || matrix[j-3][i-3] != "S" {
		return false
	}
	return true
}

func checkUpRight(matrix [][]string, i int, j int) bool {
	if matrix[j-1][i+1] != "M" || matrix[j-2][i+2] != "A" || matrix[j-3][i+3] != "S" {
		return false
	}
	return true
}

func parseInput(input string) [][]string {
	var matrix [][]string

	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	return matrix
}
