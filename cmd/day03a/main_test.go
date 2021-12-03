package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

// 00100
// 11110
// 10110
// 10111
// 10101
// 01111
// 00111
// 11100
// 10000
// 11001
// 00010
// 01010

// So, the gamma rate is the binary number 10110, or 22 in decimal.

func TestPower(t *testing.T) {

	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	assert.Equal(t, 198, power(lines))

}
