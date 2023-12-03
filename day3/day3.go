package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part_one() int {
	return -1
}

func part_two() int {
	return -1
}

func coords_to_str(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func is_symb(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

func get_symbol_cords(s bufio.Scanner) map[string]bool {
	var symbol_coords = map[string]bool{}
	for y := 0; s.Scan(); y++ {
		for x, char := range s.Text() {
			if !unicode.IsDigit(char) && char != '.' {
				symbol_coords[coords_to_str(x, y)] = true
			}
		}
	}
	return symbol_coords
}

func touches_symbol(grid []string, x int, y int, max_x int, max_y int) bool {
	for _, x_diff := range [3]int{-1, 0, 1} {
		for _, y_diff := range [3]int{-1, 0, 1} {
			y_check := y + y_diff
			x_check := x + x_diff
			if 0 <= y_check && y_check <= max_y && 0 <= x_check && x_check <= max_x && is_symb(rune(grid[y+y_diff][x+x_diff])) {
				fmt.Printf("got: true for check_pos %d,%d from source %d,%d\n", x_check, y_check, x, y)
				return true
			}
		}
	}
	return false
}

func main() {
	// file, _ := os.Open("./input.txt")
	// defer file.Close()
	file2, _ := os.ReadFile("./input.txt")
	as_rows := strings.Split(string(file2), "\n")
	max_x := len(as_rows[0]) - 1
	max_y := len(as_rows) - 1
	var curr_num string
	var to_add bool
	total := 0
	for y, list := range as_rows {
		for x, char := range list {
			if unicode.IsDigit(char) {
				curr_num += string(char)
				to_add = to_add || touches_symbol(as_rows, x, y, max_x, max_y)
			} else {
				if to_add && len(curr_num) > 0 {
					as_num, _ := strconv.Atoi(curr_num)
					total += as_num
				}
				to_add = false
				curr_num = ""
			}
		}
	}
	print(total)

	// scanner := bufio.NewScanner(file)
	// symbol_cords := get_symbol_cords(*scanner)
}
