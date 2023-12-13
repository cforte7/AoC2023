package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func check_for_extra(lookup map[int]bool, start int, end int) int {
	extra := 0

	s := min(start, end)
	e := max(start, end)
	for i := s; i <= e; i++ {
		_, found := lookup[i]
		if !found {
			extra += 1
		}
	}

	return extra
}

func part_one(data []string, multiplier int) int {
	occupied_rows := map[int]bool{}
	occupied_cols := map[int]bool{}
	planet_locations := [][2]int{}

	for i, row := range data {
		for j, c := range row {
			if c == '#' {
				occupied_rows[i] = true
				occupied_cols[j] = true
				planet_locations = append(planet_locations, [2]int{i, j})
			}
		}
	}

	total := 0
	for i := 0; i < len(planet_locations); i++ {
		for j := i + 1; j < len(planet_locations); j++ {
			start_planet := planet_locations[i]
			end_planet := planet_locations[j]
			dist := int(math.Abs(float64(end_planet[0]-start_planet[0])) + math.Abs(float64(end_planet[1]-start_planet[1])))

			extra_rows := check_for_extra(occupied_rows, start_planet[0], end_planet[0]) * multiplier
			extra_cols := check_for_extra(occupied_cols, start_planet[1], end_planet[1]) * multiplier
			trip_total := (dist + extra_rows + extra_cols)
			// fmt.Printf("for planets: %d(%d, %d) and , %d(%d, %d), total %d with dist %d with extra rows: %d and extra cols: %d \n", i+1, start_planet[0], start_planet[1], j+1, end_planet[0], end_planet[1], trip_total, dist, extra_rows, extra_cols)
			total += trip_total
		}
	}
	return total
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")

	as_rows := strings.Split(string(file), "\n")

	part_one_ans := part_one(as_rows, 1)
	part_two_ans := part_one(as_rows, 1000000-1)
	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
