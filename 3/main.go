package main

import (
	"fmt"
	"github.com/bradtgmurray/aoc2021/common"
)

func Part1(lines []string) int {
	counts := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				counts[i]++
			}
		}
	}

	gamma := 0
	epsilon := 0

	for _, c := range counts {
		gamma = gamma * 2
		epsilon = epsilon * 2

		if c > len(lines) / 2 {
			gamma++
		} else {
			epsilon++
		}
	}

	return gamma * epsilon
}

func main() {
	commands := common.ReadLines("./3/input.txt")

	fmt.Printf("part1 %v\n", Part1(commands))
}
