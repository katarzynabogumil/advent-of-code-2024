package main

import (
	"bufio"
	"fmt"
	"math"
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
	firstArr, secondArr, err := parseInput(scanner)
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(firstArr, secondArr)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(firstArr, secondArr)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(firstArr []int, secondArr []int) int {
	similarityMap := make(map[int]int, len(secondArr))
	for _, i := range secondArr {
		similarityMap[i] += 1
	}

	sum := 0
	for _, i := range firstArr {
		sum += i * similarityMap[i]
	}

	return sum
}

func part1(unsortedFirstArr []int, unsortedSecondArr []int) int {
	sum := 0

	sortedFirstArr := mergeSort(unsortedFirstArr)
	sortedSecondArr := mergeSort(unsortedSecondArr)

	for i, firstInt := range sortedFirstArr {
		sum += int(math.Abs(float64(firstInt) - float64(sortedSecondArr[i])))
	}

	return sum
}

func parseInput(scanner *bufio.Scanner) ([]int, []int, error) {
	var firstArr, secondArr []int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		firstInt, err := strconv.Atoi(line[0])
		if err != nil {
			return firstArr, secondArr, err
		}

		secondInt, err := strconv.Atoi(line[1])
		if err != nil {
			return firstArr, secondArr, err
		}

		firstArr = append(firstArr, firstInt)
		secondArr = append(secondArr, secondInt)
	}

	return firstArr, secondArr, nil
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	firstArr := mergeSort(arr[:len(arr)/2])
	secondArr := mergeSort(arr[len(arr)/2:])

	return merge(firstArr, secondArr)
}

func merge(a []int, b []int) []int {
	arr := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		arr = append(arr, a[i])
	}

	for ; j < len(b); j++ {
		arr = append(arr, b[j])
	}

	return arr
}
