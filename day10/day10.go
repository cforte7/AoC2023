package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// map from char -> enter dir -> next dir (row, col)
var index_change = map[rune]map[rune][2]int{
	'|': {'B': {-1, 0}, 'T': {1, 0}},
	'-': {'L': {0, 1}, 'R': {0, -1}},
	'L': {'T': {0, 1}, 'R': {-1, 0}},
	'J': {'L': {-1, 0}, 'T': {0, -1}},
	'7': {'L': {1, 0}, 'B': {0, -1}},
	'F': {'R': {1, 0}, 'B': {0, 1}},
}

var dir_change = map[rune]map[rune]rune{
	'|': {'B': 'B', 'T': 'T'},
	'-': {'L': 'L', 'R': 'R'},
	'L': {'T': 'L', 'R': 'B'},
	'J': {'L': 'B', 'T': 'R'},
	'7': {'L': 'T', 'B': 'R'},
	'F': {'R': 'T', 'B': 'L'},
}

func part_one(data []string, row int, col int, entering_from rune) int {
	curr_char := data[row][col]
	mvmt := index_change[rune(curr_char)][entering_from]
	next_entering_dir, found := dir_change[rune(curr_char)][entering_from]
	if !found {
		panic("no next dir found")
	}
	next_row := row + mvmt[0]
	next_col := col + mvmt[1]
	if data[next_row][next_col] == 'S' {
		return 1
	}

	return 1 + part_one(data, row+mvmt[0], col+mvmt[1], next_entering_dir)
}

// func part_two() int {
// 	return -1
// }

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")

	as_rows := strings.Split(string(file), "\n")

	part_one_ans := (1 + part_one(as_rows, 75, 54, 'L')) / 2
	part_two_ans := 0
	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
