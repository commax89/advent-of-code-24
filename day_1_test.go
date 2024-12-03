package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveDay1(t *testing.T) {
	result := solveDay1([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})
	assert.Equal(t, 11, result)
}

func TestCountOccurrences(t *testing.T) {
	result := countOccurrences(3, []int{3, 4, 2, 1, 3, 3})
	assert.Equal(t, 3, result)

	result = countOccurrences(4, []int{3, 4, 4, 2, 1, 3, 3})
	assert.Equal(t, 2, result)

	result = countOccurrences(5, []int{3, 4, 4, 2, 1, 3, 3})
	assert.Equal(t, 0, result)
}

func TestSimilarityScore(t *testing.T) {
	result := similarityScore([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})
	assert.Equal(t, 31, result)
}
