package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func is_v_reflection(pattern []string, r_start int, allowed_err int) bool {
	hlen := len(pattern[0])
	vlen := len(pattern)
	l_search := r_start - 1
	r_serach := r_start
	errors := 0
	for l_search >= 0 && r_serach < hlen {
		// search going down
		for s := 0; s < vlen; s++ {
			if pattern[s][l_search] != pattern[s][r_serach] {
				errors += 1
				if errors > allowed_err {
					return false
				}
			}
		}
		l_search -= 1
		r_serach += 1
	}
	return errors == allowed_err
}

func is_h_reflection(pattern []string, b_start int, allowed_err int) bool {
	hlen := len(pattern[0])
	vlen := len(pattern)
	t_search := b_start - 1
	b_serach := b_start
	errors := 0
	for t_search >= 0 && b_serach < vlen {
		// search going right
		for s := 0; s < hlen; s++ {
			if pattern[t_search][s] != pattern[b_serach][s] {
				errors += 1
				if errors > allowed_err {
					return false
				}
			}
		}
		t_search -= 1
		b_serach += 1
	}
	return errors == allowed_err
}

func parse_row(row string) (string, []int) {

	f_split := strings.Split(row, " ")
	records := f_split[0]
	counts := []int{}
	for _, v := range strings.Split(f_split[1], ",") {
		as_int, _ := strconv.Atoi(v)
		counts = append(counts, as_int)
	}
	return records, counts
}

func part_one(row string) int {
	records, counts := parse_row(row)
	return -1
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./test.txt")

	as_rows := strings.Split(string(file), "\n")
	part_one_ans := 0
	part_two_ans := 0
	for _, row := range as_rows {

		part_one_ans += part_one(row)
		// part_two_ans += part_one(pattern, 1)
	}

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
