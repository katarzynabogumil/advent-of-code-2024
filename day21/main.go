package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var Memo = make(map[string]map[int]int)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	codes := parseInput(string(input))

	resPart1 := solve(codes, 2)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := solve(codes, 25)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func solve(codes []string, rounds int) int {
	firstKeyboard := []string{"789", "456", "123", "X0A"}
	firstMap := getInstructionsMap(firstKeyboard, true)

	nextKeyboard := []string{"X^A", "<v>"}
	nextMap := getInstructionsMap(nextKeyboard, false)

	complexity := 0

	for _, code := range codes {
		instructions := getInstructions('A', code, firstMap)
		length := getLength('A', instructions, rounds, nextMap)
		complexity += length * getCodeInt(code)
	}

	return complexity
}

func getLength(prev rune, prevInstructions string, round int, shortestMoves map[rune]map[rune]string) int {
	key := string(prev) + prevInstructions
	if cachesInstr, okInstr := Memo[key]; okInstr {
		if value, okRound := cachesInstr[round]; okRound {
			return value
		}
	} else {
		Memo[key] = make(map[int]int)
	}

	length := 0
	instructions := getInstructions(prev, prevInstructions, shortestMoves)
	if round == 1 {
		length = len(instructions)
	} else {
		prevVal := 'A'
		for _, val := range instructions {
			length += getLength(prevVal, string(val), round-1, shortestMoves)
			prevVal = val
		}
	}

	Memo[key][round] = length
	return length
}

func getInstructions(prevChar rune, prevInstructions string, shortestMap map[rune]map[rune]string) string {
	prev := prevChar
	instructions := ""

	for _, val := range prevInstructions {
		instructions += shortestMap[prev][val]
		prev = val
	}

	return instructions
}

func getInstructionsMap(keyboard []string, first bool) map[rune]map[rune]string {
	shortestMoves := make(map[rune]map[rune]string)

	for _, row := range keyboard {
		for _, start := range row {
			for _, otherRow := range keyboard {
				for _, end := range otherRow {
					if shortestMoves[start] == nil {
						shortestMoves[start] = make(map[rune]string)
					}

					if start == end {
						shortestMoves[start][end] = "A"
					} else if end != 'X' && start != 'X' {
						shortestMoves[start][end] = getShortestInstructions(keyboard, start, end, first)
					}
				}
			}
		}
	}

	return shortestMoves
}

func getShortestInstructions(matrix []string, start rune, end rune, first bool) string {
	startX, startY := getPoint(matrix, start)
	endX, endY := getPoint(matrix, end)

	xInstructions := ""
	for range startX - endX {
		xInstructions += "<"
	}

	for range endX - startX {
		xInstructions += ">"
	}

	yInstructions := ""
	for range startY - endY {
		yInstructions += "^"
	}

	for range endY - startY {
		yInstructions += "v"
	}

	aY := 0
	if first {
		aY = 3
	}

	instructions := ""
	if startY == aY && endX == 0 {
		instructions += yInstructions + xInstructions
	} else if startX == 0 && endY == aY {
		instructions += xInstructions + yInstructions
	} else if endX-startX < 0 {
		instructions += xInstructions + yInstructions
	} else {
		instructions += yInstructions + xInstructions
	}
	instructions += "A"

	return instructions
}

func getPoint(matrix []string, value rune) (int, int) {
	for y, row := range matrix {
		for x, val := range row {
			if val == value {
				return x, y
			}
		}
	}
	return 0, 0
}

func getCodeInt(code string) int {
	res := 0
	multiplier := 100
	for _, val := range code {
		num, _ := strconv.Atoi(string(val))
		res += num * multiplier
		multiplier /= 10
	}
	return res
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(string(input)), "\n")
}
