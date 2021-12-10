package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {

	assert.Equal(t, 15, find(strings.Split(input, "\n")))

}

var input = `2199943210
3987894921
9856789892
8767896789
9899965678
`
