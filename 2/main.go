package main

import (
	"fmt"
	"github.com/bradtgmurray/aoc2021/common"
	"strconv"
	"strings"
)

func part1(commands []string) {
	position := 0
	depth := 0

	for _, line := range commands {
		parts := strings.Split(line, " ")

		command := parts[0]
		number, _ := strconv.Atoi(parts[1])

		switch {
		case command == "forward":
			position += number
		case command == "up":
			depth -= number
		case command == "down":
			depth += number
		}
	}

	fmt.Printf("part1 %v\n", position * depth)
}

func part2(commands []string) {
	position := 0
	depth := 0
	aim := 0

	for _, line := range commands {
		parts := strings.Split(line, " ")

		command := parts[0]
		number, _ := strconv.Atoi(parts[1])

		switch {
		case command == "forward":
			position += number
			depth += aim * number
		case command == "up":
			aim -= number
		case command == "down":
			aim += number
		}
	}

	fmt.Printf("part2 %v\n", position * depth)
}


func main() {
	commands := common.ReadLines("./2/input.txt")

	part1(commands)
	part2(commands)
}
