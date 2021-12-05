package main

import (
	"fmt"
	"github.com/bradtgmurray/aoc2021/common"
	"regexp"
	"strconv"
	"strings"
)

type BingoBoard struct {
	cells [5][5]string
}

func newBoard(lines []string) BingoBoard {
	re := regexp.MustCompile("[ ]+")

	bb := BingoBoard { [5][5]string{} }
	for i, line := range lines {
		cells := re.Split(strings.TrimSpace(line), -1)
		for j, _ := range bb.cells[i] {
			bb.cells[i][j] = cells[j]
		}
	}
	return bb
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func hasWon(bb BingoBoard, numbers []string) bool {
	// Check rows
	for i := 0; i < 5; i++ {
		match := true
		for j := 0; j < 5; j++ {
			if !contains(numbers, bb.cells[i][j]) {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	// Check cols
	for i := 0; i < 5; i++ {
		match := true
		for j := 0; j < 5; j++ {
			if !contains(numbers, bb.cells[j][i]) {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}

func winningScore(bb BingoBoard, numbers []string) int {
	score := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !contains(numbers, bb.cells[i][j]) {
				number, _ := strconv.Atoi(bb.cells[i][j])
				score += number
			}
		}
	}

	return score
}

func Part1(numbers []string, boards []BingoBoard) int {
	for i, winningNumberString := range numbers {
		for _, bb := range boards {
			if hasWon(bb, numbers[:i+1]) {
				score := winningScore(bb, numbers[:i+1])
				winningNumber, _ := strconv.Atoi(winningNumberString)
				return score * winningNumber
			}
		}
	}

	return 0
}

func Part2(numbers []string, boards []BingoBoard) int {
	roundWonIn := make([]int, len(boards))

	for i, _ := range numbers {
		for j, bb := range boards {
			if roundWonIn[j] == 0 && hasWon(bb, numbers[:i+1]) {
				roundWonIn[j] = i
			}
		}
	}

	lastWinningRound := 0
	lastBoardToWin := 0
	for i, v := range roundWonIn {
		if v > lastWinningRound {
			lastWinningRound = v
			lastBoardToWin = i
		}
	}

	bb := boards[lastBoardToWin]
	score := winningScore(bb, numbers[:lastWinningRound+1])
	winningNumber, _ := strconv.Atoi(numbers[lastWinningRound])
	return score * winningNumber
}

func ParseInput(filename string) ([]string, []BingoBoard) {
	lines := common.ReadLines(filename)

	numbers := strings.Split(lines[0], ",")
	boards := []BingoBoard{}

	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, newBoard(lines[i:i+5]))
	}

	return numbers, boards
}

func main() {
	numbers, boards := ParseInput("./4/input.txt")

	fmt.Printf("part1 %v\n", Part1(numbers, boards))
	fmt.Printf("part2 %v\n", Part2(numbers, boards))
}
