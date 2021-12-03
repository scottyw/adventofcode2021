package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(power(aoc.FileLinesToStringSlice("input/03.txt")))
}

func power(lines []string) int {
	counts := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				counts[i]++
			}
		}
	}
	var gamma, epsilon int
	for _, count := range counts {
		gamma <<= 1
		epsilon <<= 1
		if count >= len(lines)/2 {
			gamma |= 0x01
		} else {
			epsilon |= 0x01
		}
	}
	return gamma * epsilon
}
