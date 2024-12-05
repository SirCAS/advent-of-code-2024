package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	inputRules, error := os.ReadFile("input-rules-real.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	inputPages, error := os.ReadFile("input-pages-real.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	rules := parseRules(string(inputRules))
	pages := parsePages(string(inputPages))

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
			// Swap
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
