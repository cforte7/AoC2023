package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"
)

func calc_score(grid [][]string) int {
	tot := 0
	grid_len := len(grid)
	for row_num, v := range grid {
		row_count := 0
		for _, c := range v {
			if c == "O" {
				row_count += 1
			}
		}
		tot += row_count * (grid_len - row_num)
	}
	return tot
}

func shift_up(grid [][]string) [][]string {
	row_count := len(grid)
	col_count := len(grid[0])
	for col := col_count - 1; col >= 0; col-- {
		stop := 0
		for row := 0; row < row_count; row++ {
			if grid[row][col] == "O" {
				grid[row][col] = "."
				grid[stop][col] = "O"
				stop += 1
			} else if grid[row][col] == "#" {
				stop = row + 1
			}
		}
	}
	return grid
}
func part_one(grid [][]string) int {
	shift_up(grid)
	return calc_score(grid)
}

func rotateMatrix(matrix [][]string) [][]string {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func computeHashForList(list [][]string) [32]byte {
	var buffer bytes.Buffer
	for i, row := range list {
		for j := range row {
			buffer.WriteString(list[i][j])
		}
	}
	return sha256.Sum256((buffer.Bytes()))
}

type PastRes struct {
	Pos int
	Res [][]string
}

func part_two(grid [][]string) int {
	cache := map[[32]byte]PastRes{}
	cycle := 0
	rem := -1
	for i := 0; i < 1000000000; i++ {
		grid_hash := computeHashForList(grid)
		v, found := cache[grid_hash]
		if found {
			cycle = i - v.Pos
			rem = (1000000000 - i - cycle) % cycle
			break
		} else {
			for j := 0; j < 4; j++ {
				grid = shift_up(grid)
				grid = rotateMatrix(grid)
			}
			cache[grid_hash] = PastRes{i, grid}
		}
	}
	for i := 0; i < rem; i++ {
		for j := 0; j < 4; j++ {
			grid = shift_up(grid)
			grid = rotateMatrix(grid)
		}
	}
	return calc_score(grid)
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("./input.txt")
	as_rows := strings.Split(string(file), "\n")

	full_split_one := [][]string{}
	full_split_two := [][]string{}
	for _, v := range as_rows {
		as_rune := strings.Split(v, "")
		full_split_one = append(full_split_one, as_rune)
		full_split_two = append(full_split_two, as_rune)
	}

	part_one_ans := part_one(full_split_one)
	part_two_ans := part_two(full_split_two)

	fmt.Printf("Part one: %d\n", part_one_ans)
	fmt.Printf("Part two: %d\n", part_two_ans)
	fmt.Printf("Time elapsed: %d milliseconds\n", time.Now().UnixMilli()-start.UnixMilli())
}
