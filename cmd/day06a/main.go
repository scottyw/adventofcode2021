package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(simulate(aoc.FileToIntSlice("input/06.txt"), 80))
}

func simulate(fish []int, days int) int {

	for x := 0; x < days; x++ {

		// Tick existing fish
		var newFish int
		for i := range fish {
			if fish[i] == 0 {
				fish[i] = 6
				newFish++
			} else {
				fish[i]--
			}
		}

		// Add new fish
		for i := 0; i < newFish; i++ {
			fish = append(fish, 8)
		}

	}

	return len(fish)
}
