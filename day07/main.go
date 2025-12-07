package main

import (
	"AoC2025/input"
	"fmt"
	"strings"
)

func findStart(grid []string) int {
	return strings.IndexByte(grid[0], 'S')
}

func simulate(grid []string) (splits int, beams map[int]int) {
	beams = map[int]int{findStart(grid): 1}

	for row := 1; row < len(grid); row++ {
		next := make(map[int]int)

		for col, count := range beams {
			if grid[row][col] == '^' {
				splits++
				next[col-1] += count
				next[col+1] += count
			} else {
				next[col] += count
			}
		}

		beams = next
	}

	return splits, beams
}

func countPaths(beams map[int]int) int {
	total := 0
	for _, count := range beams {
		total += count
	}
	return total
}

func main() {
	grid := input.ReadLines("day07/in.txt")
	splits, beams := simulate(grid)
	fmt.Println("Part 1:", splits)
	fmt.Println("Part 2:", countPaths(beams))
}
