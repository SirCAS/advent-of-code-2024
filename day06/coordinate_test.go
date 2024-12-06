package main

import "testing"

type moveTest struct {
	instance Coordinate
	arg1     Direction
	expected Coordinate
}

var moveTests = []moveTest{
	{Coordinate{x: 0, y: 0}, Up, Coordinate{x: 0, y: -1}},
	{Coordinate{x: 0, y: 0}, Down, Coordinate{x: 0, y: 1}},
	{Coordinate{x: 0, y: 0}, Right, Coordinate{x: 1, y: 0}},
	{Coordinate{x: 0, y: 0}, Left, Coordinate{x: -1, y: 0}},
}

func TestMove(t *testing.T) {
	for _, test := range moveTests {
		output := test.instance.move(test.arg1)

		if test.instance.x != 0 || test.instance.y != 0 {
			t.Fatalf("Input %v has changed which is not allowed", test.instance)
		}

		if output != test.expected {
			t.Fatalf("Instance moved %d and was expected at %v but was at %v", test.arg1, test.expected, output)
		}
	}
}

type isWithinBoundaryTest struct {
	pos      Coordinate
	height   int
	width    int
	expected bool
}

var isWithinBoundaryTests = []isWithinBoundaryTest{
	{Coordinate{x: 0, y: 0}, 10, 10, true},
	{Coordinate{x: 0, y: 9}, 10, 10, true},
	{Coordinate{x: 9, y: 0}, 10, 10, true},
	{Coordinate{x: 9, y: 9}, 10, 10, true},
	{Coordinate{x: -1, y: 0}, 10, 10, false},
	{Coordinate{x: 0, y: -1}, 10, 10, false},
	{Coordinate{x: -1, y: -1}, 10, 10, false},
	{Coordinate{x: 10, y: 0}, 10, 10, false},
	{Coordinate{x: 0, y: 10}, 10, 10, false},
	{Coordinate{x: 10, y: 10}, 10, 10, false},
}

func TestIsWithinBoundary(t *testing.T) {
	for _, test := range isWithinBoundaryTests {
		if output := test.pos.isWithinBoundary(test.height, test.width); output != test.expected {
			t.Fatalf("Coordinate at %v in a 10x10 system was expected to be %t but was %t", test.pos, test.expected, output)
		}
	}
}
