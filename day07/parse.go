package main

import (
	"strconv"
	"strings"
)

type Equation struct {
	result  int
	numbers []int
}

func parseEquation(eq string) Equation {
	parsedEq := Equation{}

	parts := strings.Split(eq, ": ")
	parsedEq.result, _ = strconv.Atoi(parts[0])

	numbers := strings.Split(parts[1], " ")

	for _, numberText := range numbers {
		number, _ := strconv.Atoi(numberText)
		parsedEq.numbers = append(parsedEq.numbers, number)
	}

	return parsedEq
}
