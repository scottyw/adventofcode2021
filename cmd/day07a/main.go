package main

import (
	"fmt"
	"sort"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(fuel(aoc.FileToIntSlice("input/07.txt")))
}

func fuel(crabs []int) int {

	sort.Ints(crabs)

	if len(crabs)%2 == 0 {
		median1 := crabs[len(crabs)/2]
		cost1 := cost(crabs, median1)
		median2 := crabs[len(crabs)/2+1]
		cost2 := cost(crabs, median2)
		if cost1 > cost2 {
			return cost2
		}
		return cost1
	}

	median := crabs[len(crabs)/2]
	return cost(crabs, median)

}

func cost(crabs []int, target int) int {
	var total int
	for _, crab := range crabs {
		diff := target - crab
		if diff < 0 {
			total -= diff
		} else {
			total += diff
		}
	}
	return total
}
