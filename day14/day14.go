package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func part_one(grid [][]string) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "O" {
				moveup := 0
				for i-moveup > 0 && grid[i-moveup-1][j] == "." {
					moveup += 1
				}

				if moveup > 0 {
					grid[i][j] = "."
					grid[i-moveup][j] = "O"
				}

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
	part_two_ans := 0

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
