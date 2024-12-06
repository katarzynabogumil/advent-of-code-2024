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

var (
	Path       = []Point{}
	NextVector = map[string]Point{
		"-1,0": {0, -1},
		"0,-1": {1, 0},
		"1,0":  {0, 1},
		"0,1":  {-1, 0},
	}
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	matrix, point := parseInput(string(input))

	resPart1 := part1(matrix, point)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(matrix)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(matrix *[][]string) int {
	counter := 0
	obstacles := map[string]int{}

	for i := 0; i < len(Path)-1; i++ {
		p := Point{Path[i].x, Path[i].y}
		nextP := Point{Path[i+1].x, Path[i+1].y}

		val := (*matrix)[nextP.y][nextP.x]
		(*matrix)[nextP.y][nextP.x] = "#"
		vector := Point{nextP.x - p.x, nextP.y - p.y}

		if obstacles[str(nextP)] != 1 && checkIfLoop(matrix, p, vector) && !slices.Contains(Path[:i], nextP) {
			obstacles[str(nextP)] = 1
			counter += 1
		}

		(*matrix)[nextP.y][nextP.x] = val
	}

	// printMatrix(matrix, Point{0, 0}, Point{0, 1})

	return counter
}

func checkIfLoop(matrix *[][]string, p1 Point, vector1 Point) bool {
	corners := map[string]int{}

	point := p1
	vector := vector1
	for {
		if corners[strVector(point, vector)] == 1 {
			return true
		}

		check, nextPoint := checkIfNextObstacle(matrix, point, vector)
		if !check {
			return false
		}

		corners[strVector(point, vector)] = 1
		vector = NextVector[str(vector)]
		point = nextPoint
	}
}

func checkIfNextObstacle(matrix *[][]string, p Point, vector Point) (bool, Point) {
	switch {
	case vector.x == -1 && vector.y == 0:
		for i := p.x - 1; i >= 0; i-- {
			if (*matrix)[p.y][i] == "#" {
				return true, Point{i + 1, p.y}
			}
		}
		return false, Point{0, 0}
	case vector.x == 0 && vector.y == -1:
		for i := p.y - 1; i >= 0; i-- {
			if (*matrix)[i][p.x] == "#" {
				return true, Point{p.x, i + 1}
			}
		}
		return false, Point{0, 0}
	case vector.x == 1 && vector.y == 0:
		for i := p.x + 1; i < len((*matrix)[0]); i++ {
			if (*matrix)[p.y][i] == "#" {
				return true, Point{i - 1, p.y}
			}
		}
		return false, Point{0, 0}
	case vector.x == 0 && vector.y == 1:
		for i := p.y + 1; i < len(*matrix); i++ {
			if (*matrix)[i][p.x] == "#" {
				return true, Point{p.x, i - 1}
			}
		}
		return false, Point{0, 0}
	}
	return false, Point{0, 0}
}

func part1(matrix *[][]string, p Point) int {
	return 1 + traverseMatrix(matrix, p, Point{0, -1})
}

func traverseMatrix(matrix *[][]string, prev Point, vector Point) int {
	// printMatrix(matrix, prev, vector)

	p := Point{prev.x + vector.x, prev.y + vector.y}

	if p.y < 0 || p.y > len(*matrix)-1 || p.x < 0 || p.x > len((*matrix)[0])-1 {
		return 0
	}

	if (*matrix)[p.y][p.x] != "#" {
		Path = append(Path, p)

		if (*matrix)[p.y][p.x] == "." {
			(*matrix)[p.y][p.x] = getVectorChar(vector)
			return 1 + traverseMatrix(matrix, p, vector)
		}

		(*matrix)[p.y][p.x] = getVectorChar(vector)
		return traverseMatrix(matrix, p, vector)
	}

	return traverseMatrix(matrix, prev, NextVector[str(vector)])
}

func parseInput(input string) (*[][]string, Point) {
	var matrix [][]string
	var p Point

	for j, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		row := strings.Split(line, "")
		for i, val := range row {
			if val == "^" {
				p.x = i
				p.y = j
				row[i] = "X"
			}
		}
		matrix = append(matrix, row)
	}
	return &matrix, p
}

func getVectorChar(v Point) string {
	switch {
	case v.x == 0 && v.y == -1:
		return "^"
	case v.x == 1 && v.y == 0:
		return ">"
	case v.x == 0 && v.y == 1:
		return "v"
	case v.x == -1 && v.y == 0:
		return "<"
	}
	return ""
}

func str(p Point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func strVector(p, v Point) string {
	return fmt.Sprintf("%d,%d (%d,%d)", p.x, p.y, v.x, v.y)
}

// func printMatrix(matrix *[][]string, p Point, v Point) {
// 	for j, line := range *matrix {
// 		for i, val := range line {
// 			if Obstacles[str(Point{i, j})] == 1 {
// 				fmt.Print("O")
// 			} else if p.x == i && p.y == j {
// 				fmt.Print(getVectorChar(v))
// 			} else {
// 				fmt.Print(val)
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// }
