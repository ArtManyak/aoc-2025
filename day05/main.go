package main

import (
	"AoC2025/input"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Segment struct {
	Start, Finish int
}

func parseInput(lines []string) ([]Segment, []int) {
	i := 0
	var segments []Segment
	for lines[i] != "" {
		parts := strings.Split(lines[i], "-")
		start, _ := strconv.Atoi(parts[0])
		finish, _ := strconv.Atoi(parts[1])
		segments = append(segments, Segment{start, finish})
		i++
	}

	i++
	var points []int
	for ; i < len(lines); i++ {
		point, _ := strconv.Atoi(lines[i])
		points = append(points, point)
	}

	return mergeSegments(segments), points
}

func mergeSegments(segments []Segment) []Segment {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].Start < segments[j].Start
	})

	merged := []Segment{segments[0]}
	for _, seg := range segments[1:] {
		last := &merged[len(merged)-1]
		if seg.Start <= last.Finish {
			if seg.Finish > last.Finish {
				last.Finish = seg.Finish
			}
		} else {
			merged = append(merged, seg)
		}
	}
	return merged
}

func part1(segments []Segment, points []int) int {
	count := 0
	for _, point := range points {
		for _, seg := range segments {
			if point >= seg.Start && point <= seg.Finish {
				count++
				break
			}
		}
	}
	return count
}

func part2(segments []Segment) int {
	count := 0
	for _, seg := range segments {
		count += seg.Finish - seg.Start + 1
	}
	return count
}

func main() {
	segments, points := parseInput(input.ReadLines("day05/in.txt"))
	fmt.Println("Part 1:", part1(segments, points))
	fmt.Println("Part 2:", part2(segments))
}
