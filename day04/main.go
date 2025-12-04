package main

import (
	"AoC2025/input"
	"AoC2025/utils"
	"fmt"
)

func changeChar(s string, i int, c byte) string {
	b := []byte(s)
	b[i] = c
	return string(b)
}

func countRollsAround(lines []string, i, j int) int {
	count := 0
	for _, dir := range utils.Directions8 {
		ni, nj := i+dir.Di, j+dir.Dj
		if ni >= 0 && ni < len(lines) && nj >= 0 && nj < len(lines[i]) && lines[ni][nj] == '@' {
			count++
		}
	}
	return count
}

func part1(lines []string) int {
	ans := 0
	for i, line := range lines {
		for j := range line {
			if line[j] == '@' && countRollsAround(lines, i, j) < 4 {
				ans++
			}
		}
	}
	return ans
}

func part2(lines []string) int {
	ans := 0
	for {
		cur := 0
		for i, line := range lines {
			for j := range line {
				if line[j] == '@' && countRollsAround(lines, i, j) < 4 {
					cur++
					lines[i] = changeChar(lines[i], j, '.')
				}
			}
		}
		ans += cur
		if cur == 0 {
			break
		}
	}
	return ans
}

func main() {
	lines := input.ReadLines("day04/in.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
