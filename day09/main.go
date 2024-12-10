package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	disc, error := os.ReadFile("input1.txt")
	if error != nil {
		fmt.Println(error)
		os.Exit(-1)
	}

	filesystem := getFilesystem(string(disc))
	defragBlocks := defragmentBlocks(filesystem)
	checksum := calculateChecksum(defragBlocks)
	fmt.Printf(" Checksum of defragmented blocks : %d\n", checksum)

}

func getFilesystem(disc string) []int {
	var filesystem []int

	for index := 0; index < len(disc); index++ {
		var val int
		if index%2 == 0 {
			val = index / 2
		} else {
			val = -1
		}

		size, _ := strconv.Atoi(string(disc[index]))
		test := slices.Repeat([]int{val}, size)
		filesystem = append(filesystem, test...)
	}

	return filesystem
}

func defragmentBlocks(filesystem []int) []int {
	for {
		firstSpaceIdx := slices.Index(filesystem, -1)
		clone := slices.Clone(filesystem)
		slices.Reverse(clone)
		lastDigitIdx := len(clone) - 1 - slices.IndexFunc(clone, func(n int) bool {
			return n != -1
		})

		if firstSpaceIdx >= lastDigitIdx {
			return filesystem
		}

		filesystem[firstSpaceIdx] = filesystem[lastDigitIdx]
		filesystem[lastDigitIdx] = -1
	}
}

func calculateChecksum(filesystem []int) int {
	sum := 0

	emptySpaceIdx := slices.Index(filesystem, -1)
	for i := range filesystem[:emptySpaceIdx] {
		sum += i * filesystem[i]
	}

	return sum
}
