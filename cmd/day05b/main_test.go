package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for you to review. For example:

// 0,9 -> 5,9
// 8,0 -> 0,8
// 9,4 -> 3,4
// 2,2 -> 2,1
// 7,0 -> 7,4
// 6,4 -> 2,0
// 0,9 -> 2,9
// 3,4 -> 1,4
// 0,0 -> 8,8
// 5,5 -> 8,2

// You still need to determine the number of points where at least two lines overlap. In the above example, this is still anywhere in the diagram with a 2 or larger - now a total of 12 points.

func TestOverlaps(t *testing.T) {

	assert.Equal(t, 12, overlaps(lines))

}

var lines = strings.Split(input, "\n")

var input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`
