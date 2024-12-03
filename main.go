package main

import (
	"fmt"
)

func main() {
	left, right := readDay1File(false)

	dayOneResult := solveDay1(left, right)
	fmt.Printf("the day one, part one, result is: %d \n", dayOneResult)

	dayOneResult2 := similarityScore(left, right)
	fmt.Printf("the day one, part two, result is: %d \n", dayOneResult2)

	reports := readDay2File(false)
	safeCount := 0
	for _, r := range reports {
		safe := isReportSafe(r)
		if safe {
			safeCount++
		}
	}

	fmt.Printf("the day two, part one, result is: %d\n", safeCount)

	safeCountDampened := 0
	for _, r := range reports {
		safe := isReportSafeDampened(r)
		if safe {
			safeCountDampened++
		}
	}
	fmt.Printf("the day two, part two, result is: %d\n", safeCountDampened)
}
