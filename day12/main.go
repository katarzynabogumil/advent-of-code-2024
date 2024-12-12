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

var Checked = map[Point]bool{}

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

	matrix, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1, resPart2 := solve(matrix)
	time := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, time)
	fmt.Printf("Part 2: %v in %s\n", resPart2, time)
}

func solve(matrix *[][]string) (int, int) {
	sumP1 := 0
	sumP2 := 0

	for y, line := range *matrix {
		for x := range line {
			if !Checked[Point{x, y}] {
				points := []Point{}
				area, perimeter := checkArea(matrix, x, y, 0, 0, &points)
				sumP1 += area * perimeter

				sides := calculateSides(&points)
				sumP2 += area * sides
			}
		}
	}

	return sumP1, sumP2
}

func calculateSides(shape *[]Point) int {
	corners := map[Point]bool{}
	duplicateCorners := map[Point]bool{}

	for _, point := range *shape {
		for _, corner := range getCorners(point) {
			values := getValuesAroundCorner(corner, shape)

			switch {
			case values[0] && !values[1] && !values[2] && !values[3], // ABBB
				!values[0] && values[1] && !values[2] && !values[3], // BABB
				!values[0] && !values[1] && values[2] && !values[3], // BBAB
				!values[0] && !values[1] && !values[2] && values[3], // BBBA

				!values[0] && values[1] && values[2] && values[3], // BAAA
				values[0] && !values[1] && values[2] && values[3], // ABAA
				values[0] && values[1] && !values[2] && values[3], // AABA
				values[0] && values[1] && values[2] && !values[3]: // AAAB

				corners[corner] = true

			case values[0] && !values[1] && !values[2] && values[3], // ABBA
				!values[0] && values[1] && values[2] && !values[3]: // BAAB

				corners[corner] = true
				duplicateCorners[corner] = true
			}
		}
	}

	return len(corners) + len(duplicateCorners)
}

func getCorners(point Point) []Point {
	corners := []Point{}
	for _, vector := range []Point{
		{0, 0},
		{-1, 0},
		{-1, -1},
		{0, -1},
	} {
		corners = append(corners, Point{point.x + vector.x, point.y + vector.y})
	}
	return corners
}

func getValuesAroundCorner(corner Point, shape *[]Point) []bool {
	values := []bool{}
	for _, vector := range []Point{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
	} {
		value := slices.Contains(*shape, Point{corner.x + vector.x, corner.y + vector.y})
		values = append(values, value)
	}
	return values
}

func checkArea(matrix *[][]string, prevX int, prevY int, area int, perimeter int, points *[]Point) (int, int) {
	sizeX := len((*matrix)[0])
	sizeY := len(*matrix)
	prevValue := (*matrix)[prevY][prevX]

	Checked[Point{prevX, prevY}] = true
	*points = append(*points, Point{prevX, prevY})
	area += 1
	perimeter += 4

	for _, vector := range Vectors {
		x := prevX + vector.x
		y := prevY + vector.y
		if x >= 0 && y >= 0 && x < sizeX && y < sizeY {
			value := (*matrix)[y][x]

			if value == prevValue {
				perimeter -= 1

				if !Checked[Point{x, y}] {
					area, perimeter = checkArea(matrix, x, y, area, perimeter, points)
				}
			}
		}
	}

	return area, perimeter
}

func parseInput(input string) (*[][]string, error) {
	matrix := [][]string{}
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		arr := strings.Split(line, "")
		matrix = append(matrix, arr)
	}

	return &matrix, nil
}
