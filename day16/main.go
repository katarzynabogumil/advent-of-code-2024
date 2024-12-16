package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Point struct {
	x         int
	y         int
	wall      bool
	end       bool
	start     bool
	score     int
	direction Vector
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
	startTime := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	matrix, start := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	score, tilesCount := solve(matrix, start)
	time := time.Since(startTime)
	fmt.Printf("Part 1: %v in %s\n", score, time)
	fmt.Printf("Part 2: %v in %s\n", tilesCount, time)
}

func solve(matrix [][]Point, start Point) (int, int) {
	q := []Point{start}

	start.score = 0
	visited := map[string]int{
		str(start): start.score,
	}

	smallestScore := len(matrix[0]) * len(matrix) * 1000
	bestPaths := []string{str(start)}

	for len(q) > 0 {
		point := q[0]
		q = q[1:]

		if point.end {
			if point.score < smallestScore {
				bestPaths = []string{str(start)}
				smallestScore = point.score
			}

			if point.score == smallestScore {
				currPoint := point
				for !currPoint.start {
					if !slices.Contains(bestPaths, str(currPoint)) {
						bestPaths = append([]string{str(currPoint)}, bestPaths...)
					}
					currPoint = *currPoint.prev
				}
			}
		}

		for _, neighbour := range getPointNeighbours(matrix, point) {
			neighbour.prev = &point
			neighbour.direction = Vector{neighbour.x - point.x, neighbour.y - point.y}
			neighbour.score = point.score + 1
			if neighbour.direction != point.direction {
				neighbour.score += 1000
			}

			if val, ok := visited[str(neighbour)]; !ok || neighbour.score <= val+1000 {
				visited[str(neighbour)] = neighbour.score
				q = append(q, neighbour)
			}
		}
	}

	return smallestScore, len(bestPaths)
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
			if point.direction == vector && !newPoint.wall {
				neighbours = append(neighbours, matrix[y][x])
			} else if point.direction.x != vector.x && point.direction.y != vector.y && !newPoint.wall {
				neighbours = append(neighbours, matrix[y][x])
			}
		}
	}

	return neighbours
}

func str(p Point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func parseInput(input string) ([][]Point, Point) {
	matrix := [][]Point{}
	var start Point

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for y, inputLine := range lines {
		line := strings.Split(inputLine, "")
		arr := make([]Point, 0)
		for x, value := range line {
			isWall := false
			if value == "#" {
				isWall = true
			}

			newPoint := Point{
				x:    x,
				y:    y,
				wall: isWall,
			}

			if value == "S" {
				newPoint.direction = Vector{-1, 0}
				newPoint.start = true
				start = newPoint
			} else if value == "E" {
				newPoint.end = true
			}

			arr = append(arr, newPoint)
		}

		matrix = append(matrix, arr)
	}

	return matrix, start
}

// func printMatrix(matrix [][]Point, path []string) {
// 	for _, line := range matrix {
// 		for _, p := range line {
// 			if slices.Contains(path, str(p)) {
// 				fmt.Print("O")
// 			} else if p.wall {
// 				fmt.Print("#")
// 			} else {
// 				fmt.Print(".")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// }
