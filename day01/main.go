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
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

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

	for idx := 0; idx < len(lines); idx++ {
		totalDistance += math.Abs(float64(col0[idx]) - float64(col1[idx]))
	}

	fmt.Printf("Distance is %0.f", totalDistance)
}
