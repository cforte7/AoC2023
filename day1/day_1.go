package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_is_int(s string, multiplier int64) (int64, error) {
	val, int_err := strconv.ParseInt(s, 10, 64)
	if int_err == nil {
		return val * multiplier, nil
	}
	return -1, errors.New("not a digit")
}

func part_one(data_as_list []string) int64 {
	var total int64

	for _, s := range data_as_list {
		for i := 0; i < len(s); i++ {
			val, int_err := check_is_int(string(s[i]), 10)
			if int_err == nil {
				// fmt.Printf("adding new value: %d from string %s of list %s", val, string(s[i]), s)
				total = total + val
				break
			}
		}

		for j := len(s) - 1; j >= 0; j-- {
			val, int_err := check_is_int(string(s[j]), 1)
			if int_err == nil {
				total = total + val
				break
			}
		}
	}

	return total
}

// var lookup = map[string]int{
// 	"one":   1,
// 	"two":   2,
// 	"three": 3,
// 	"four":  4,
// 	"five":  5,
// 	"six":   6,
// 	"seven": 7,
// 	"eight": 8,
// 	"nine":  9,
// }

// func check_is_text(s string, index int) {

// }

// func part_two(data_as_list []string) int64 {

// }

func main() {
	data, parse_err := os.ReadFile("day_1_input.txt")
	if parse_err != nil {
		fmt.Printf("Error reading file: %s\n", parse_err)
		return
	}
	as_string := string(data)
	as_list := strings.Split(as_string, "\n")
	part_one_ans := part_one(as_list)
	print(part_one_ans)

}
