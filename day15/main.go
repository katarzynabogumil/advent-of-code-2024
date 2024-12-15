package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	matrixP1, robotP1, matrixP2, robotP2, instructions := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := solve(matrixP1, robotP1, instructions, false)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := solve(matrixP2, robotP2, instructions, true)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func solve(matrix [][]string, robot Point, instructions []Point, part2 bool) int {
	for _, instruction := range instructions {
		robot = moveRobot(matrix, robot, instruction, part2)
	}
	return countCoordinates(matrix)
}

func moveRobot(matrix [][]string, robot Point, instruction Point, part2 bool) Point {
	x := robot.x + instruction.x
	y := robot.y + instruction.y
	switch matrix[y][x] {
	case "#":
		return robot
	case ".":
		return Point{x, y}
	default:
		if moveBoxes(matrix, robot, instruction, part2) {
			return Point{x, y}
		} else {
			return robot
		}
	}
}

func moveBoxes(matrix [][]string, robot Point, instruction Point, part2 bool) bool {
	firstBox := Point{robot.x + instruction.x, robot.y + instruction.y}
	x := firstBox.x + instruction.x
	y := firstBox.y + instruction.y

	for {
		switch matrix[y][x] {
		case "#":
			return false
		case ".":
			if !part2 {
				matrix[firstBox.y][firstBox.x] = "."
				matrix[y][x] = "O"
				return true
			}

			if instruction.y == 0 {
				moveBoxesHorizontally(matrix, robot, Point{x, y}, instruction)
				return true
			}

			return moveBoxesVertically(matrix, robot, instruction)
		}

		x = x + instruction.x
		y = y + instruction.y
	}
}

func moveBoxesHorizontally(matrix [][]string, start Point, stop Point, instruction Point) {
	y := start.y
	for x := stop.x - instruction.x; x != start.x; x -= instruction.x {
		matrix[y][x+instruction.x] = matrix[y][x]
	}

	matrix[y][start.x+instruction.x] = "."
}

func moveBoxesVertically(matrix [][]string, start Point, instruction Point) bool {
	boxesToMove := getBoxesToMove(matrix, start, instruction)
	if len(boxesToMove) == 0 {
		return false
	}

	for i := len(boxesToMove) - 1; i >= 0; i-- {
		box := boxesToMove[i]
		matrix[box.y][box.x] = "."
		matrix[box.y][box.x+1] = "."
		matrix[box.y+instruction.y][box.x] = "["
		matrix[box.y+instruction.y][box.x+1] = "]"

	}

	return true
}

func getBoxesToMove(matrix [][]string, start Point, instruction Point) []Point {
	firstBox := getBoxStart(matrix, Point{start.x, start.y + instruction.y})
	allBoxes := []Point{firstBox}
	lastBoxes := []Point{firstBox}

	for {
		nextBoxes := []Point{}
		allFree := true

		for _, box := range lastBoxes {
			y := box.y + instruction.y

			for _, x := range []int{box.x, box.x + 1} {
				if matrix[y][x] == "#" {
					return nil
				}

				if matrix[y][x] == "[" || matrix[y][x] == "]" {
					allFree = false
					start := getBoxStart(matrix, Point{x, y})
					if !slices.Contains(nextBoxes, start) {
						nextBoxes = append(nextBoxes, start)
					}
				}
			}
		}

		if allFree {
			return allBoxes
		}

		allBoxes = append(allBoxes, nextBoxes...)
		lastBoxes = nextBoxes
	}
}

func getBoxStart(matrix [][]string, point Point) Point {
	if matrix[point.y][point.x] == "[" {
		return Point{point.x, point.y}
	} else {
		return Point{point.x - 1, point.y}
	}
}

func countCoordinates(matrix [][]string) int {
	count := 0
	for j, line := range matrix {
		for i, value := range line {
			if value == "O" {
				count += j*100 + i
			}
			if value == "[" {
				count += j*100 + i
			}
		}
	}
	return count
}

func parseInput(input string) ([][]string, Point, [][]string, Point, []Point) {
	matrixP1 := [][]string{}
	matrixP2 := [][]string{}
	parstingInstructions := false
	instructions := []Point{}
	var startP1 Point
	var startP2 Point

	for j, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if len(line) == 0 {
			parstingInstructions = true
		}

		if parstingInstructions {
			for _, val := range strings.Split(line, "") {
				var vector Point
				switch val {
				case "<":
					vector = Point{-1, 0}
				case ">":
					vector = Point{1, 0}
				case "^":
					vector = Point{0, -1}
				case "v":
					vector = Point{0, 1}
				}
				instructions = append(instructions, vector)
			}

		} else {
			arrP1 := strings.Split(line, "")
			arrP2 := make([]string, len(arrP1)*2)

			for i, value := range arrP1 {
				if value == "@" {
					startP1 = Point{i, j}
					arrP1[i] = "."

					startP2 = Point{i * 2, j}
					arrP2[i*2] = "."
					arrP2[i*2+1] = "."

				} else if value == "O" {
					arrP2[i*2] = "["
					arrP2[i*2+1] = "]"

				} else {
					arrP2[i*2] = value
					arrP2[i*2+1] = value
				}
			}

			matrixP1 = append(matrixP1, arrP1)
			matrixP2 = append(matrixP2, arrP2)
		}
	}

	return matrixP1, startP1, matrixP2, startP2, instructions
}

// func printMatrix(matrix [][]string, robot Point) {
// 	for j, line := range matrix {
// 		for i, val := range line {
// 			if robot.x == i && robot.y == j {
// 				fmt.Print("@")
// 			} else {
// 				fmt.Print(val)
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// }
