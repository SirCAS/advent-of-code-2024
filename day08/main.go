package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type coordinate struct {
	x int
	y int
}

func main() {
	input, error := os.ReadFile("input.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}
	lines := strings.Split(string(input), "\n")

	width, height := len(lines[0]), len(lines)

	antennaGroup := map[rune][]coordinate{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			field := rune(lines[y][x])
			if unicode.IsDigit(field) || unicode.IsLetter(field) {
				antenna := coordinate{
					x: x,
					y: y,
				}

				antennaGroup[field] = append(antennaGroup[field], antenna)
			}
		}
	}

	antinode := map[coordinate]bool{}

	for _, antennas := range antennaGroup {
		for i := range antennas {
			for j := range antennas {
				if i != j {

					coordinate := coordinate{
						x: antennas[i].x + int(antennas[i].x-antennas[j].x),
						y: antennas[i].y + int(antennas[i].y-antennas[j].y),
					}

					if 0 <= coordinate.x && coordinate.x < width && 0 <= coordinate.y && coordinate.y < height {
						antinode[coordinate] = true
					}
				}
			}
		}
	}

	fmt.Printf("There is %d", len(antinode))
}
