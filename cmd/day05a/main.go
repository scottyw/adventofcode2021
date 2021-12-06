package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

// 8,0 -> 0,8
var r = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	fmt.Println(overlaps(aoc.FileLinesToStringSlice("input/05.txt")))
}

func overlaps(lines []string) int {

	var field [1000][1000]int

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		matches := r.FindStringSubmatch(line)
		x1 := aoc.Atoi(matches[1])
		y1 := aoc.Atoi(matches[2])
		x2 := aoc.Atoi(matches[3])
		y2 := aoc.Atoi(matches[4])

		if x1 != x2 && y1 != y2 {
			continue
		}

		if x1 == x2 {
			start, end := sort(y1, y2)
			for i := start; i <= end; i++ {
				field[x1][i]++
			}
		}

		if y1 == y2 {
			start, end := sort(x1, x2)
			for i := start; i <= end; i++ {
				field[i][y1]++
			}
		}

	}

	var total int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if field[x][y] > 1 {
				total++
			}
		}
	}
	return total

}

func sort(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}
