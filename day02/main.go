package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	input, _ := os.ReadFile(filename)

	reports := strings.Split(string(input), "\n")
	data := make([][]int, len(reports))

	countSafeReports := 0

	for r, report := range reports {

		levels := strings.Split(report, " ")
		data[r] = make([]int, len(levels))

		for l, level := range levels {
			data[r][l], _ = strconv.Atoi(level)
		}

		isSafe := true
		isIncreasing := data[r][0] < data[r][1]
		for l := 0; l < len(levels)-2; l++ {
			delta1 := math.Abs(float64(data[r][l]) - float64(data[r][l+1]))
			delta2 := math.Abs(float64(data[r][l+1]) - float64(data[r][l+2]))

			if !(0 < delta1 && delta1 < 4) || !(0 < delta2 && delta2 < 4) {
				isSafe = false
				continue
			}

			if isIncreasing {
				if !(data[r][l] < data[r][l+1] && data[r][l+1] < data[r][l+2]) {
					isSafe = false
					continue
				}
			} else {
				if !(data[r][l] > data[r][l+1] && data[r][l+1] > data[r][l+2]) {
					isSafe = false
					continue
				}
			}
		}

		if isSafe {
			countSafeReports++
		}
	}

	fmt.Printf("There is : %d safe reports\n", countSafeReports)
}
