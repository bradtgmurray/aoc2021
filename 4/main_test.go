package main

import "testing"

func TestPart1Example(t *testing.T) {
	numbers, boards := ParseInput("./example_input.txt")

	part1 := Part1(numbers, boards)
	if part1 != 4512 {
		t.Errorf("It failed %v", part1)
	}
}

func TestPart2Example(t *testing.T) {
	numbers, boards := ParseInput("./example_input.txt")

	part2 := Part2(numbers, boards)
	if part2 != 1924 {
		t.Errorf("It failed %v", part2)
	}
}

