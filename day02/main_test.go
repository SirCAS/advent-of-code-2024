package main

import (
	"testing"
)

type validateTest struct {
	arg1     string
	expected bool
}

var validateTests = []validateTest{
	{"7 6 4 2 1", true},
	{"1 2 7 8 9", false},
	{"9 7 6 2 1", false},
	{"1 3 2 4 5", false},
	{"8 6 4 4 1", false},
	{"1 3 6 7 9", true},
}

var validateWithProblemDampenerTests = []validateTest{
	{"7 6 4 2 1 ", true},
	{"1 2 7 8 9", false},
	{"9 7 6 2 1", false},
	{"1 3 2 4 5", true},
	{"8 6 4 4 1", true},
	{"1 3 6 7 9", true},
	{"59 57 54 57 59 62", false},
	{"6 7 10 13 14 15 17 19", true},
	{"84 83 83 86 83", false},
	{"30 27 27 30 30", false},
	{"36 35 36 38 41 43 47 53", false},
	{"36 35 36 38 41 43 47 50", false},
	{"62 62 65 66 68 69 71 75", false},
}

func TestValidate(t *testing.T) {
	for _, test := range validateTests {
		if output := validate(test.arg1); output != test.expected {
			t.Fatalf("Plan %s has safety equal to %t and expected %t", test.arg1, output, test.expected)
		}
	}
}

func TestValidateWithProblemDampener(t *testing.T) {
	for _, test := range validateWithProblemDampenerTests {
		if output := validateWithProblemDampener(test.arg1); output != test.expected {
			t.Fatalf("Plan %s has safety equal to %t and expected %t", test.arg1, output, test.expected)
		}
	}
}
