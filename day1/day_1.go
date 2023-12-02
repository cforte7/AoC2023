package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_is_int(s string) (int64, error) {
	val, int_err := strconv.ParseInt(s, 10, 64)
	if int_err == nil {
		return val, nil
	}
	return -1, errors.New("not a digit")
}

var lookup = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func check_is_text(s string, index int) (int64, error) {
	// check if the index and proceeding characters match a
	// text number. Must check that you are not out of bounds
	// of the string
	str_len := len(s)
	for str_val, int_val := range lookup {
		if str_len-index >= len(str_val) && s[index:index+len(str_val)] == str_val {
			return int64(int_val), nil
		}
	}
	return -1, errors.New("no string found")
}

func check_line(s string, include_text bool) int64 {
	var total int64
	for i := 0; i < len(s); i++ {
		val, int_err := check_is_int(string(s[i]))
		if int_err == nil {
			total = total + val*10
			break
		}
		if include_text {
			str_val, str_err := check_is_text(s, i)
			if str_err == nil {
				total = total + str_val*10
				break
			}
		}
	}

	for j := len(s) - 1; j >= 0; j-- {
		val, int_err := check_is_int(string(s[j]))
		if int_err == nil {
			total = total + val
			break
		}
		if include_text {
			str_val, str_err := check_is_text(s, j)
			if str_err == nil {
				total = total + str_val
				break
			}
		}
	}
	return total
}

func calc_results(data_as_list []string) (int64, int64) {
	var one_total int64
	var two_total int64
	for _, s := range data_as_list {
		one_total = one_total + check_line(s, false)
		two_total = two_total + check_line(s, true)
	}
	return one_total, two_total
}

func main() {
	data, parse_err := os.ReadFile("day_1_input.txt")
	if parse_err != nil {
		fmt.Printf("Error reading file: %s\n", parse_err)
		return
	}
	as_string := string(data)
	as_list := strings.Split(as_string, "\n")
	part_one_ans, part_two_ans := calc_results(as_list)
	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d", part_two_ans)
}
