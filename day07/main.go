package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Equation struct {
	result  int
	numbers []int
}

func parseEquation(eq string) Equation {
	result := Equation{}

	parts := strings.Split(eq, ": ")
	result.result, _ = strconv.Atoi(parts[0])

	numers := strings.Split(parts[1], " ")

	for _, i := range numers {
		c, _ := strconv.Atoi(i)
		result.numbers = append(result.numbers, c)
	}

	return result
}

type Operator int

const (
	Add Operator = iota
	Multiply
)

func hasSolution(result int, nums []int, ops []Operator) []int {
	var res []int
	for op := range ops {
		var a int
		switch op {
		case int(Add):
			a = result + nums[0]
		case int(Multiply):
			a = result * nums[0]
		}

		if len(nums) > 1 {
			res = append(res, hasSolution(a, nums[1:], ops)...)
		} else {
			res = append(res, a)
		}

	}

	return res
}

func main() {
	input, error := os.ReadFile("input-real.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	resultsSum := 0
	for _, line := range strings.Split(string(input), "\n") {

		eq := parseEquation(line)
		solutions := hasSolution(eq.numbers[0], eq.numbers[1:], []Operator{Add, Multiply})
		if slices.Contains(solutions, eq.result) {
			resultsSum += eq.result
		}

	}

	fmt.Printf("Sum of possible equations are %d\n", resultsSum)

}
