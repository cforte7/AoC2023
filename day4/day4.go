package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func splitNumbers(s string) []int {
	split_list := regexp.MustCompile(`\s+`).Split(s, -1)
	var out []int
	for _, v := range split_list {
		num, _ := strconv.Atoi(v)
		out = append(out, num)
	}
	return out
}
func parse_line(line string) ([]int, []int) {
	re := regexp.MustCompile(`:\s*([0-9\s]+)\s*\|\s*([0-9\s]+)`)
	matches := re.FindStringSubmatch(line)
	firstList := strings.TrimSpace(matches[1])
	secondList := strings.TrimSpace(matches[2])
	firstNumbers := splitNumbers(firstList)
	secondNumbers := splitNumbers(secondList)
	return firstNumbers, secondNumbers
}

func get_line_score(line string) int {
	winners, to_check := parse_line(line)
	winner_lookup := map[int]bool{}
	for _, v := range winners {
		winner_lookup[v] = true
	}
	matches := 0
	for _, v := range to_check {
		_, match := winner_lookup[v]
		if match {
			matches += 1
		}
	}
	return matches
}

type Card struct {
	Score int
	Count int
}

func main() {
	file2, _ := os.ReadFile("./input.txt")
	as_rows := strings.Split(string(file2), "\n")

	part_one_ans := 0
	part_two_ans := 0
	cards := []Card{}

	for range as_rows {
		cards = append(cards, Card{Count: 1})
	}

	for i, v := range as_rows {
		score := get_line_score(v)
		cards[i].Score = score
		part_one_ans += int(math.Pow(float64(2), float64(score-1)))
		for j := (i + 1); j <= i+cards[i].Score && j < len(cards); j++ {
			cards[j].Count += cards[i].Count
		}
		part_two_ans += cards[i].Count
	}

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d", part_two_ans)
}
