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

	middlepageNumberSum := 0

	for _, page := range pages {

		if validateUpdate(page, rules) {

			middlepageNumberSum += findMiddlepageNumber(page)
		}
	}

	fmt.Printf("The sum of valid middlepage nubmers are: %d", middlepageNumberSum)

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
