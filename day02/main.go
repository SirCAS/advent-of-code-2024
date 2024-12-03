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

func convertReportToLevels(report string) []int {
	levels := strings.Split(report, " ")
	data := make([]int, len(levels))

	for l, level := range levels {
		data[l], _ = strconv.Atoi(level)
	}
	return data
}

func validate(report string) bool {

	levels := convertReportToLevels(report)

	isIncreasing := levels[0] < levels[1]
	for l := 0; l < len(levels)-1; l++ {

		step := levels[l] - levels[l+1]
		if step < -3 || step > 3 {
			return false

		}

		if levels[l] == levels[l+1] {
			return false
		}

		if isIncreasing {
			if levels[l] > levels[l+1] {
				return false

			}
		} else {
			if levels[l] < levels[l+1] {
				return false

			}
		}
	}

	return true
}

func validateWithProblemDampener(report string) bool {
	// old rules still apply - short circut validation if possible
	if validate(report) {
		return true
	}

	// allow a second chance
	levels := convertReportToLevels(report)

	for l := range levels {
		// copy report for mutation
		mutatedReport := make([]int, len(levels))
		copy(mutatedReport, levels)

		// remove l'th element from slice and convert it to string
		mutatedReport = append(mutatedReport[:l], mutatedReport[l+1:]...)
		reportStr := strings.Trim(fmt.Sprint(mutatedReport), "[]")

		// if safe, then we don't need to keep mutating
		if validate(reportStr) {
			return true
		}
	}

	// no luck - still a unsafe report
	return false
}
