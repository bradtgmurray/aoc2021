package main

import (
	"fmt"
	"github.com/bradtgmurray/aoc2021/common"
	"regexp"
	"strconv"
	"strings"
)

type Line struct {
	startX int
	startY int
	endX   int
	endY   int
}

func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func Part1(lines []Line, gridSize int) int {
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	for _, l := range lines {
		if l.startX == l.endX {
			for i := min(l.startY, l.endY); i <= max(l.startY, l.endY); i++ {
				grid[l.startX][i] += 1
			}
		} else if l.startY == l.endY {
			for i := min(l.startX, l.endX); i <= max(l.startX, l.endX); i++ {
				grid[i][l.startY] += 1
			}
		}
	}

	counter := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] >= 2 {
				counter++
			}
		}
	}

	return counter
}

func Part2(lines []Line, gridSize int) int {
	grid := make([][]int, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}


	for _, l := range lines {
		lineLength := max(abs(l.endX - l.startX), abs(l.endY - l.startY))
		stepX := (l.endX - l.startX) / lineLength
		stepY := (l.endY - l.startY) / lineLength

		for i := 0; i < lineLength + 1; i++ {
			grid[l.startX + (stepX * i)][l.startY + (stepY * i)]++
		}
	}

	counter := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] >= 2 {
				counter++
			}
		}
	}

	return counter
}

func ParseLines(inputLines []string) []Line {
	re := regexp.MustCompile(" -> ")

	lines := make([]Line, len(inputLines))

	for i, l := range inputLines {
		ends := re.Split(l, -1)
		start := strings.Split(ends[0], ",")
		end := strings.Split(ends[1], ",")

		lines[i] = Line{toInt(start[0]), toInt(start[1]), toInt(end[0]), toInt(end[1])}
	}

	return lines
}

func main() {
	inputLines := common.ReadLines("./5/input.txt")
	lines := ParseLines(inputLines)

	fmt.Printf("part1 %v\n", Part1(lines, 1000))
	fmt.Printf("part2 %v\n", Part2(lines, 1000))
}
