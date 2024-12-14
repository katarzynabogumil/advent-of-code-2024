package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type FloatPoint struct {
	x, y float64
}

var (
	aCost      = 3
	bCost      = 1
	limit      = 100
	max        = aCost*limit + bCost*limit + 1
	addToPrize = 10000000000000
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	parsedInput, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(parsedInput)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v, in %s\n", resPart1, timePart1)

	resPart2 := part2(parsedInput)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(input [][]Point) int {
	tokens := 0
	for _, machine := range input {
		a := FloatPoint{float64(machine[0].x), float64(machine[0].y)}
		b := FloatPoint{float64(machine[1].x), float64(machine[1].y)}
		c := FloatPoint{float64(machine[2].x + addToPrize), float64(machine[2].y + addToPrize)}
		tokens += countTokens(a, b, c)
	}
	return tokens
}

func countTokens(a FloatPoint, b FloatPoint, prize FloatPoint) int {
	aClicksFloat := (prize.y - b.y*prize.x/b.x) / (a.y - b.y*a.x/b.x)
	bClicksFloat := (prize.x - a.x*aClicksFloat) / b.x
	bClicks := int(bClicksFloat + 0.5)
	aClicks := int(aClicksFloat + 0.5)

	if aClicks < 0 || bClicks < 0 {
		return 0
	}

	if aClicks*int(a.x)+bClicks*int(b.x) == int(prize.x) && aClicks*int(a.y)+bClicks*int(b.y) == int(prize.y) {
		return aClicks*aCost + bClicks*bCost
	}

	return 0
}

func part1(input [][]Point) int {
	tokens := 0
	for _, machine := range input {
		checkedPositions := map[Point]int{}
		thisTokens := checkNextClick(machine[0], machine[1], machine[2], 0, 0, Point{0, 0}, 0, checkedPositions)
		if thisTokens != max {
			tokens += thisTokens
		}
	}
	return tokens
}

func checkNextClick(a Point, b Point, prize Point, aClicks int, bClicks int, position Point, tokens int, checked map[Point]int) int {
	if val, ok := checked[position]; ok && val <= tokens {
		return max
	}
	checked[position] = tokens

	if position.x == prize.x && position.y == prize.y {
		return tokens
	}

	if position.x > prize.x || position.y > prize.y {
		return max
	}

	return min(
		checkNextClick(a, b, prize, aClicks+1, bClicks, Point{position.x + a.x, position.y + a.y}, tokens+aCost, checked),
		checkNextClick(a, b, prize, aClicks, bClicks+1, Point{position.x + b.x, position.y + b.y}, tokens+bCost, checked),
	)
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func parseInput(input string) ([][]Point, error) {
	lines := [][]Point{}
	inputLines := strings.Split(strings.TrimSpace(string(input)), "\n")

	for i := 0; i < len(inputLines); i += 4 {
		var arr []Point

		for j := range 3 {
			lineStr := inputLines[i+j]
			switch j {
			case 0:
				lineStr = strings.Replace(lineStr, "Button A: ", "", 1)
			case 1:
				lineStr = strings.Replace(lineStr, "Button B: ", "", 1)
			case 2:
				lineStr = strings.Replace(lineStr, "Prize: ", "", 1)
			}

			lineArr := strings.Split(lineStr, ", ")

			x, err := strconv.Atoi(lineArr[0][2:])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(lineArr[1][2:])
			if err != nil {
				return nil, err
			}

			arr = append(arr, Point{x, y})
		}

		lines = append(lines, arr)

	}

	return lines, nil
}
