package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(simulate(aoc.FileToIntSlice("input/06.txt"), 256))
}

func simulate(fish []int, days int) int {

	// Reorganize the input into counts
	fs := [9]int{}
	for _, f := range fish {
		fs[f]++
	}

	// Iterate days
	for x := 0; x < days; x++ {
		new := fs[0]
		for i := 0; i < 8; i++ {
			fs[i] = fs[i+1]
		}
		fs[6] += new
		fs[8] = new
	}

	// Count all the fish
	var total int
	for _, f := range fs {
		total += f
	}

	return total
}
