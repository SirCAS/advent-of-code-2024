package main

import (
	"testing"
)

type findXmasTest struct {
	arg1     []string
	expected int
}

var findXmasTests = []findXmasTest{
	{[]string{".XMAS.", "......", "......", "......", "......", "......"}, 1},
	{[]string{".SAMX.", "......", "......", "......", "......", "......"}, 1},
	{[]string{".S....", ".A....", ".M....", ".X....", "......", "......"}, 1},
	{[]string{".X....", ".M....", ".A....", ".S....", "......", "......"}, 1},
	{[]string{"....X.", "...M..", "..A...", ".S....", "......", "......"}, 1},
	{[]string{".X....", "..M...", "...A..", "....S.", "......", "......"}, 1},
}

func TestFindXmas(t *testing.T) {
	for _, test := range findXmasTests {
		if output := findXmas(test.arg1); output != test.expected {
			t.Fatalf("Input %v has count %d and expected %d", test.arg1, output, test.expected)
		}
	}
}
