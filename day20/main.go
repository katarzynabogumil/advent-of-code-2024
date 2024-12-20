package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Point struct {
	x     int
	y     int
	wall  bool
	end   bool
	start bool
	time  int
	prev  *Point
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

	path := findPath(matrix, start)

	resPart1 := solve(path, 2)
	timePart1 := time.Since(startTime)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := solve(path, 20)
	timePart2 := time.Since(startTime)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func solve(path []Point, limit int) int {
	counter := 0

	for i := 0; i < len(path); i++ {
		point1 := path[i]
		for j := i + 1; j < len(path); j++ {
			point2 := path[j]
			distance := getDistance(point1, point2)
			if distance > limit {
				continue
			} else if point2.time+distance+100 <= point1.time {
				counter++
			}
		}
	}

	return counter
}

func getDistance(p1 Point, p2 Point) int {
	return absInt(p1.x-p2.x) + absInt(p1.y-p2.y)
}

func findPath(matrix [][]Point, start Point) []Point {
	q := []Point{start}

	start.time = 0
	visited := map[string]int{
		str(start): start.time,
	}

	for len(q) > 0 {
		point := q[0]
		q = q[1:]

		if point.end {
			path := make([]Point, 0)
			currPoint := point
			for !currPoint.start {
				path = append(path, currPoint)
				currPoint = *currPoint.prev
			}
			path = append(path, start)
			return path
		}

		for _, neighbour := range getPointNeighbours(matrix, point) {
			neighbour.prev = &point
			neighbour.time = point.time + 1

			if val, ok := visited[str(neighbour)]; !ok || neighbour.time < val {
				// save time in the matrix
				matrix[neighbour.y][neighbour.x] = neighbour

				visited[str(neighbour)] = neighbour.time
				q = append(q, neighbour)
			}

		}
	}

	return nil
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
			if !newPoint.wall {
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

func absInt(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
