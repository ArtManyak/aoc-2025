package main

import (
	"AoC2025/input"
	"fmt"
)

func part1(lines []string) int {
	return calcSum(lines, 2)
}

func part2(lines []string) int {
	return calcSum(lines, 12)
}

func calcSum(lines []string, count int) int {
	sum := 0
	for _, line := range lines {
		current := 0
		lastIdx := -1
		for k := count - 1; k >= 0; k-- {
			maxDigit := uint8(0)
			for i := lastIdx + 1; i < len(line)-k; i++ {
				digit := line[i] - '0'
				if digit > maxDigit {
					maxDigit = digit
					lastIdx = i
				}
			}
			current = current*10 + int(maxDigit)
		}
		sum += current
	}
	return sum
}

func main() {
	lines := input.ReadLines("day03/in.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
