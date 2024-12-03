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
}
