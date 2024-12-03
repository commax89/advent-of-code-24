package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsReportSafe(t *testing.T) {
	var res bool
	res = isReportSafe([]int{7, 6, 4, 2, 1})
	assert.True(t, res, "7, 6, 4, 2, 1")

	res = isReportSafe([]int{1, 2, 7, 8, 9})
	assert.False(t, res, "1, 2, 7, 8, 9")

	res = isReportSafe([]int{9, 7, 6, 2, 1})
	assert.False(t, res, "9, 7, 6, 2, 1")

	res = isReportSafe([]int{1, 3, 2, 4, 5})
	assert.False(t, res, "1, 3, 2, 4, 5")

	res = isReportSafe([]int{8, 6, 4, 4, 1})
	assert.False(t, res, "8, 6, 4, 4, 1")

	res = isReportSafe([]int{1, 3, 6, 7, 9})
	assert.True(t, res, "1, 3, 6, 7, 9")

	return
}

func TestIsReportSafeDampened(t *testing.T) {
	var res bool
	res = isReportSafeDampened([]int{7, 6, 4, 2, 1})
	assert.True(t, res, "7, 6, 4, 2, 1")

	res = isReportSafeDampened([]int{1, 2, 7, 8, 9})
	assert.False(t, res, "1, 2, 7, 8, 9")

	res = isReportSafeDampened([]int{9, 7, 6, 2, 1})
	assert.False(t, res, "9, 7, 6, 2, 1")

	res = isReportSafeDampened([]int{1, 3, 2, 4, 5})
	assert.True(t, res, "1, 3, 2, 4, 5")

	res = isReportSafeDampened([]int{8, 6, 4, 4, 1})
	assert.True(t, res, "8, 6, 4, 4, 1")

	res = isReportSafeDampened([]int{1, 3, 6, 7, 9})
	assert.True(t, res, "1, 3, 6, 7, 9")

	return
}
