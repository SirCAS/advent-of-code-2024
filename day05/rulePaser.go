package main

import (
	"fmt"
	"strconv"
	"strings"
)

type sortRule struct {
	before int
	after  int
}

func parseRules(ruleText string) []sortRule {
	lines := strings.Split(ruleText, "\n")

	var rules []sortRule

	for _, line := range lines {
		rule, errors := parseRule(line)
		if len(errors) > 0 {
			fmt.Printf("Warning: couldn't parse rule - skipping it")
		} else {
			rules = append(rules, rule)
		}
	}

	return rules
}

func parseRule(rule string) (sortRule, []error) {
	parts := strings.Split(rule, "|")

	var errors []error

	before, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		errors = append(errors, err1)
	}
	after, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		errors = append(errors, err2)
	}

	return sortRule{
		before: before,
		after:  after,
	}, errors
}
