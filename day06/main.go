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

	uniquePos := getUniquePos(data)

	fmt.Printf("Guard visited %d unique positions\n", len(uniquePos))

	loopCount := countPositionsForLoops(data, uniquePos)
	fmt.Printf("There is %d positions which results in guard looping", loopCount)
}

func getUniquePos(data GuardMap) map[Coordinate]bool {
	width, height := data.getSize()

	uniquePos := map[Coordinate]bool{}

	direction := Up
	guard := data.getGuardPosition()

	for {
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

	return uniquePos
}

func countPositionsForLoops(data GuardMap, uniquePos map[Coordinate]bool) int {
	width, height := data.getSize()

	posCount := 0

	// Bruteforce solutions based on every visisted positions
	for pos := range uniquePos {
		// Ensure proper clone of map
		newMap := GuardMap{data: make([]string, len(data.data))}
		copy(newMap.data, data.data)
		newMap.setObstacle(pos)

		// Set start params
		direction := Up
		guard := newMap.getGuardPosition()
		i := 0

		// Simulate guard position
		for guard.isWithinBoundary(height, width) {
			newPos := guard.move(direction)
			if newPos.isWithinBoundary(height, width) && newMap.isObstacle(newPos) {
				direction = (direction + 1) % 4
			} else {
				guard = newPos
			}

			// Well consider the map with a loop if we have been running double the length of the non-looped flow
			if i > len(uniquePos)*2 {
				posCount++
				break
			}

			i++
		}
	}

	return posCount
}
