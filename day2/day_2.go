package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lookup = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func eval_game(game string) bool {
	arr_of_dice := strings.Split(game, ", ")
	for _, die := range arr_of_dice {
		res := strings.Split(die, " ")
		roll_count, _ := strconv.Atoi(res[0])
		if lookup[res[1]] < roll_count {
			return false
		}
	}
	return true
}

func evaluate_games(games string) bool {
	arr_of_games := strings.Split(games, "; ")
	for _, game := range arr_of_games {
		if !eval_game(game) {
			return false
		}
	}
	return true
}

func parse_line_part_one(line string) int {
	res := strings.Split(line, ": ")
	game_title := strings.Split(res[0], " ")
	game_num := game_title[1]
	if evaluate_games(res[1]) {
		as_num, _ := strconv.Atoi(game_num)
		return as_num
	}
	return 0
}

func parse_line_part_two(line string) int {
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
	return max_die["red"] * max_die["green"] * max_die["blue"]
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	sum_one := 0
	sum_two := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum_one += parse_line_part_one(scanner.Text())
		sum_two += parse_line_part_two(scanner.Text())
	}

	fmt.Printf("Part one: %d\n", sum_one)
	fmt.Printf("Part two: %d", sum_two)
}
