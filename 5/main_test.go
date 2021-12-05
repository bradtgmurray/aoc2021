package main

import "testing"

func TestPart1Example(t *testing.T) {
	lines := ParseLines([]string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	})

	part1 := Part1(lines, 10)
	if part1 != 5 {
		t.Errorf("It failed %v", part1)
	}
}

func TestPart2Example(t *testing.T) {
	lines := ParseLines([]string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	})

	part2 := Part2(lines, 10)
	if part2 != 12 {
		t.Errorf("It failed %v", part2)
	}
}