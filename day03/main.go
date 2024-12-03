package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	input, _ := os.ReadFile(filename)

	fmt.Printf("Sum of multiplications are: %d\n", part1(string(input)))
	fmt.Printf("Sum of multiplications are: %d\n", part2(string(input)))
}

func part1(memory string) int {

	r := regexp.MustCompile(`mul\((?P<a>\d+),(?P<b>\d+)\)`)
	matches := r.FindAllStringSubmatch(memory, -1)

	sum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		sum += a * b
	}

	return sum
}

func part2(memory string) int {

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
	matches := r.FindAllString(memory, -1)

	sum := 0
	allowed := true
	for _, match := range matches {

		if match == "do()" {
			allowed = true
		} else if match == "don't()" {
			allowed = false
		} else {
			if allowed {
				numbers := strings.Split(match[4:len(match)-1], ",")

				a, _ := strconv.Atoi(numbers[0])
				b, _ := strconv.Atoi(numbers[1])

				sum += a * b
			}
		}
	}

	return sum
}
