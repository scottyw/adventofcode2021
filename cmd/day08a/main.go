package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

var r = regexp.MustCompile(`(.+) \| (.+)`)

func main() {
	fmt.Println(countUnique(aoc.FileLinesToStringSlice("input/08.txt")))
}

func countUnique(lines []string) int {

	var total int
	for _, line := range lines {

		if line == "" {
			continue
		}

		matches := r.FindStringSubmatch(line)
		// inputs := strings.Split(matches[1], " ")
		outputs := strings.Split(matches[2], " ")

		total += count(outputs)

	}

	return total
}

func count(outputs []string) int {
	var total int
	for _, output := range outputs {
		switch len(output) {
		case 2, 3, 4, 7:
			total++
		}
	}
	return total
}
