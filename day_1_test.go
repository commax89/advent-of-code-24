package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay1(t *testing.T) {
	result := solveDay1([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 4, 9, 3})
	assert.Equal(t, 11, result)
}
