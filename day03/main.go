package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	filename := "input.txt"
	input, _ := os.ReadFile(filename)

	r := regexp.MustCompile(`mul\((?P<first>\d+),(?P<second>\d+)\)`)
	matches := r.FindAllStringSubmatch(string(input), -1)

	sum := 0

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}

	fmt.Println(sum)

}
