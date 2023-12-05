package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type MappingSet struct {
	mappings []Mapping
}
type Mapping struct {
	Source      int
	Destination int
	Range       int
}

func (m MappingSet) apply_mapping(val int) int {
	for _, mapping := range m.mappings {
		if mapping.Source <= val && val < mapping.Source+mapping.Range {
			diff := val - mapping.Source
			return mapping.Destination + diff
		}
	}
	return val
}

func parse_maps(maps []string) []MappingSet {
	var mappings []MappingSet
	for _, map_val := range maps {
		var map_list []Mapping
		for _, line := range strings.Split(map_val, "\n")[1:] {
			as_ints := strings.Split(line, " ")
			dest_start, _ := strconv.Atoi(as_ints[0])
			source_start, _ := strconv.Atoi(as_ints[1])
			range_len, _ := strconv.Atoi(as_ints[2])
			map_list = append(map_list, Mapping{
				Source:      source_start,
				Destination: dest_start,
				Range:       range_len,
			})
		}
		mappings = append(mappings, MappingSet{mappings: map_list})
	}
	return mappings
}

func part_one(seeds []int, map_sets []MappingSet) int {
	ans := math.MaxInt32
	for _, map_set := range map_sets {
		for i, seed := range seeds {
			seeds[i] = map_set.apply_mapping(seed)
		}
	}
	for _, v := range seeds {
		ans = min(ans, v)
	}
	return ans
}

func part_two(seeds []int, map_sets []MappingSet) int {
	ans := math.MaxInt32
	for i := 0; i <= len(seeds)/2; i = i + 2 {
		seed_start := seeds[i]
		seed_range := seeds[i+1]
		for seed := seed_start; seed < seed_start+seed_range; seed++ {
			seed_val := seed
			for _, map_set := range map_sets {
				seed_val = map_set.apply_mapping(seed_val)
			}
			ans = min(ans, seed_val)
		}

	}
	return ans
}

func main() {
	start := time.Now()
	file2, _ := os.ReadFile("./input.txt")
	as_rows := strings.Split(string(file2), "\n\n")

	var seeds_part_one []int
	var seeds_part_two []int
	for _, v := range strings.Split(strings.TrimPrefix(as_rows[0], "seeds: "), " ") {
		as_int, _ := strconv.Atoi(v)
		seeds_part_one = append(seeds_part_one, as_int)
		seeds_part_two = append(seeds_part_two, as_int)
	}

	map_sets := parse_maps(as_rows[1:])
	part_one_ans := part_one(seeds_part_one, map_sets)
	part_two_ans := part_two(seeds_part_two, map_sets)

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
