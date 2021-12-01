package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For example, suppose you had the following report:

// 199
// 200
// 208
// 210
// 200
// 207
// 240
// 269
// 260
// 263

// In this example, there are 7 measurements that are larger than the previous measurement.

func TestFindMatch(t *testing.T) {

	depths := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	assert.Equal(t, 7, findIncreases(depths))

}
