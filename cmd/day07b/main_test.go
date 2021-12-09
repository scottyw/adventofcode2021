package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T) {

	assert.Equal(t, 168, fuel([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))

}
