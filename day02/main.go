package main

import (
	"AoC2025/input"
	"fmt"
	"strconv"
	"strings"
)

func parseRange(rng string) (int, int) {
	parts := strings.Split(rng, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return start, end
}

func isHalfMatch(s string) bool {
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func part1(line string) int {
	sum := 0
	for _, rng := range strings.Split(line, ",") {
		start, end := parseRange(rng)
		for num := start; num <= end; num++ {
			if isHalfMatch(strconv.Itoa(num)) {
				sum += num
			}
		}
	}
	return sum
}

func hasRepeatingPattern(s string) bool {
	for patternLen := 1; patternLen <= len(s)/2; patternLen++ {
		if len(s)%patternLen != 0 {
			continue
		}
		pattern := s[:patternLen]
		matches := true
		for i := patternLen; i < len(s); i += patternLen {
			if s[i:i+patternLen] != pattern {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}

func part2(line string) int {
	sum := 0
	for _, rng := range strings.Split(line, ",") {
		start, end := parseRange(rng)
		for num := start; num <= end; num++ {
			if hasRepeatingPattern(strconv.Itoa(num)) {
				sum += num
			}
		}
	}
	return sum
}

func main() {
	lines := input.ReadLines("day02/in.txt")
	for _, line := range lines {
		fmt.Println("Part 1:", part1(line))
		fmt.Println("Part 2:", part2(line))
	}
}
