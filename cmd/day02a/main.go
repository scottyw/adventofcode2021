package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

// forward 5
var r = regexp.MustCompile(`(\w+) (\d+)`)

func main() {
	x, y := findPosition(aoc.FileLinesToStringSlice("input/02.txt"))
	fmt.Println(x * y)
}

func findPosition(lines []string) (int, int) {
	var x, y int
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		direction := matches[1]
		distance := aoc.Atoi(matches[2])
		switch direction {
		case "forward":
			x += distance
		case "up":
			y -= distance
		case "down":
			y += distance
		default:
			panic(direction)
		}
	}
	return x, y
}
