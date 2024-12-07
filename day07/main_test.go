package main

import (
	"reflect"
	"slices"
	"testing"
)

type findSolutionsTest struct {
	nums     []int
	ops      []Operator
	expected []int
}

var findSolutionsTests = []findSolutionsTest{
	{[]int{10, 19}, []Operator{Add, Multiply}, []int{190, 29}},
	{[]int{81, 40, 27}, []Operator{Add, Multiply}, []int{87480, 3267, 3267, 148}},
	{[]int{10, 19}, []Operator{Add, Multiply, Concat}, []int{29, 190, 1019}},
	{[]int{81, 40, 27}, []Operator{Add, Multiply, Concat}, []int{148, 3267, 3267, 8167, 12127, 87480, 219780, 324027, 814027}},
}

func TestFindSolutions(t *testing.T) {
	for _, test := range findSolutionsTests {
		outputs := findSolutions(test.nums[0], test.nums[1:], test.ops)
		slices.Sort(outputs)
		slices.Sort(test.expected)
		if !reflect.DeepEqual(outputs, test.expected) {
			t.Fatalf("Input of %v yielded %v but should be %v", test.nums, outputs, test.expected)
		}
	}
}
