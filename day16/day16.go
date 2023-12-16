package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func backslash_switcher(dir [2]int) [2]int {
	return [2]int{dir[1], dir[0]}
}

func forwardslash_switcher(dir [2]int) [2]int {
	return [2]int{-1 * dir[1], -1 * dir[0]}
}

var seen_with_dir = map[[4]int]bool{}
var seen = map[[2]int]bool{}

func mover(loc [2]int, dir [2]int) [2]int {
	return [2]int{loc[0] + dir[0], loc[1] + dir[1]}
}

func part_one(data []string, loc [2]int, dir [2]int) {
	if loc[0] >= len(data) || loc[1] >= len(data[0]) || loc[0] < 0 || loc[1] < 0 || seen_with_dir[[4]int{loc[0], loc[1], dir[0], dir[1]}] {
		return
	}
	seen_with_dir[[4]int{loc[0], loc[1], dir[0], dir[1]}] = true
	seen[loc] = true

	curr_char := data[loc[0]][loc[1]]
	if curr_char == '.' {
		next_pos := mover(loc, dir)
		part_one(data, next_pos, dir)
	} else if curr_char == '\\' {
		new_dir := backslash_switcher(dir)
		next_pos := mover(loc, new_dir)
		part_one(data, next_pos, new_dir)
	} else if curr_char == '/' {
		new_dir := forwardslash_switcher(dir)
		next_pos := mover(loc, new_dir)
		part_one(data, next_pos, new_dir)
	} else if curr_char == '|' {
		if dir[0] == 0 {
			top_loc := [2]int{loc[0] - 1, loc[1]}
			bot_loc := [2]int{loc[0] + 1, loc[1]}
			part_one(data, top_loc, [2]int{-1, 0})
			part_one(data, bot_loc, [2]int{1, 0})
		} else {
			next_pos := mover(loc, dir)
			part_one(data, next_pos, dir)
		}
	} else if curr_char == '-' {
		if dir[1] == 0 {
			left_loc := [2]int{loc[0], loc[1] - 1}
			right_loc := [2]int{loc[0], loc[1] + 1}
			part_one(data, left_loc, [2]int{0, -1})
			part_one(data, right_loc, [2]int{0, 1})
		} else {
			next_pos := mover(loc, dir)
			part_one(data, next_pos, dir)
		}
	}
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")

	as_rows := strings.Split(string(file), "\n")

	part_one(as_rows, [2]int{0, 0}, [2]int{0, 1})
	fmt.Printf("Part one: %d\n", len(seen))
	seen = map[[2]int]bool{}
	seen_with_dir = map[[4]int]bool{}
	part_two_ans := 0

	for i := 0; i < len(as_rows); i++ {
		for j := 0; j <= len(as_rows[0]); j++ {
			if i == 0 || i == len(as_rows)-1 || j == 0 || j == len(as_rows[0])-1 {

				for rodir := range [3]int{0, 1, -1} {
					for coldiur := range [3]int{0, 1, -1} {
						part_one(as_rows, [2]int{i, j}, [2]int{rodir, coldiur})
						part_two_ans = max(part_two_ans, len(seen))
						seen = map[[2]int]bool{}
						seen_with_dir = map[[4]int]bool{}
					}
				}
			}
		}
	}

	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
