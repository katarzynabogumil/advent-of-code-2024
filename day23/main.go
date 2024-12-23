package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can't read file"))
	}

	network := parseInput(string(input))

	resPart1 := part1(network)
	timePart1 := time.Since(start)
	fmt.Printf("Part 1: %v in %s\n", resPart1, timePart1)

	resPart2 := part2(network)
	timePart2 := time.Since(start)
	fmt.Printf("Part 2: %v in %s\n", resPart2, timePart2)
}

func part2(network map[string][]string) string {
	groups := make([][]string, 0)

	for curr, connections := range network {
		for i := 0; i < len(connections); i++ {
			found := false
			for k, group := range groups {
				contains1 := slices.Contains(group, curr)
				contains2 := slices.Contains(group, connections[i])

				if contains1 || contains2 {
					for l, other := range group {
						if (other != curr && !slices.Contains(network[other], curr)) ||
							(other != connections[i] && !slices.Contains(network[other], connections[i])) {
							break
						}

						if l == len(group)-1 {
							found = true

							if !contains1 {
								groups[k] = append(group, curr)
							}

							if !contains2 {
								groups[k] = append(group, connections[i])
							}
						}
					}

				}
			}

			if !found {
				groups = append(groups, []string{curr, connections[i]})
			}
		}
	}

	maxLength := 0
	idx := 0
	for k, group := range groups {
		if len(group) > maxLength {
			maxLength = len(group)
			idx = k
		}
	}

	slices.Sort(groups[idx])
	return strings.Join(groups[idx], ",")
}

func part1(network map[string][]string) int {
	found := make(map[string]bool)

	for curr, connections := range network {
		for i := 0; i < len(connections); i++ {
			for j := 0; j < len(connections); j++ {
				if i == j {
					continue
				}

				if !slices.Contains(network[connections[i]], connections[j]) ||
					!slices.Contains(network[connections[j]], connections[i]) {
					continue
				}

				if curr[0] != 't' && connections[i][0] != 't' && connections[j][0] != 't' {
					continue
				}

				connection := make([]string, 0)
				connection = append(connection, curr)
				connection = append(connection, connections[i])
				connection = append(connection, connections[j])
				slices.Sort(connection)
				key := strings.Join(connection, ",")
				found[key] = true
			}
		}
	}

	return len(found)
}

func parseInput(input string) map[string][]string {
	network := make(map[string][]string)

	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		curruters := strings.Split(line, "-")
		if network[curruters[0]] == nil {
			network[curruters[0]] = make([]string, 0)
		}
		network[curruters[0]] = append(network[curruters[0]], curruters[1])

		if network[curruters[1]] == nil {
			network[curruters[1]] = make([]string, 0)
		}
		network[curruters[1]] = append(network[curruters[1]], curruters[0])
	}

	return network
}
