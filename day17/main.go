package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var Registers = make(map[int]int)

var Pointer int

var Instructions = map[int]func(o int) string{
	0: func(o int) string {
		operand := getComboOperand(o)
		value := Registers[4] / int(math.Pow(2, float64(operand)))
		Registers[4] = value
		Pointer += 2
		return ""
	},
	1: func(o int) string {
		value := Registers[5] ^ o
		Registers[5] = value
		Pointer += 2
		return ""
	},
	2: func(o int) string {
		value := getComboOperand(o) % 8
		Registers[5] = value
		Pointer += 2
		return ""
	},
	3: func(o int) string {
		if Registers[4] == 0 {
			Pointer += 2
			return ""
		}
		Pointer = o
		return ""
	},
	4: func(o int) string {
		value := Registers[5] ^ Registers[6]
		Registers[5] = value
		Pointer += 2
		return ""
	},
	5: func(o int) string {
		value := getComboOperand(o) % 8
		Pointer += 2
		return strings.Join(strings.Split(strconv.Itoa(value), ""), ",")
	},
	6: func(o int) string {
		operand := getComboOperand(o)
		value := Registers[4] / int(math.Pow(2, float64(operand)))
		Registers[5] = value
		Pointer += 2
		return ""
	},
	7: func(o int) string {
		operand := getComboOperand(o)
		value := Registers[4] / int(math.Pow(2, float64(operand)))
		Registers[6] = value
		Pointer += 2
		return ""
	},
}

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	program, registers, err := parseInput(string(input))
	if err != nil {
		fmt.Println(fmt.Errorf("can't parse file: %w", err))
	}

	resPart1 := part1(program, registers)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %s in %s\n", resPart1, timePart1)

	resPart2 := part2(program, registers)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

// improved brute force after noticing that first Register A values affect last program output values
func part2(program []int, registers map[int]int) int {
	programStr := make([]string, 0)

	for _, val := range program {
		programStr = append(programStr, strconv.Itoa(val))
	}

	goal := strings.Join(programStr, ",")

	length := 0
	result := ""
	registerA := 0
	registerB := registers[5]
	registerC := registers[6]

	// estimated initial power to calculate the increment value, f.e.:
	// if the result is not close to the goal,
	// increment by 10^(<length of the program instructions> - 3)
	initialPower := len(programStr) - 3

	for length <= len(programStr) {
		result, length = runProgram(program, map[int]int{
			4: registerA,
			5: registerB,
			6: registerC,
		})

		if goal == result {
			return registerA
		}

		contains := false
		if length == len(programStr) {
			for i := len(programStr); i > 0; i-- {
				if strings.Contains(result+"E", strings.Join(programStr[len(programStr)-i:], ",")+"E") {
					// increment by a value proportional to how close we are to the goal result
					registerA += int(math.Pow(10, getPower(i, initialPower)))
					contains = true
					break
				}
			}
		}

		if !contains || length != len(programStr) {
			// increment by a bigger value if not close to the goal result
			registerA += int((math.Pow(10, getPower(0, initialPower))))
		}
	}

	return 0
}

func getPower(i, initialPower int) float64 {
	if i < initialPower {
		return float64(initialPower - i)
	}
	return float64(0)
}

func part1(program []int, registers map[int]int) string {
	output, _ := runProgram(program, registers)
	return output
}

func runProgram(program []int, registers map[int]int) (string, int) {
	Registers = registers
	output := make([]string, 0)
	Pointer = 0

	for Pointer < len(program)-1 {
		opcode := program[Pointer]
		operand := program[Pointer+1]
		value := Instructions[opcode](operand)
		if len(value) != 0 {
			output = append(output, value)
		}
	}

	return strings.Join(output, ","), len(output)
}

func getComboOperand(code int) int {
	if code < 4 {
		return code
	}
	return Registers[code]
}

func parseInput(input string) ([]int, map[int]int, error) {
	program := make([]int, 0)
	registers := make(map[int]int)

	isReadingProgram := false
	for i, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if len(line) == 0 {
			isReadingProgram = true
			continue
		}
		if isReadingProgram {
			for _, val := range strings.Split(strings.Replace(line, "Program: ", "", 1), ",") {
				num, err := strconv.Atoi(val)
				if err != nil {
					return nil, nil, err
				}

				program = append(program, num)
			}
		} else {
			var numStr string
			switch i {
			case 0:
				numStr = strings.Replace(line, "Register A: ", "", 1)
			case 1:
				numStr = strings.Replace(line, "Register B: ", "", 1)
			case 2:
				numStr = strings.Replace(line, "Register C: ", "", 1)
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}

			registers[i+4] = num
		}
	}

	return program, registers, nil
}
