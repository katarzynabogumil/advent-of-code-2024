package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't open file"))
	}

	scanner := bufio.NewScanner(file)
	lines, err := parseInput(scanner)
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(lines)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(lines)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(lines [][]int) int {
	sum := 0
	for _, line := range lines {
		if checkValuePart2(line[0], line[1], line[2:]) {
			sum += line[0]
		}
	}
	return sum
}

func checkValuePart2(testValue int, val int, arr []int) bool {
	if len(arr) == 0 && testValue == val {
		return true
	}

	if len(arr) == 0 {
		return false
	}

	return checkValuePart2(testValue, val*arr[0], arr[1:]) ||
		checkValuePart2(testValue, val+arr[0], arr[1:]) ||
		checkValuePart2(testValue, combine(val, arr[0]), arr[1:])
}

func combine(a int, b int) int {
	i, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return i
}

func part1(lines [][]int) int {
	sum := 0
	for _, line := range lines {
		if checkValuePart1(line[0], line[1], line[2:]) {
			sum += line[0]
		}
	}
	return sum
}

func checkValuePart1(testValue int, val int, arr []int) bool {
	if len(arr) == 0 && testValue == val {
		return true
	}

	if len(arr) == 0 {
		return false
	}

	return checkValuePart1(testValue, val*arr[0], arr[1:]) || checkValuePart1(testValue, val+arr[0], arr[1:])
}

func parseInput(scanner *bufio.Scanner) ([][]int, error) {
	var arr [][]int

	for scanner.Scan() {
		var lineArr []int
		line := strings.Split(scanner.Text(), ": ")

		testValue, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		lineArr = append(lineArr, testValue)

		nums := strings.Split(line[1], " ")
		for _, n := range nums {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			lineArr = append(lineArr, num)
		}
		arr = append(arr, lineArr)
	}

	return arr, nil
}
