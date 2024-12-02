package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	counterPart1, counterPart2 := 0, 0
	for scanner.Scan() {
		var arr []int
		line := strings.Split(scanner.Text(), " ")

		for _, num := range line {
			i, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println(err)
			}
			arr = append(arr, i)
		}

		if checkIfSafePart1(arr) {
			counterPart1 += 1
		}

		if checkIfSafePart2(arr, false) {
			counterPart2 += 1
		}
	}

	time := time.Since(start)
	fmt.Printf("Part 1: %v\nPart 2: %v in %s\n", counterPart1, counterPart2, time)
}

func checkIfSafePart2(arr []int, wasLevelRemoved bool) bool {
	isIncreasing := false
	isDecreasing := false

	if arr[1] > arr[0] {
		isIncreasing = true
	} else if arr[0] > arr[1] {
		isDecreasing = true
	} else if !wasLevelRemoved {
		return checkIfSafePart2(arr[1:], true)
	} else {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if isIncreasing && arr[i] <= arr[i-1] {
			return checkRemoving(arr, i, wasLevelRemoved)
		}

		if isDecreasing && arr[i] >= arr[i-1] {
			return checkRemoving(arr, i, wasLevelRemoved)
		}

		if absInt(arr[i]-arr[i-1]) > 3 {
			return checkRemoving(arr, i, wasLevelRemoved)
		}
	}

	return true
}

func checkRemoving(arr []int, i int, wasLevelRemoved bool) bool {
	if wasLevelRemoved {
		return false
	}

	if i == (len(arr) - 1) {
		return true
	}

	if checkIfSafePart2(arr[1:], true) {
		return true
	}

	if checkIfSafePart2(append(slices.Clone(arr[:i-1]), arr[i:]...), true) {
		return true
	}

	return checkIfSafePart2(append(slices.Clone(arr[:i]), arr[i+1:]...), true)
}

func checkIfSafePart1(arr []int) bool {
	isIncreasing := false
	isDecreasing := false

	if arr[1] > arr[0] {
		isIncreasing = true
	} else if arr[0] > arr[1] {
		isDecreasing = true
	} else {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if isIncreasing && arr[i] <= arr[i-1] {
			return false
		}

		if isDecreasing && arr[i] >= arr[i-1] {
			return false
		}

		if absInt(arr[i]-arr[i-1]) > 3 {
			return false
		}
	}

	return true
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
