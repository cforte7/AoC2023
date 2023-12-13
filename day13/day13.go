package main

import (
	"fmt"
	"os"
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

func part_one(pattern []string, allowed_err int) int {
	hlen := len(pattern[0])
	vlen := len(pattern)

	vrefelct := 0
	for r_start := 1; r_start < hlen; r_start++ {
		if is_v_reflection(pattern, r_start, allowed_err) {
			vrefelct = r_start
			break
		}
	}

	hreflect := 0
	for b_start := 1; b_start < vlen; b_start++ {
		if is_h_reflection(pattern, b_start, allowed_err) {
			hreflect = b_start
			break
		}
	}

	return vrefelct + hreflect*100
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")

	as_rows := strings.Split(string(file), "\n\n")
	part_one_ans := 0
	part_two_ans := 0
	for _, row := range as_rows {
		pattern := strings.Split(row, "\n")
		part_one_ans += part_one(pattern, 0)
		part_two_ans += part_one(pattern, 1)
	}

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
