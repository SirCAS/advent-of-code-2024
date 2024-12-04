package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	input, _ := os.ReadFile(filename)

	data := strings.Split(string(input), "\n")

	xmasOccurrences := findXmas(data)
	fmt.Printf("There is %d xmas'es\n", xmasOccurrences)

	crossMasOccurrences := findCrossMas(data)
	fmt.Printf("There is %d cross-mas'es\n", crossMasOccurrences)
}

func findXmas(data []string) int {
	occurrences := 0

	width := len(data[0])
	height := len(data)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// horitontal
			if x < width-3 && (data[y][x:x+4] == "XMAS" || data[y][x:x+4] == "SAMX") {
				occurrences++
			}

			// vertical
			if y < height-3 && data[y][x] == 'X' && data[y+1][x] == 'M' && data[y+2][x] == 'A' && data[y+3][x] == 'S' {
				occurrences++
			}
			if y < height-3 && data[y][x] == 'S' && data[y+1][x] == 'A' && data[y+2][x] == 'M' && data[y+3][x] == 'X' {
				occurrences++
			}

			// cross
			if x < width-3 && y < height-3 {
				if data[y][x] == 'X' && data[y+1][x+1] == 'M' && data[y+2][x+2] == 'A' && data[y+3][x+3] == 'S' {
					occurrences++
				}
				if data[y][x] == 'S' && data[y+1][x+1] == 'A' && data[y+2][x+2] == 'M' && data[y+3][x+3] == 'X' {
					occurrences++
				}
			}

			if x >= 3 && y < height-3 {
				if data[y][x] == 'X' && data[y+1][x-1] == 'M' && data[y+2][x-2] == 'A' && data[y+3][x-3] == 'S' {
					occurrences++
				}
				if data[y][x] == 'S' && data[y+1][x-1] == 'A' && data[y+2][x-2] == 'M' && data[y+3][x-3] == 'X' {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func findCrossMas(data []string) int {
	occurrences := 0

	width := len(data[0])
	height := len(data)

	for y := 0; y < height-2; y++ {
		for x := 0; x < width-2; x++ {

			diagonal1 := data[y][x] == 'M' && data[y+1][x+1] == 'A' && data[y+2][x+2] == 'S'
			diagonal1_reverse := data[y][x] == 'S' && data[y+1][x+1] == 'A' && data[y+2][x+2] == 'M'

			diagonal2 := data[y][x+2] == 'M' && data[y+1][x+1] == 'A' && data[y+2][x] == 'S'
			diagonal2_reverse := data[y][x+2] == 'S' && data[y+1][x+1] == 'A' && data[y+2][x] == 'M'

			if (diagonal1 || diagonal1_reverse) && (diagonal2 || diagonal2_reverse) {
				occurrences++
			}

		}
	}

	return occurrences
}
