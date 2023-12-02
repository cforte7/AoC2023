package main

import (
	"bufio"
	"fmt"
	"os"
)

func part_one() int {
	return -1
}

func part_two() int {
	return -1
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}

	part_one_ans := part_one()
	part_two_ans := part_two()

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d", part_two_ans)
}
