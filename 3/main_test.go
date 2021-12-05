package main

import "testing"

func TestPart1Simple(t *testing.T) {
	input := []string{
		"00100",
	}

	expected := 0b00100 * 0b11011

	part1 := Part1(input)
	if part1 != expected {
		t.Errorf("It failed %v != %v", part1, expected)
	}
}

func TestPart1Example(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	part1 := Part1(input)
	if part1 != 198 {
		t.Errorf("It failed %v", part1)
	}
}
