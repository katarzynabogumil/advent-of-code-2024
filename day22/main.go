package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Num struct {
	price, diff string
}

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	numbers, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(numbers, 2000)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(numbers, 2000)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(numbers []int, limit int) int {
	sequences := make(map[int][]Num)

	for i, num := range numbers {
		sequences[i] = make([]Num, 0)
		prev := num % 10

		for range limit - 1 {
			num = generateNum(num)
			price := num % 10
			diff := price - prev

			sequences[i] = append(sequences[i], Num{strconv.Itoa(price), strconv.Itoa(diff)})

			prev = price
		}
	}

	memo := make(map[string]map[int]string)
	for j, arr := range sequences {
		for i := 0; i < len(arr)-3; i++ {
			key := strings.Join([]string{arr[i].diff, arr[i+1].diff, arr[i+2].diff, arr[i+3].diff}, ",")

			if memo[key] == nil {
				memo[key] = make(map[int]string)
			}

			if _, ok := memo[key][j]; !ok {
				memo[key][j] = arr[i+3].price
			}
		}
	}

	return findMax(memo)
}

func findMax(memo map[string]map[int]string) int {
	max := 0
	for _, value := range memo {
		sum := 0

		for _, price := range value {
			num, _ := strconv.Atoi(price)
			sum += num
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

func part1(numbers []int, limit int) int {
	sum := 0

	for _, num := range numbers {
		for range limit {
			num = generateNum(num)
		}
		sum += num
	}

	return sum
}

func generateNum(num int) int {
	num1 := (num ^ (num * 64)) % 16777216
	num2 := (num1 ^ num1/32) % 16777216
	num3 := (num2 ^ (num2 * 2048)) % 16777216
	return num3
}

func parseInput(input string) ([]int, error) {
	arr := make([]int, 0)

	for _, val := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		arr = append(arr, num)
	}

	return arr, nil
}
