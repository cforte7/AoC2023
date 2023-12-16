package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func get_value(data []byte) int {
	value := 0
	for _, c := range data {
		value += int(c)
		value *= 17
		value %= 256
	}
	return value
}

func part_one(data [][]byte) int {
	tot := 0
	for _, v := range data {
		tot += get_value(v)
	}
	return tot
}

const DASH_VAL byte = 45
const EQUAL_VAL byte = 61

type Lens struct {
	Label    string
	FocalLen int
}

type Box struct {
	Slots []Lens
}

func (b *Box) add_or_replace_lens(incoming Lens) *Box {
	for i, v := range b.Slots {
		if v.Label == incoming.Label {
			b.Slots[i] = incoming
			return b
		}
	}
	b.Slots = append(b.Slots, incoming)
	return b
}

func (b *Box) remove_lens(label string) {
	new_slots := []Lens{}
	for _, v := range b.Slots {
		if v.Label != label {
			new_slots = append(new_slots, v)
		}
	}
	b.Slots = new_slots
}
func (b Box) focusing_power(box_num int) int {
	power := 0
	for i, v := range b.Slots {

		lpower := (box_num + 1) * (i + 1) * v.FocalLen
		power += lpower

	}
	return power
}

func build_boxes(data [][]byte) []*Box {
	boxes := []*Box{}
	for i := 0; i < 256; i++ {
		boxes = append(boxes, &Box{[]Lens{}})
	}

	for _, v := range data {
		action_ind := 0
		for v[action_ind] != DASH_VAL && v[action_ind] != EQUAL_VAL {
			action_ind += 1
		}
		// var focal_len byte = 0
		label := string(v[:action_ind])
		action := v[action_ind]
		box_num := get_value(v[:action_ind])
		box := boxes[box_num]
		if action == 61 {
			focal_len, _ := strconv.Atoi(string(v[action_ind+1]))
			lens := Lens{label, focal_len}
			boxes[box_num] = box.add_or_replace_lens(lens)
		} else {
			box.remove_lens(label)
		}
		boxes[box_num] = box

	}
	return boxes
}

func part_two(data [][]byte) int {
	// first index is the box number
	// second index is the position of the lens in the box
	total := 0
	boxes := build_boxes(data)
	for i, v := range boxes {
		total += v.focusing_power(i)
	}
	return total
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")

	as_chunks := [][]byte{}
	chunk := []byte{}

	for i, v := range file {
		if v != 44 {
			chunk = append(chunk, v)
		} else {
			as_chunks = append(as_chunks, chunk)
			chunk = []byte{}
		}

		if i == len(file)-1 && len(chunk) > 0 {
			as_chunks = append(as_chunks, chunk)
		}
	}

	part_one_ans := part_one(as_chunks)
	part_two_ans := part_two(as_chunks)

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
