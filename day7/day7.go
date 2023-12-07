package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// func part_one(line string) int {
// s := strings.Split(line, ' ')
// hand := s[0]
// score := s[1]
// 	utf8.RuneCountInString()
// 	return -1
// }

var card_str = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9, '9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1}
var hand_str = map[int]int{5: 7, 6: 5, 3: 4, 2: 2, 1: 1}

/*
7 - Five of a kind, where all five cards have the same label: AAAAA 5
6- Four of a kind, where four cards have the same label and one card has a different label: AA8AA 4
5 -Full house, where three cards have the same label, and the remaining two cards share a different label: 23332 6
4- Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98 3
3- Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432 4
2- One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4 2
1 - High card, where all cards' labels are distinct: 23456 1
*/
type Hand struct {
	TypeScore int
	Cards     string
	Score     int
}

func build_hand(line string) Hand {
	s := strings.Split(line, " ")
	println(s[0], s[1])
	hand := s[0]
	score, _ := strconv.Atoi(s[1])
	hand_lookup := map[rune]int{}
	mult_score := 1
	hand_score := 0
	for _, v := range hand {
		count := hand_lookup[v]
		hand_lookup[v] = count + 1
	}

	for _, v := range hand_lookup {
		mult_score *= v
	}

	if mult_score == 4 && len(hand_lookup) == 2 {

		hand_score = 6
	} else if mult_score == 4 {
		hand_score = 3
	} else {
		score := hand_str[mult_score]
		hand_score = score
	}

	return Hand{
		TypeScore: hand_score,
		Cards:     hand,
		Score:     score,
	}
}

type Hands []Hand

func (h Hands) Len() int { return len(h) }
func (h Hands) Less(i, j int) bool {
	if h[i].TypeScore < h[j].TypeScore {
		return true
	} else if h[i].TypeScore > h[j].TypeScore {
		return false
	} else {
		for k := 0; k < 5; k++ {
			i_str := card_str[rune(h[i].Cards[k])]
			j_str := card_str[rune(h[j].Cards[k])]
			if i_str < j_str {
				return true
			} else if i_str > j_str {
				return false
			}
		}
	}
	return true
}
func (h Hands) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func main() {
	start := time.Now()
	file, _ := os.Open("./input.txt")
	defer file.Close()

	hands := Hands{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands = append(hands, build_hand(scanner.Text()))
	}
	sort.Sort(hands)
	part_one_ans := 0
	for i, v := range hands {
		part_one_ans += (i + 1) * v.Score
	}
	// part_two_ans := part_two()

	fmt.Printf("Part one: %d\n", part_one_ans)
	// fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
