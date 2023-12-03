package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func is_symb(char rune) bool {
	return !unicode.IsDigit(char) && char != '.'
}

func touches_symbol(grid []string, x int, y int, max_x int, max_y int) bool {
	for _, x_diff := range [3]int{-1, 0, 1} {
		for _, y_diff := range [3]int{-1, 0, 1} {
			y_check := y + y_diff
			x_check := x + x_diff
			if 0 <= y_check && y_check <= max_y && 0 <= x_check && x_check <= max_x && is_symb(rune(grid[y+y_diff][x+x_diff])) {
				return true
			}
		}
	}
	return false
}
func part_one(as_rows []string, max_x int, max_y int) int {
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
	return total
}

func get_number(line string, start int, linenum int) (int, string) {
	num := ""
	back_ind := start
	forward_ind := start

	for back_ind > 0 {
		back_ind -= 1
		if !unicode.IsDigit(rune(line[back_ind])) {
			break
		}
		num = fmt.Sprintf("%s%s", string(line[back_ind]), num)
	}
	num = fmt.Sprintf("%s%s", num, string(line[start]))
	for forward_ind < len(line)-1 {
		forward_ind += 1
		if !unicode.IsDigit(rune(line[forward_ind])) {
			break
		}
		num = fmt.Sprintf("%s%s", num, string(line[forward_ind]))
	}
	res, _ := strconv.Atoi(num)
	return res, fmt.Sprintf("%d,%d,%d", back_ind, forward_ind, linenum)
}

func touches_number(grid []string, x int, y int, max_x int, max_y int) []int {
	// key is {start},{end} index so we can avoid double counting the same number but inclue the same number in two places
	found_numbers := map[string]int{}

	for _, x_diff := range [3]int{-1, 0, 1} {
		for _, y_diff := range [3]int{-1, 0, 1} {
			y_check := y + y_diff
			x_check := x + x_diff
			if 0 <= y_check && y_check <= max_y && 0 <= x_check && x_check <= max_x && unicode.IsDigit(rune(grid[y_check][x_check])) {
				val, r := get_number(grid[y+y_diff], x_check, y_diff)
				_, found := found_numbers[r]
				if !found {
					found_numbers[r] = val
				}
				found_numbers[r] = val
			}
		}
	}
	vals := make([]int, 0, len(found_numbers))
	for _, v := range found_numbers {
		vals = append(vals, v)
	}
	return vals
}

func part_two(as_rows []string, max_x int, max_y int) int {
	total := 0
	for y, list := range as_rows {
		for x, char := range list {
			if char == '*' {
				res := touches_number(as_rows, x, y, max_x, max_y)
				if len(res) == 2 {
					total += res[0] * res[1]
				}
			}
		}
	}
	return total
}

func main() {

	file2, _ := os.ReadFile("./input.txt")
	as_rows := strings.Split(string(file2), "\n")
	max_x := len(as_rows[0]) - 1
	max_y := len(as_rows) - 1
	part_one_ans := part_one(as_rows, max_x, max_y)
	part_two_ans := part_two(as_rows, max_x, max_y)
	println(part_one_ans)
	println(part_two_ans)
}
