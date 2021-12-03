package main

import (
	"fmt"
	"strconv"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	fmt.Println(support(aoc.FileLinesToStringSlice("input/03.txt")))
}

func support(lines []string) int {

	oxygenCandidates := make([]string, len(lines))
	copy(oxygenCandidates, lines)

	co2Candidates := make([]string, len(lines))
	copy(co2Candidates, lines)

	for i := 0; i < len(lines[0]) && len(oxygenCandidates) > 1; i++ {
		_, oxygenCandidates = split(oxygenCandidates, i)
	}

	for i := 0; i < len(lines[0]) && len(co2Candidates) > 1; i++ {
		co2Candidates, _ = split(co2Candidates, i)
	}

	o2, err := strconv.ParseInt(oxygenCandidates[0], 2, 64)
	if err != nil {
		panic(err)
	}

	co2, err := strconv.ParseInt(co2Candidates[0], 2, 64)
	if err != nil {
		panic(err)
	}

	return int(o2 * co2)
}

func split(xs []string, i int) ([]string, []string) {
	var as []string
	var bs []string
	for _, x := range xs {
		if x[i] == '0' {
			as = append(as, x)
		} else {
			bs = append(bs, x)
		}
	}
	if len(as) > len(bs) {
		return bs, as
	}
	return as, bs
}
