package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func part_one() int {
	return -1
}

func part_two() int {
	return -1
}

func main() {
	start := time.Now()
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}

	part_one_ans := part_one()
	part_two_ans := part_two()

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	end := time.Now()
	fmt.Printf("Time elapsed: %d milliseconds\n", end.UnixMilli()-start.UnixMilli())
}
