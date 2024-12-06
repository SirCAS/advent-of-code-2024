package main

import (
	"fmt"
	"os"
	"strings"
)

// type Direction int

// const (
// 	Up Direction = iota
// 	Down
// 	Left
// 	Right
// )

// type coordinate struct {
// 	x int
// 	y int
// }

// func (c *coordinate) move(d Direction) int {

// 	return r.width * r.height
// }

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	data := strings.Split(string(input), "\n")

	guardX, guardY := getGuardCoordiates(data)

	height := len(data) - 1
	width := len(data[0]) - 1

	var pos []coordinate

	direction := Up

	for !(0 >= guardY || guardY >= height || 0 >= guardX || guardX >= width) {

		// take step
		if direction == Up {
			if string(data[guardY-1][guardX]) != "#" {
				guardY--
			} else {
				direction = Right
			}
		} else if direction == Down {
			if string(data[guardY+1][guardX]) != "#" {
				guardY++
			} else {
				direction = Left
			}
		} else if direction == Left {
			if string(data[guardY][guardX-1]) != "#" {
				guardX--
			} else {
				direction = Up
			}
		} else if direction == Right {
			if string(data[guardY][guardX+1]) != "#" {
				guardX++
			} else {
				direction = Down
			}
		}
		pos = append(pos, coordinate{x: guardX, y: guardY})
		fmt.Printf("x: %d, y: %d\n", guardX, guardY)
	}

	fmt.Printf("Visisted %d positions", len(pos))

	unique := map[coordinate]bool{}
	for _, v := range pos {
		unique[v] = true
	}

	fmt.Printf("Visisted %d uniqe positions", len(unique))

}

func getGuardCoordiates(data []string) (int, int) {
	for y, line := range data {
		idx := strings.Index(line, "^")
		if idx != -1 {
			return idx, y
		}
	}
	return 0, 0
}
