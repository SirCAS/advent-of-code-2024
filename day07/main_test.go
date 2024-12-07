package main

import (
	"reflect"
	"slices"
	"testing"
)

type hasSolutionTest struct {
	nums     []int
	ops      []Operator
	expected []int
}

var hasSolutionTests = []hasSolutionTest{
	{[]int{10, 19}, []Operator{Add, Multiply}, []int{190, 29}},
	{[]int{81, 40, 27}, []Operator{Add, Multiply}, []int{87480, 3267, 3267, 148}},
}

func TestHasSolution(t *testing.T) {
	for _, test := range hasSolutionTests {
		outputs := hasSolution(test.nums[0], test.nums[1:], test.ops)
		slices.Sort(outputs)
		slices.Sort(test.expected)
		if !reflect.DeepEqual(outputs, test.expected) {
			t.Fatalf("Input of %v yielded %v but should be %v", test.nums, outputs, test.expected)
		}
	}
}
