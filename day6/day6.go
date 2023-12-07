package main

import (
	"fmt"
	"time"
)

func solve_race(time int, distance int) int {
	count := 0
	for i := 0; i < time; i++ {
		speed := i
		time_remaining := time - i
		distance_travelled := speed * time_remaining
		if distance_travelled > distance {
			count += 1
		}
	}
	return count
}

func part_one(times []int, distances []int) int {
	ans := 1
	for i := 0; i < len(times); i++ {
		res := solve_race(times[i], distances[i])
		ans = ans * res
	}
	return ans
}

func calc_dist(time_pressed int, total_time int) int {
	return time_pressed * (total_time - time_pressed)
}

func find_smallest_true(time int, distance int) int {
	right := time
	left := 0
	half_way := (right + left) / 2
	for right-left > 1 {
		if calc_dist(half_way, time) > distance {
			right = half_way
		} else {
			left = half_way
		}
		half_way = (right + left) / 2
	}
	return right
}

func find_largest_true(time int, distance int) int {
	right := time
	left := 0
	half_way := (right + left) / 2
	for right-left > 1 {
		if calc_dist(half_way, time) >= distance {
			left = half_way
		} else {
			right = half_way
		}
		half_way = (right + left) / 2
	}
	return right
}
func part_two() int {
	time := 47986698
	distance := 400121310111540
	min_ans := find_smallest_true(time, distance)
	max_ans := find_largest_true(time, distance)
	return max_ans - min_ans
}

func main() {
	start := time.Now()
	times := []int{47, 98, 66, 98}
	distances := []int{400, 1213, 1011, 1540}

	part_one_ans := part_one(times, distances)
	part_two_ans := part_two()

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
