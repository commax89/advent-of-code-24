package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readDay2File(isTest bool) (reports [][]int) {
	filename := "assets/day_2.test.txt"
	if !isTest {
		filename = "assets/day_2.txt"
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

	reports = make([][]int, 0)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		lineNumbers := make([]int, 0)
		for _, str := range strings.Split(line, " ") {
			n, err := strconv.Atoi(str)
			if err != nil {
				lineNumbers = append(lineNumbers, 0)
				continue
			}
			lineNumbers = append(lineNumbers, n)
		}

		reports = append(reports, lineNumbers)
	}
	return
}

func isReportSafe(r []int) (safe bool) {
	safe = true

	for i, v := range r {
		if i == 0 {
			continue
		}

		if r[i-1] == v {
			safe = false
			break
		}

		d := max(r[i-1], v) - min(r[i-1], v)
		if d > 3 {
			safe = false
			break
		}

		if i+1 == len(r) {
			continue
		}

		if math.Signbit(float64(r[i-1]-v)) != math.Signbit(float64(v-r[i+1])) {
			safe = false
			break
		}
	}

	return
}

func isReportSafeDampened(r []int) (safe bool) {
	safe = true
	levelToRemove := -1

	for i, v := range r {
		if i == 0 {
			continue
		}

		if r[i-1] == v {
			safe = false
			levelToRemove = i
			break
		}

		d := max(r[i-1], v) - min(r[i-1], v)
		if d > 3 {
			safe = false
			levelToRemove = i
			break
		}

		if i+1 == len(r) {
			continue
		}

		if math.Signbit(float64(r[i-1]-v)) != math.Signbit(float64(v-r[i+1])) {
			safe = false
			levelToRemove = i
			break
		}
	}

	if levelToRemove != -1 {
		c := make([]int, len(r))
		copy(c, r)
		c = append(c[:levelToRemove], c[levelToRemove+1:]...)
		return isReportSafe(c)
	}

	return
}
