package main

import (
	"fmt"
)

func main() {
	left, right := readDay1File(false)

	dayOneResult := solveDay1(left, right)
	fmt.Printf("the day one result is: %d \n", dayOneResult)
}
