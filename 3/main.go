package main

import (
	"fmt"
	"github.com/bradtgmurray/aoc2021/common"
	"strconv"
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

		if c > len(lines)/2 {
			gamma++
		} else {
			epsilon++
		}
	}

	return gamma * epsilon
}

func getOxygen(lines []string) string {
	remainingLines := lines[:]

	for i := 0; i < len(lines[0]); i++ {
		count := 0
		for _, line := range remainingLines {
			if line[i] == '1' {
				count++
			}
		}

		nextRemainingLines := []string{}
		for _, line := range remainingLines {
			if count * 2 >= len(remainingLines) {
				if line[i] == '1' {
					nextRemainingLines = append(nextRemainingLines, line)
				}
		    } else {
				if line[i] == '0' {
					nextRemainingLines = append(nextRemainingLines, line)
				}
			}
		}

		if len(nextRemainingLines) == 1 {
			return nextRemainingLines[0]
		}

		remainingLines = nextRemainingLines
	}

	return ""
}

func getCO2(lines []string) string {
	remainingLines := lines[:]

	for i := 0; i < len(lines[0]); i++ {
		count := 0
		for _, line := range remainingLines {
			if line[i] == '1' {
				count++
			}
		}

		nextRemainingLines := []string{}
		for _, line := range remainingLines {
			if count * 2 >= len(remainingLines) {
				if line[i] == '0' {
					nextRemainingLines = append(nextRemainingLines, line)
				}
			} else {
				if line[i] == '1' {
					nextRemainingLines = append(nextRemainingLines, line)
				}
			}
		}

		if len(nextRemainingLines) == 1 {
			return nextRemainingLines[0]
		}

		remainingLines = nextRemainingLines
	}

	return ""
}

func baseTwoStringToInt(s string) int64 {
	r, _ := strconv.ParseInt(s, 2, 32)
	return r
}

func Part2(lines []string) int64 {
	oxygen := baseTwoStringToInt(getOxygen(lines))
	co2 := baseTwoStringToInt(getCO2(lines))

	return oxygen * co2
}

func main() {
	commands := common.ReadLines("./3/input.txt")

	fmt.Printf("part1 %v\n", Part1(commands))
	fmt.Printf("part2 %v\n", Part2(commands))
}
