package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(findIncreases(aoc.FileLinesToIntSlice("input/01.txt")))
}

func findIncreases(depths []int) int {
	var count int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return count
}
