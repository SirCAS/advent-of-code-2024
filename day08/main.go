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

	canvas := len(lines)

	groups := locateAntennasGroupedByFrequency(lines, canvas)

	antinodes := locateAntinodes(groups, canvas)
	fmt.Printf("There is %d antinodes\n", len(antinodes))

	antinodesWithHarmonics := locateAntinodesWithHarmonics(groups, canvas)
	fmt.Printf("There is %d antinodes when considering harmonics too\n", len(antinodesWithHarmonics))
}

func locateAntennasGroupedByFrequency(lines []string, canvas int) map[rune][]coordinate {
	groups := map[rune][]coordinate{}

	for y := 0; y < canvas; y++ {
		for x := 0; x < canvas; x++ {

			field := rune(lines[y][x])
			if unicode.IsDigit(field) || unicode.IsLetter(field) {
				groups[field] = append(groups[field], coordinate{x: x, y: y})
			}
		}
	}

	return groups
}

func locateAntinodes(antennaGroup map[rune][]coordinate, canvas int) map[coordinate]bool {
	antinodes := map[coordinate]bool{}

	for _, antennas := range antennaGroup {
		for i := range antennas {
			for j := range antennas {
				// Skip pair with own location
				if i != j {
					coordinate := coordinate{
						x: antennas[i].x + (antennas[i].x - antennas[j].x),
						y: antennas[i].y + (antennas[i].y - antennas[j].y),
					}

					// Add to results
					if isWithinCanvas(coordinate, canvas) {
						antinodes[coordinate] = true
					}
				}
			}
		}
	}

	return antinodes
}

func locateAntinodesWithHarmonics(antennaGroup map[rune][]coordinate, canvas int) map[coordinate]bool {
	antinodes := map[coordinate]bool{}
	for _, antennas := range antennaGroup {
		for i := range antennas {
			for j := range antennas {
				// "Worst case" would be every possible position from left top corner to bottom
				// right corner (we could use pyhagotas but for the sake of simplicity assume 2x)
				for a := 1; a < canvas*2; a++ {
					coordinate := coordinate{
						x: antennas[i].x + (antennas[i].x-antennas[j].x)*a,
						y: antennas[i].y + (antennas[i].y-antennas[j].y)*a,
					}

					if isWithinCanvas(coordinate, canvas) {
						antinodes[coordinate] = true
					}
				}
			}
		}
	}

	return antinodes
}

func isWithinCanvas(position coordinate, canvas int) bool {
	return 0 <= position.x && position.x < canvas && 0 <= position.y && position.y < canvas
}
