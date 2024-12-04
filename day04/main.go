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

	occurrences := findXmas(data)

	fmt.Printf("there is %d\n", occurrences)
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
