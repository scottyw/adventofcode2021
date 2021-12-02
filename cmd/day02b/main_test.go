package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

// forward 5
// down 5
// forward 8
// up 3
// down 8
// forward 2

// After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

func TestPosition(t *testing.T) {

	lines := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	x, y := findPosition(lines)
	assert.Equal(t, 15, x)
	assert.Equal(t, 60, y)

}
