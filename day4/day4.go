package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

func parse_line(line string) ([]string, []string) {
	re := regexp.MustCompile(`:\s*([0-9\s]+)\s*\|\s*([0-9\s]+)`)
	matches := re.FindStringSubmatch(line)
	firstList := regexp.MustCompile(`\s+`).Split(matches[1], -1)
	secondList := regexp.MustCompile(`\s+`).Split(matches[2], -1)
	return firstList, secondList
}

func get_line_score(line string) int {
	winners, to_check := parse_line(line)
	winner_lookup := map[string]bool{}
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
	start := time.Now()
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
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
