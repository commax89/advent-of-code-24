package main

import (
	"slices"
)

// Be aware that since slices in go are essentially pointers to the underlying array
// we are basically altering the original slice here, but for this purpose it is ok
func solveDay1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	//fmt.Printf("what is left? %+v", left)

	var result int

	for idx, leftVal := range left {
		rightVal := right[idx]
		result += leftVal - rightVal
	}

	return result
}
