package main

import "strings"

type GuardMap struct {
	data []string
}

func (m *GuardMap) isObstacle(pos Coordinate) bool {
	return string(m.data[pos.y][pos.x]) == "#"
}

func (m *GuardMap) getSize() (int, int) {
	return len(m.data[0]), len(m.data)
}

func (m *GuardMap) getGuardPosition() Coordinate {
	for y, line := range m.data {
		idx := strings.Index(line, "^")
		if idx != -1 {
			return Coordinate{x: idx, y: y}
		}
	}
	return Coordinate{}
}