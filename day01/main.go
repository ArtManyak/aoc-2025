package main

import (
	"AoC2025/input"
	"fmt"
	"strconv"
)

const trackSize = 100

func part1(lines []string) int {
	pos := 50
	crossings := 0

	for _, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			pos -= n
		} else {
			pos += n
		}
		pos = ((pos % trackSize) + trackSize) % trackSize
		if pos == 0 {
			crossings++
		}
	}
	return crossings
}

func part2(lines []string) int {
	pos := 50
	laps := 0

	for _, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			if pos == 0 {
				laps--
			}
			pos -= n
			laps += (trackSize - pos) / trackSize
		} else {
			pos += n
			laps += pos / trackSize
		}
		pos = ((pos % trackSize) + trackSize) % trackSize
	}
	return laps
}

func main() {
	lines := input.ReadLines("day01/in.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
