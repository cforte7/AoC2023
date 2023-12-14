package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func part_one(grid [][]string) int {
	row_count := len(grid)
	col_count := len(grid[0])
	for col := col_count - 1; col >= 0; col-- {
		stop := 0
		for row := 0; row < row_count; row++ {
			if grid[row][col] == "O" {
				grid[row][col] = "."
				grid[stop][col] = "O"
				stop += 1
			} else if grid[row][col] == "#" {
				stop = row + 1
			}
		}
	}
	tot := 0
	grid_len := len(grid)
	for row_num, v := range grid {
		row_count := 0
		for _, c := range v {
			if c == "O" {
				row_count += 1
			}
		}
		tot += row_count * (grid_len - row_num)
	}
	return tot
}

func part_two(grid [][]string) int {
	return -1
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")

	as_rows := strings.Split(string(file), "\n")

	full_split := [][]string{}
	for _, v := range as_rows {
		as_rune := strings.Split(v, "")
		full_split = append(full_split, as_rune)
	}

	part_one_ans := part_one(full_split)
	part_two_ans := part_two(full_split)

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
