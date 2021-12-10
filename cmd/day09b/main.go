package main

import (
	"fmt"
	"sort"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(find(aoc.FileLinesToStringSlice("input/09.txt")))
}

func find(lines []string) int {
	depths := parse(lines)

	var totals []int
	for y := range depths {
		for x := range depths[y] {
			if lower(depths, x, y, x, y-1) &&
				lower(depths, x, y, x, y+1) &&
				lower(depths, x, y, x-1, y) &&
				lower(depths, x, y, x+1, y) {
				totals = append(totals, size(depths, x, y))
			}
		}
	}

	sort.Ints(totals)

	return totals[len(totals)-1] * totals[len(totals)-2] * totals[len(totals)-3]
}

func size(depths [][]int, x, y int) int {
	if y < 0 || y == len(depths) || x < 0 || x == len(depths[y]) {
		return 0
	}
	if depths[y][x] >= 0 && depths[y][x] < 9 {
		depths[y][x] = -1
		return 1 + size(depths, x, y-1) + size(depths, x, y+1) + size(depths, x-1, y) + size(depths, x+1, y)
	}
	return 0
}

func lower(depths [][]int, x, y, x2, y2 int) bool {
	if y2 < 0 || y2 == len(depths) || x2 < 0 || x2 == len(depths[y]) {
		return true
	}
	return depths[y][x] < depths[y2][x2]
}

func parse(lines []string) [][]int {
	var depths [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		row := []int{}
		for _, r := range line {
			row = append(row, aoc.Atoi(string(r)))
		}
		depths = append(depths, row)
	}
	return depths
}
