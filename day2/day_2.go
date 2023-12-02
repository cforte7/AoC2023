package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var max_per_game = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func get_max_per_line(line string) map[string]int {
	res := strings.Split(line, ": ")
	arr_of_games := strings.Split(res[1], "; ")
	var max_die = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, game := range arr_of_games {
		arr_of_die := strings.Split(game, ", ")
		for _, dice := range arr_of_die {
			count_and_color := strings.Split(dice, " ")
			count, _ := strconv.Atoi(count_and_color[0])
			color := count_and_color[1]
			max_die[color] = max(max_die[color], count)
		}
	}
	return max_die
}

func game_is_valid(max_lookup map[string]int) bool {
	for key, val := range max_per_game {
		if max_lookup[key] > val {
			return false
		}
	}
	return true
}

func calc_power_level(max_lookup map[string]int) int {
	power := 1
	for _, val := range max_lookup {
		power *= val
	}
	return power
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	sum_one := 0
	sum_two := 0
	scanner := bufio.NewScanner(file)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		max_lookup := get_max_per_line(scanner.Text())
		if game_is_valid(max_lookup) {
			sum_one += lineNumber
		}
		sum_two += calc_power_level(max_lookup)
	}
	fmt.Printf("Part one: %d\n", sum_one)
	fmt.Printf("Part two: %d\n", sum_two)
}
