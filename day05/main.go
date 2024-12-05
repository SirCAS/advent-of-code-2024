package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	// input file is in two sections divided by two new-line chars
	inputParts := strings.Split(string(input), "\n\n")

	rules := parseRules(inputParts[0]) // first is rules in file
	pages := parsePages(inputParts[1]) // .. then follows pages

	middlepageNumberValidSum := 0
	middlepageNumberFixedSum := 0

	for _, page := range pages {

		if validateUpdate(page, rules) {

			middlepageNumberValidSum += findMiddlepageNumber(page)
		} else {
			fixed := fixUpdate(page, rules)
			middlepageNumberFixedSum += findMiddlepageNumber(fixed)
		}
	}

	fmt.Printf("The sum of valid middlepage nubmers are: %d\n", middlepageNumberValidSum)
	fmt.Printf("The sum of fixed middlepage nubmers are: %d\n", middlepageNumberFixedSum)

}

func fixUpdate(page []int, rules []sortRule) []int {
	for j := 0; j < len(rules); j++ {

		before := slices.Index(page, rules[j].before)
		after := slices.Index(page, rules[j].after)

		if before != -1 && after != -1 && before > after {
			// Swap numbers then try again
			page[before], page[after] = page[after], page[before]
			j = 0
		}
	}

	return page
}

func validateUpdate(page []int, rules []sortRule) bool {
	for _, rule := range rules {
		before := slices.Index(page, rule.before)
		after := slices.Index(page, rule.after)
		if before != -1 && after != -1 && before > after {
			return false
		}
	}

	return true
}

func findMiddlepageNumber(pages []int) int {
	size := len(pages)
	middleIndex := size / 2
	return pages[middleIndex]
}
