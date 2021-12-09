package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(fuel(aoc.FileToIntSlice("input/07.txt")))
}

func fuel(crabs []int) int {

	var mean int
	for _, crab := range crabs {
		mean += crab
	}
	mean /= len(crabs)

	cost1 := cost(crabs, mean)
	cost2 := cost(crabs, mean+1)
	if cost1 > cost2 {
		return cost2
	}
	return cost1

}

func cost(crabs []int, target int) int {
	var total int
	for _, crab := range crabs {
		diff := target - crab
		if diff < 0 {
			diff = -diff
		}
		total += (diff * (diff + 1)) / 2
	}
	return total
}
