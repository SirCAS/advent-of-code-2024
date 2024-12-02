package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	data, _ := os.ReadFile(filename)

	lines := strings.Split(string(data), "\n")

	col0 := make([]int, len(lines))
	col1 := make([]int, len(lines))

	for idx, line := range lines {
		cols := strings.Split(line, "   ")
		col0[idx], _ = strconv.Atoi(cols[0])
		col1[idx], _ = strconv.Atoi(cols[1])
	}

	slices.Sort(col0)
	slices.Sort(col1)

	totalDistance := 0.0
	similarityScore := 0

	for idxCol0 := 0; idxCol0 < len(lines); idxCol0++ {
		totalDistance += math.Abs(float64(col0[idxCol0]) - float64(col1[idxCol0]))

		occur := 0
		for idxCol1 := 0; idxCol1 < len(lines); idxCol1++ {
			if col0[idxCol0] == col1[idxCol1] {
				occur++
			}
		}

		similarityScore += col0[idxCol0] * occur
	}

	fmt.Printf("  Total distance is : %0.f\n", totalDistance)
	fmt.Printf("Similarity score is : %d\n", similarityScore)
}
