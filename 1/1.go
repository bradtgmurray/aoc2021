package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []int {
	file, err := os.Open("./1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		input = append(input, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func part1(input []int) {
	counter := 0
	for i, _ := range input {
		if i == 0 { continue }

		if input[i] > input[i - 1] {
			counter += 1
		}
	}

	fmt.Printf("part1 %v\n", counter)
}

func part2(input []int) {
	counter := 0
	for i, _ := range input {
		if i < 3 { continue }

		if input[i] + input[i - 1] + input[i - 2] > input[i - 1] + input[i - 2] + input[i - 3] {
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
