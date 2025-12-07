package main

import (
	"AoC2025/input"
	"fmt"
	"strconv"
	"strings"
)

func calc(op rune, nums []int) int {
	res := nums[0]
	for _, n := range nums[1:] {
		if op == '*' {
			res *= n
		} else {
			res += n
		}
	}
	return res
}

func part1(lines []string) any {
	var cols [][]int
	for i := 0; i < len(lines)-1; i++ {
		for j, f := range strings.Fields(lines[i]) {
			n, _ := strconv.Atoi(f)
			for len(cols) <= j {
				cols = append(cols, nil)
			}
			cols[j] = append(cols[j], n)
		}
	}

	sum, ops := 0, strings.Fields(lines[len(lines)-1])
	for i, col := range cols {
		sum += calc(rune(ops[i][0]), col)
	}
	return sum
}

func part2(lines []string) any {
	w := 0
	for _, l := range lines {
		if len(l) > w {
			w = len(l)
		}
	}

	var cols [][]rune
	for c := w - 1; c >= 0; c-- {
		var col []rune
		for _, l := range lines {
			if c < len(l) {
				col = append(col, rune(l[c]))
			} else {
				col = append(col, ' ')
			}
		}
		cols = append(cols, col)
	}

	var grps [][][]rune
	var g [][]rune
	for _, col := range cols {
		empty := strings.TrimSpace(string(col)) == ""
		if empty {
			if len(g) > 0 {
				grps = append(grps, g)
				g = nil
			}
		} else {
			g = append(g, col)
		}
	}
	if len(g) > 0 {
		grps = append(grps, g)
	}

	sum := 0
	for _, g := range grps {
		var op rune
		for _, col := range g {
			if ch := col[len(col)-1]; ch == '*' || ch == '+' {
				op = ch
				break
			}
		}

		var nums []int
		for _, col := range g {
			var b strings.Builder
			for i := 0; i < len(col)-1; i++ {
				if ch := col[i]; ch >= '0' && ch <= '9' {
					b.WriteRune(ch)
				}
			}
			if s := b.String(); s != "" {
				n, _ := strconv.Atoi(s)
				nums = append(nums, n)
			}
		}

		sum += calc(op, nums)
	}
	return sum
}

func main() {
	lines := input.ReadLines("day06/in.txt")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}
