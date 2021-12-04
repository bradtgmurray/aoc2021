package main

import (
	"fmt"
	"github.com/bradtgmurray/aoc2021/common"
	"strconv"
)

func readInput() []int {
	strLines := common.ReadLines("./1/input.txt")

	input := make([]int, len(strLines))
	for i, v := range strLines {
		number, _ := strconv.Atoi(v)
		input[i] = number
	}
	return input
}

func part1(input []int) {
	counter := 0
	for i, _ := range input {
		if i == 0 {
			continue
		}

		if input[i] > input[i-1] {
			counter += 1
		}
	}

	fmt.Printf("part1 %v\n", counter)
}

func part2(input []int) {
	counter := 0
	for i, _ := range input {
		if i < 3 {
			continue
		}

		if input[i]+input[i-1]+input[i-2] > input[i-1]+input[i-2]+input[i-3] {
			counter += 1
		}
	}

	fmt.Printf("part2 %v\n", counter)
}

func main() {
	input := readInput()

	part1(input)
	part2(input)
}
