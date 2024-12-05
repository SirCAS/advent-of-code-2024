package main

import (
	"reflect"
	"testing"
)

type findMiddlepageNumberTest struct {
	arg1     []int
	expected int
}

var findMiddlepageNumberTests = []findMiddlepageNumberTest{
	{[]int{75, 47, 61, 53, 29}, 61},
	{[]int{97, 61, 53, 29, 13}, 53},
	{[]int{75, 29, 13}, 29},
}

func TestFindMiddlepageNumber(t *testing.T) {
	for _, test := range findMiddlepageNumberTests {
		if output := findMiddlepageNumber(test.arg1); output != test.expected {
			t.Fatalf("Input %v has count %d and expected %d", test.arg1, output, test.expected)
		}
	}
}

type validateUpdateTest struct {
	arg1     []int
	expected bool
}

var validateUpdateTests = []validateUpdateTest{
	{[]int{75, 47, 61, 53, 29}, true},
	{[]int{97, 61, 53, 29, 13}, true},
	{[]int{75, 29, 13}, true},
	{[]int{75, 97, 47, 61, 53}, false},
	{[]int{61, 13, 29}, false},
	{[]int{97, 13, 75, 29, 47}, false},
}

func TestValidateUpdate(t *testing.T) {
	rules := []sortRule{
		{before: 47, after: 53},
		{before: 97, after: 13},
		{before: 97, after: 61},
		{before: 97, after: 47},
		{before: 75, after: 29},
		{before: 61, after: 13},
		{before: 75, after: 53},
		{before: 29, after: 13},
		{before: 97, after: 29},
		{before: 53, after: 29},
		{before: 61, after: 53},
		{before: 97, after: 53},
		{before: 61, after: 29},
		{before: 47, after: 13},
		{before: 75, after: 47},
		{before: 97, after: 75},
		{before: 47, after: 61},
		{before: 75, after: 61},
		{before: 47, after: 29},
		{before: 75, after: 13},
		{before: 53, after: 13},
	}

	for _, test := range validateUpdateTests {
		if output := validateUpdate(test.arg1, rules); output != test.expected {
			t.Fatalf("Pages %v validated %t but expected was %t", test.arg1, output, test.expected)
		}
	}
}

type fixUpdateTest struct {
	arg1     []int
	expected []int
}

var fixUpdateTests = []fixUpdateTest{
	{[]int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
	{[]int{61, 13, 29}, []int{61, 29, 13}},
	{[]int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
}

func TestFixUpdate(t *testing.T) {
	rules := []sortRule{
		{before: 47, after: 53},
		{before: 97, after: 13},
		{before: 97, after: 61},
		{before: 97, after: 47},
		{before: 75, after: 29},
		{before: 61, after: 13},
		{before: 75, after: 53},
		{before: 29, after: 13},
		{before: 97, after: 29},
		{before: 53, after: 29},
		{before: 61, after: 53},
		{before: 97, after: 53},
		{before: 61, after: 29},
		{before: 47, after: 13},
		{before: 75, after: 47},
		{before: 97, after: 75},
		{before: 47, after: 61},
		{before: 75, after: 61},
		{before: 47, after: 29},
		{before: 75, after: 13},
		{before: 53, after: 13},
	}

	for _, test := range fixUpdateTests {
		if output := fixUpdate(test.arg1, rules); !reflect.DeepEqual(output, test.expected) {
			t.Fatalf("Unsorted pages %v was expected to be %v but was infact %v", test.arg1, output, test.expected)
		}
	}
}
