package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x         int
	y         int
	corrupted bool
	end       bool
	distance  int
	prev      *Point
}

type Vector struct {
	x int
	y int
}

var (
	Vectors = []Vector{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	size := 71
	limit := 1024
	matrix, nextPoints, err := parseInput(string(input), size, limit)
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(matrix)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(matrix, nextPoints)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(matrix [][]Point, nextPoints []Vector) string {
	for _, point := range nextPoints {
		matrix[point.y][point.x].corrupted = true
		if part1(matrix) == 0 {
			return fmt.Sprintf("%d,%d", point.x, point.y)
		}
	}
	return ""
}

func part1(matrix [][]Point) int {
	start := matrix[0][0]
	start.distance = 0

	q := []Point{start}
	visited := map[string]int{
		str(start): start.distance,
	}

	for len(q) > 0 {
		point := q[0]
		q = q[1:]

		if point.end {
			return point.distance
		}

		for _, neighbour := range getPointNeighbours(matrix, point) {
			neighbour.prev = &point
			neighbour.distance = point.distance + 1

			if val, ok := visited[str(neighbour)]; !ok || neighbour.distance < val {
				visited[str(neighbour)] = neighbour.distance
				q = append(q, neighbour)
			}
		}
	}

	return 0
}

func getPointNeighbours(matrix [][]Point, point Point) []Point {
	sizeX := len(matrix[0])
	sizeY := len(matrix)
	neighbours := make([]Point, 0)

	for _, vector := range Vectors {
		x := point.x + vector.x
		y := point.y + vector.y
		if x >= 0 && y >= 0 && x < sizeX && y < sizeY {
			newPoint := matrix[y][x]
			if !newPoint.corrupted {
				neighbours = append(neighbours, matrix[y][x])
			}
		}
	}

	return neighbours
}

func str(p Point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func parseInput(input string, size int, limit int) ([][]Point, []Vector, error) {
	matrix := [][]Point{}
	corrupted := make(map[Vector]bool)
	nextBytes := make([]Vector, 0)

	for i, inputLine := range strings.Split(strings.TrimSpace(string(input)), "\n") {

		nums := strings.Split(inputLine, ",")

		x, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, nil, err
		}

		y, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, err
		}

		if i < limit {
			corrupted[Vector{x, y}] = true
		} else {
			nextBytes = append(nextBytes, Vector{x, y})
		}
	}

	for y := range size {
		arr := make([]Point, 0)
		for x := range size {
			newPoint := Point{
				x:         x,
				y:         y,
				corrupted: corrupted[Vector{x, y}],
			}

			if x == size-1 && y == size-1 {
				newPoint.end = true
			}

			arr = append(arr, newPoint)
		}

		matrix = append(matrix, arr)
	}

	return matrix, nextBytes, nil
}

// func printMatrix(matrix [][]Point, path []string) {
// 	for _, line := range matrix {
// 		for _, p := range line {
// 			if slices.Contains(path, str(p)) {
// 				fmt.Print("O")
// 			} else if p.corrupted {
// 				fmt.Print("#")
// 			} else {
// 				fmt.Print(".")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// }
