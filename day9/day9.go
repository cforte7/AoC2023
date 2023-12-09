package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func part_one(line []int) int {
	all_zero := true
	for _, v := range line {
		if v != 0 {
			all_zero = false
		}
	}

	if all_zero {
		return 0
	}
	next := []int{}

	for i := 0; i < len(line)-1; i += 1 {
		next = append(next, line[i+1]-line[i])
	}

	return line[len(line)-1] + part_one(next)
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")
	as_rows := strings.Split(string(file), "\n")
	part_one_ans := 0
	part_two_ans := 0
	for _, line := range as_rows {
		l := strings.Split(line, " ")
		as_int := []int{}
		for _, c := range l {
			ii, _ := strconv.Atoi(c)
			as_int = append(as_int, ii)
		}
		part_one_ans += part_one(as_int)
		// reverse list for part 2
		for i, j := 0, len(as_int)-1; i < j; i, j = i+1, j-1 {
			as_int[i], as_int[j] = as_int[j], as_int[i]
		}

		part_two_ans += part_one(as_int)
	}

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
