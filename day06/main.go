package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	data := GuardMap{data: strings.Split(string(input), "\n")}
	width, height := data.getSize()

	uniquePos := map[Coordinate]bool{}

	direction := Up
	guard := data.getGuardPosition()

	for guard.isWithinBoundary(height, width) {

		newPos := guard.move(direction)
		if !newPos.isWithinBoundary(height, width) {
			break
		}

		if data.isObstacle(newPos) {
			direction = (direction + 1) % 4
		} else {
			guard = newPos
			uniquePos[newPos] = true
		}
	}

	fmt.Printf("Visited %d unique positions", len(uniquePos))

}
