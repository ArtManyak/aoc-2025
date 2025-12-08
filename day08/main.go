package main

import (
	"AoC2025/input"
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type V3 struct {
	x, y, z int
}

type Edge struct {
	i, j int
	dist float64
}

type UnionFind struct {
	parent []int
	size   []int
}

func newUnionFind(n int) UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf UnionFind) union(x, y int) bool {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX == rootY {
		return false
	}
	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

func (uf UnionFind) groupSizes() []int {
	groups := make(map[int]int)
	for i := range uf.parent {
		groups[uf.find(i)]++
	}
	return slices.Collect(maps.Values(groups))
}

func parsePoints(lines []string) []V3 {
	points := make([]V3, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points[i] = V3{x, y, z}
	}
	return points
}

func distance(p1, p2 V3) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func buildEdges(points []V3) []Edge {
	n := len(points)
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, Edge{i, j, distance(points[i], points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	return edges
}

func part1(lines []string, n int) int {
	points := parsePoints(lines)
	edges := buildEdges(points)
	uf := newUnionFind(len(points))

	for i := 0; i < n; i++ {
		uf.union(edges[i].i, edges[i].j)
	}

	sizes := uf.groupSizes()
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
}

func part2(lines []string) int {
	points := parsePoints(lines)
	edges := buildEdges(points)
	uf := newUnionFind(len(points))

	numGroups := len(points)
	for _, e := range edges {
		if uf.union(e.i, e.j) {
			numGroups--
			if numGroups == 1 {
				return points[e.i].x * points[e.j].x
			}
		}
	}
	return 0
}

func main() {
	lines := input.ReadLines("day08/in.txt")
	fmt.Println("Part 1:", part1(lines, 1000))
	fmt.Println("Part 2:", part2(lines))
}
