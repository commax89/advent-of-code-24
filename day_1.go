package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readDay1File(isTest bool) (left, right []int) {
	filename := "assets/day_1.test.txt"
	if !isTest {
		filename = "assets/day_1.txt"
	}
	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file handle %s\n", err)
		}
	}(file)
	if err != nil {
		log.Fatalf("Error reading file %s\n", err)
	}

	left = make([]int, 0)
	right = make([]int, 0)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		line = strings.Replace(line, " ", "", 2)
		lineNumbers := strings.Split(line, " ")

		leftVal, err := strconv.Atoi(lineNumbers[0])
		if err != nil {
			log.Printf("skipping line with %s", line)
		}
		rightVal, err := strconv.Atoi(lineNumbers[1])
		if err != nil {
			log.Printf("skipping line with %s", line)
		}
		left = append(left, leftVal)
		right = append(right, rightVal)
	}
	return
}

// Be aware that since slices in go are essentially pointers to the underlying array
// we are basically altering the original slice here, but for this purpose it is ok
func solveDay1(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	var result int

	for idx, leftVal := range left {
		rightVal := right[idx]
		result += max(leftVal, rightVal) - min(leftVal, rightVal)
	}

	return result
}
