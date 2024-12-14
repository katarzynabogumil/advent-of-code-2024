package main

import (
	"fmt"
	"os"
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

	matrix, antennas := parseInput(string(input))

	resPart1 := part1(matrix, antennas)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(matrix, antennas)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(matrix [][]string, antennas map[string][]Point) int {
	counter := 0
	antidotes := map[string]int{}
	sizeX := len(matrix[0])
	sizeY := len(matrix)

	for _, arr := range antennas {
		if len(arr) == 0 {
			continue
		}

		if antidotes[str(arr[0])] != 1 {
			antidotes[str(arr[0])] = 1
			counter++
		}

		for i := 0; i < len(arr)-1; i++ {
			for j := 1; j < len(arr); j++ {
				if antidotes[str(arr[j])] != 1 {
					antidotes[str(arr[j])] = 1
					counter++
				}

				p1 := arr[i]
				p2 := arr[j]
				diffY := p2.y - p1.y
				diffX := p2.x - p1.x

				for _, p := range []Point{p1, p2} {
					newP := Point{p.x + diffX, p.y + diffY}

					for newP.x >= 0 && newP.x < sizeX && newP.y >= 0 && newP.y < sizeY && (diffX != 0 || diffY != 0) {
						if newP != p1 && newP != p2 &&
							antidotes[str(newP)] != 1 {
							antidotes[str(newP)] = 1
							counter++

							// matrix[newP.y][newP.x] = "#" // debug
						}
						newP = Point{newP.x + diffX, newP.y + diffY}
					}

					newP = Point{p.x - diffX, p.y - diffY}
					for newP.x >= 0 && newP.x < sizeX && newP.y >= 0 && newP.y < sizeY && (diffX != 0 || diffY != 0) {
						if newP != p1 && newP != p2 &&
							antidotes[str(newP)] != 1 {
							antidotes[str(newP)] = 1
							counter++

							// matrix[newP.y][newP.x] = "#" // debug
						}
						newP = Point{newP.x - diffX, newP.y - diffY}
					}
				}
			}
		}
	}

	// printMatrix(matrix)
	return counter
}

func part1(matrix [][]string, antennas map[string][]Point) int {
	counter := 0
	antidotes := map[string]int{}
	sizeX := len(matrix[0])
	sizeY := len(matrix)

	for _, arr := range antennas {
		if len(arr) == 0 {
			continue
		}

		for i := 0; i < len(arr); i++ {
			for j := 1; j < len(arr); j++ {
				p1 := arr[i]
				p2 := arr[j]
				diffY := p2.y - p1.y
				diffX := p2.x - p1.x

				for _, p := range []Point{p1, p2} {
					newP1 := Point{p.x + diffX, p.y + diffY}
					newP2 := Point{p.x - diffX, p.y - diffY}

					for _, newP := range []Point{newP1, newP2} {
						if newP != p1 && newP != p2 && newP.x >= 0 &&
							newP.x < sizeX && newP.y >= 0 && newP.y < sizeY &&
							antidotes[str(newP)] != 1 {
							antidotes[str(newP)] = 1
							counter++
						}
					}
				}
			}
		}
	}

	return counter
}

func str(p Point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func parseInput(input string) ([][]string, map[string][]Point) {
	var matrix [][]string
	antennas := map[string][]Point{}

	for j, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		row := strings.Split(line, "")
		for i, val := range row {
			if val != "." {
				antennas[val] = append(antennas[val], Point{i, j})
			}
		}
		matrix = append(matrix, row)
	}
	return matrix, antennas
}

// func printMatrix(matrix [][]string) {
// 	for _, line := range matrix {
// 		for _, val := range line {
// 			fmt.Print(val)
// 		}
// 		fmt.Print("\n")
// 	}
// 	fmt.Print("\n")
// }
