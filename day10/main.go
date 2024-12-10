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

	matrix, trailheads, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1, resPart2 := solve(matrix, trailheads)

	time := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, time)
	fmt.Printf("Part 2: %v in %s\n", resPart2, time)
}

func solve(matrix *[][]int, trailheads *[]Point) (int, int) {
	sumP1 := 0
	sumP2 := 0

	for _, start := range *trailheads {
		trailends := map[string]int{}
		x := start.x
		y := start.y
		scoreP1, scoreP2 := checkNextStep(matrix, x, y, 0, 0, &trailends)
		sumP1 += scoreP1
		sumP2 += scoreP2
	}

	return sumP1, sumP2
}

func checkNextStep(matrix *[][]int, prevX int, prevY int, scoreP1 int, scoreP2 int, trailends *map[string]int) (int, int) {
	sizeX := len((*matrix)[0])
	sizeY := len(*matrix)
	prevValue := (*matrix)[prevY][prevX]

	for _, vector := range Vectors {
		x := prevX + vector.x
		y := prevY + vector.y
		if x >= 0 && y >= 0 && x < sizeX && y < sizeY {
			value := (*matrix)[y][x]

			if value == prevValue+1 {
				if value == 9 {
					scoreP2++
					if (*trailends)[str(Point{x, y})] != 1 {
						scoreP1++
						(*trailends)[str(Point{x, y})] = 1
					}
				} else {
					scoreP1, scoreP2 = checkNextStep(matrix, x, y, scoreP1, scoreP2, trailends)
				}
			}
		}
	}

	return scoreP1, scoreP2
}

func str(p Point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func parseInput(input string) (*[][]int, *[]Point, error) {
	matrix := [][]int{}
	trailheads := []Point{}

	for j, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var arr []int

		for i, val := range strings.Split(line, "") {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, nil, err
			}

			arr = append(arr, num)
			if num == 0 {
				trailheads = append(trailheads, Point{i, j})
			}
		}

		matrix = append(matrix, arr)
	}

	return &matrix, &trailheads, nil
}
