package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	inputs, outputs, outputReversed, outputLength, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(inputs, outputs, outputLength)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(outputs, outputReversed, outputLength)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(outputs map[string][]string, outputReversed map[string]string, outputLength int) string {
	var wiresToSwap = make([]string, 0)

	var tempsum, carry, tempcarry0, tempcarry1 string
	for i := range outputLength - 1 {
		xKey := getKey('x', i)
		yKey := getKey('y', i)
		zKey := getKey('z', i)

		if i == 0 {
			sum := getWire(outputReversed, "XOR", xKey, yKey)
			carry = getWire(outputReversed, "AND", xKey, yKey)

			if sum != zKey {
				wiresToSwap = append(wiresToSwap, zKey)
				wiresToSwap = append(wiresToSwap, sum)
				swapWires(outputReversed, zKey, sum)
			}
			continue
		}

		tempsum = getWire(outputReversed, "XOR", xKey, yKey)
		sum := getWire(outputReversed, "XOR", carry, tempsum)

		if sum == "" {
			rightSumOperation := outputs[zKey]
			key1 := rightSumOperation[0]
			key2 := rightSumOperation[2]

			correctTempSum := key1
			if key2 == carry {
				correctTempSum = key2
			}

			if correctTempSum == key2 {
				wiresToSwap = append(wiresToSwap, key1)
				wiresToSwap = append(wiresToSwap, tempsum)
				swapWires(outputReversed, key1, tempsum)
				tempsum = key1
			} else if correctTempSum == key1 {
				wiresToSwap = append(wiresToSwap, key2)
				wiresToSwap = append(wiresToSwap, tempsum)
				swapWires(outputReversed, key2, tempsum)
				tempsum = key2
			}
		} else if sum != zKey {
			wiresToSwap = append(wiresToSwap, zKey)
			wiresToSwap = append(wiresToSwap, sum)
			swapWires(outputReversed, zKey, sum)
		}

		tempcarry0 = getWire(outputReversed, "AND", xKey, yKey)
		tempcarry1 = getWire(outputReversed, "AND", carry, tempsum)
		carry = getWire(outputReversed, "OR", tempcarry0, tempcarry1)
	}

	slices.Sort(wiresToSwap)
	return strings.Join(wiresToSwap, ",")
}

func getWire(outputs map[string]string, operator string, x string, y string) string {
	key1 := fmt.Sprintf("%s %s %s", x, operator, y)
	key2 := fmt.Sprintf("%s %s %s", y, operator, x)

	if val, ok := outputs[key1]; ok {
		return val
	}

	if val, ok := outputs[key2]; ok {
		return val
	}

	return ""
}

func swapWires(outputs map[string]string, key1 string, key2 string) {
	for key, value := range outputs {
		if value == key1 {
			outputs[key] = key2
		}

		if value == key2 {
			outputs[key] = key1
		}
	}
}

func getKey(c rune, i int) string {
	var key string
	if i <= 9 {
		key = string(c) + "0" + strconv.Itoa(i)
	} else {
		key = string(c) + strconv.Itoa(i)
	}
	return key
}

func part1(inputs map[string]int, outputs map[string][]string, outputLength int) int64 {
	res := make([]int, outputLength)

	for key := range outputs {
		if key[0] != 'z' {
			continue
		}

		output := getValue(inputs, outputs, key)

		num, _ := strconv.Atoi(key[1:])
		res[outputLength-num-1] = output
	}

	return convertBinToInt(res)
}

func getValue(inputs map[string]int, outputs map[string][]string, key string) int {
	if val, ok := inputs[key]; ok {
		return val
	}

	arr := outputs[key]
	a := getValue(inputs, outputs, arr[0])
	b := getValue(inputs, outputs, arr[2])

	var res int
	switch arr[1] {
	case "XOR":
		res = a ^ b
	case "AND":
		if a == 1 && b == 1 {
			res = 1
		} else {
			res = 0
		}
	case "OR":
		if a == 0 && b == 0 {
			res = 0
		} else {
			res = 1
		}
	}

	inputs[key] = res
	return res
}

func convertBinToInt(arr []int) int64 {
	str := ""
	for _, num := range arr {
		str += strconv.Itoa(num)
	}

	converted, _ := strconv.ParseInt(str, 2, 64)

	return converted
}

func parseInput(input string) (map[string]int, map[string][]string, map[string]string, int, error) {
	inputs := make(map[string]int)
	outputs := make(map[string][]string)
	outputsReversed := make(map[string]string)
	outputLength := 0

	isParsingWires := false
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if len(line) == 0 {
			isParsingWires = true
			continue
		}

		if isParsingWires {
			splitLine := strings.Split(line, " -> ")
			outputs[splitLine[1]] = strings.Split(splitLine[0], " ")
			outputsReversed[splitLine[0]] = splitLine[1]

			if splitLine[1][0] == 'z' {
				outputLength++
			}
		} else {
			splitLine := strings.Split(line, ": ")

			num, err := strconv.Atoi(splitLine[1])
			if err != nil {
				return nil, nil, nil, 0, err
			}

			inputs[splitLine[0]] = num

		}
	}

	return inputs, outputs, outputsReversed, outputLength, nil
}
