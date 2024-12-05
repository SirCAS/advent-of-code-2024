package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parsePages(pageText string) [][]int {
	lines := strings.Split(pageText, "\n")

	var pages [][]int

	for _, line := range lines {
		page, errors := parsePage(line)
		if len(errors) > 0 {
			fmt.Printf("Warning: couldn't parse page - skipping it")
		} else {
			pages = append(pages, page)
		}
	}

	return pages
}

func parsePage(page string) ([]int, []error) {
	numbersText := strings.Split(page, ",")

	var pages []int
	var errors []error

	for _, numberText := range numbersText {
		number, error := strconv.Atoi(numberText)
		pages = append(pages, number)
		if error != nil {
			errors = append(errors, error)
		}
	}

	return pages, errors
}
