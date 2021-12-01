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
	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			count++
		}
	}
	return count
}
