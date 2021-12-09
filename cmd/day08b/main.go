package main

import (
	"fmt"
	"math/bits"
	"regexp"
	"strings"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

var r = regexp.MustCompile(`(.+) \| (.+)`)

func main() {
	fmt.Println(sum(aoc.FileLinesToStringSlice("input/08.txt")))
}

func sum(lines []string) int {

	var total int
	for _, line := range lines {

		if line == "" {
			continue
		}

		var inputs, outputs []uint8
		matches := r.FindStringSubmatch(line)
		for _, input := range strings.Split(matches[1], " ") {
			inputs = append(inputs, toUint8(input))
		}
		for _, output := range strings.Split(matches[2], " ") {
			outputs = append(outputs, toUint8(output))
		}

		total += identify(inputs, outputs)

	}

	return total
}

func identify(inputs, outputs []uint8) int {

	var one, four, nine uint8
	var fives []uint8
	var sixes []uint8

	digits := map[uint8]int{}

	for _, input := range inputs {
		switch bits.OnesCount8(input) {
		case 2:
			digits[input] = 1 // The input is a 1
			one = input
		case 3:
			digits[input] = 7 // The input is a 7
		case 4:
			digits[input] = 4 // The input is a 4
			four = input
		case 5:
			fives = append(fives, input) // The input has 5 segments
		case 6:
			sixes = append(sixes, input) // The input has 6 segments
		case 7:
			digits[input] = 8 // The input is an 8
		}
	}

	for _, input := range sixes {
		if input&four == four {
			digits[input] = 9 // The input is a 9
			nine = input
		} else if input&one == one {
			digits[input] = 0 // The input is a 0
		} else {
			digits[input] = 6 // The input is a 6
		}
	}

	for _, input := range fives {
		if input&one == one {
			digits[input] = 3 // The input is a 3
		} else if input&nine == input {
			digits[input] = 5 // The input is a 5
		} else {
			digits[input] = 2 // The input is a 2
		}
	}

	return digits[outputs[0]]*1000 +
		digits[outputs[1]]*100 +
		digits[outputs[2]]*10 +
		digits[outputs[3]]
}

func toUint8(s string) uint8 {
	var i uint8
	if strings.Contains(s, "a") {
		i |= 0x01
	}
	if strings.Contains(s, "b") {
		i |= 0x02
	}
	if strings.Contains(s, "c") {
		i |= 0x04
	}
	if strings.Contains(s, "d") {
		i |= 0x08
	}
	if strings.Contains(s, "e") {
		i |= 0x10
	}
	if strings.Contains(s, "f") {
		i |= 0x20
	}
	if strings.Contains(s, "g") {
		i |= 0x40
	}
	return i
}
