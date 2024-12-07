package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	part1, part2 := 0, 0

	for _, line := range strings.Split(string(input), "\n") {

		eq := parseEquation(line)

		solutions1 := findSolutions(eq.numbers[0], eq.numbers[1:], []Operator{Add, Multiply})
		if slices.Contains(solutions1, eq.result) {
			part1 += eq.result
		}

		solutions2 := findSolutions(eq.numbers[0], eq.numbers[1:], []Operator{Add, Multiply, Concat})
		if slices.Contains(solutions2, eq.result) {
			part2 += eq.result
		}
	}

	fmt.Printf("Sum of possible equations using Add, Multiply are %d\n", part1)
	fmt.Printf("Sum of possible equations using Add, Multiply, Concat are %d\n", part2)

}

type Operator int

const (
	Add Operator = iota
	Multiply
	Concat
)

func findSolutions(result int, nums []int, operators []Operator) []int {
	var results []int
	for operator := range operators {
		var interResult int
		switch operator {
		case int(Add):
			interResult = result + nums[0]
		case int(Multiply):
			interResult = result * nums[0]
		case int(Concat):
			interResult, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(nums[0]))
		}

		if len(nums) > 1 {
			results = append(results, findSolutions(interResult, nums[1:], operators)...)
		} else {
			results = append(results, interResult)
		}
	}

	return results
}
