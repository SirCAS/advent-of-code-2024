package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	input, _ := os.ReadFile(filename)

	reports := strings.Split(string(input), "\n")

	// Part 1
	countSafeReports := 0
	for _, report := range reports {
		if validate(report) {
			countSafeReports++
		}
	}
	fmt.Printf("There is %d safe reports\n", countSafeReports)

	// Part 2
	countSafeReports = 0
	for _, report := range reports {
		if validateWithProblemDampener(report) {
			countSafeReports++
		}
	}
	fmt.Printf("There is %d safe reports when tolerating one bad level\n", countSafeReports)
}

func validate(report string) bool {
	levels := strings.Split(report, " ")
	data := make([]int, len(levels))

	for l, level := range levels {
		data[l], _ = strconv.Atoi(level)
	}

	isIncreasing := data[0] < data[1]
	for l := 0; l < len(levels)-1; l++ {

		step := data[l] - data[l+1]
		if step < -3 || step > 3 {
			return false

		}

		if data[l] == data[l+1] {
			return false
		}

		if isIncreasing {
			if data[l] > data[l+1] {
				return false

			}
		} else {
			if data[l] < data[l+1] {
				return false

			}
		}
	}

	return true
}

func validateWithProblemDampener(report string) bool {

	levels := strings.Split(report, " ")
	data := make([]int, len(levels))

	for l, level := range levels {
		data[l], _ = strconv.Atoi(level)
	}

	tolerate := 1

	for l := 0; l < len(data)-1; l++ {

		// handle repeatitions
		if data[l] == data[l+1] {
			if tolerate > 0 {
				// remove left side or skip
				data = remove(data, l)
				l--
				tolerate--
				continue
			}
			return false

		}

		// handle too large steps
		step := data[l] - data[l+1]
		if step < -3 || step > 3 {
			if tolerate > 0 {
				data = remove(data, l+1)
				l--
				tolerate--
				continue
			}
			return false
		}

		// skip on first run - we'll validate in subsequent checks
		if l > 0 {
			if data[l-1] > data[l] && data[l] < data[l+1] {
				if tolerate > 0 {
					data = remove(data, l+1)
					l--
					tolerate--
					continue
				}
				return false
			}

			if data[l-1] < data[l] && data[l] > data[l+1] {
				if tolerate > 0 {
					data = remove(data, l)
					l--
					tolerate--
					continue
				}
				return false

			}
		}
	}

	return true

}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
