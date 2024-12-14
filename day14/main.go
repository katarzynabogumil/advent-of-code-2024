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

var (
	Vectors = []Point{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
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

	resPart1 := part1(copyInput(parsedInput), 101, 103, 100)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(parsedInput, 101, 103)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(robots [][]Point, sizeX int, sizeY int) int {
	seconds := 1

	for {
		for _, robot := range robots {
			robot[0] = Point{(robot[0].x + robot[1].x + sizeX) % sizeX, (robot[0].y + robot[1].y + sizeY) % sizeY}
		}

		if checkConnected(robots, sizeX, sizeY) > 100 {
			return seconds
		}

		seconds++
	}
}

func checkConnected(robots [][]Point, sizeX int, sizeY int) int {
	checkedRobots := map[Point]bool{}
	counter := 0

	for _, robot := range robots {
		if !checkedRobots[robot[0]] {
			thisCounter := checkNextRobot(robots, robot[0], sizeX, sizeY, 0, &checkedRobots)
			if thisCounter > counter {
				counter = thisCounter
			}
		}
	}
	return counter
}

func checkNextRobot(robots [][]Point, robot Point, sizeX int, sizeY int, counter int, checkedRobots *map[Point]bool) int {
	(*checkedRobots)[robot] = true
	counter++

	for _, vector := range Vectors {
		x := robot.x + vector.x
		y := robot.y + vector.y
		if x >= 0 && y >= 0 && x < sizeX && y < sizeY {
			nextRobot := Point{x, y}
			if isRobot(robots, nextRobot) && !(*checkedRobots)[nextRobot] {
				counter = checkNextRobot(robots, nextRobot, sizeX, sizeY, counter, checkedRobots)
			}
		}
	}

	return counter
}

func isRobot(robots [][]Point, robotToCheck Point) bool {
	for _, robot := range robots {
		if robot[0].x == robotToCheck.x && robot[0].y == robotToCheck.y {
			return true
		}
	}
	return false
}

func part1(robots [][]Point, sizeX int, sizeY int, seconds int) int {
	for range seconds {
		for _, robot := range robots {
			robot[0] = Point{(robot[0].x + robot[1].x + sizeX) % sizeX, (robot[0].y + robot[1].y + sizeY) % sizeY}
		}
	}
	return calculateSafety(&robots, sizeX, sizeY)
}

func calculateSafety(robots *[][]Point, sizeX int, sizeY int) int {
	sizeX1 := []int{0, sizeX / 2}
	sizeX2 := []int{sizeX - sizeX/2, sizeX}
	sizeY1 := []int{0, sizeY / 2}
	sizeY2 := []int{sizeY - sizeY/2, sizeY}

	res := 1

	for _, sizeY := range [][]int{sizeY1, sizeY2} {
		for _, sizeX := range [][]int{sizeX1, sizeX2} {
			count := 0

			for _, robot := range *robots {
				if robot[0].x >= sizeX[0] && robot[0].x < sizeX[1] &&
					robot[0].y >= sizeY[0] && robot[0].y < sizeY[1] {
					count++
				}
			}

			res *= count
		}
	}

	return res
}

func copyInput(input [][]Point) [][]Point {
	copiedInput := [][]Point{}
	for _, line := range input {
		copiedLine := make([]Point, len(line))
		copy(copiedLine, line)
		copiedInput = append(copiedInput, copiedLine)
	}
	return copiedInput
}

func parseInput(input string) ([][]Point, error) {
	lines := [][]Point{}

	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var arr []Point

		lineArr := strings.Split(line, " ")
		point := strings.Replace(lineArr[0], "p=", "", 1)
		velocity := strings.Replace(lineArr[1], "v=", "", 1)

		for _, val := range []string{point, velocity} {
			nums := strings.Split(val, ",")
			x, err := strconv.Atoi(nums[0])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(nums[1])
			if err != nil {
				return nil, err
			}

			arr = append(arr, Point{x, y})
		}

		lines = append(lines, arr)
	}

	return lines, nil
}

// func printMatrix(robots [][]Point, sizeX int, sizeY int) {
// 	for y := range sizeY {
// 		line := []int{}

// 		for x := range sizeX {
// 			count := 0
// 			for _, robot := range robots {
// 				if robot[0].x == x && robot[0].y == y {
// 					count++
// 				}
// 			}
// 			line = append(line, count)
// 		}

// 		for _, val := range line {
// 			if val == 0 {
// 				fmt.Print(".")
// 			} else {
// 				fmt.Print(val)
// 			}
// 		}

// 		fmt.Print("\n")
// 	}

// 	fmt.Print("\n")
// }
