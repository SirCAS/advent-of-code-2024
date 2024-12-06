package main

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) move(d Direction) Coordinate {
	newCoord := c
	switch d {
	case Up:
		newCoord.y--
	case Down:
		newCoord.y++
	case Right:
		newCoord.x++
	case Left:
		newCoord.x--
	}
	return newCoord
}

func (c *Coordinate) isWithinBoundary(height, width int) bool {
	// Taking into account that coordinates are zero-indexed
	return 0 <= c.y && c.y <= height-1 &&
		0 <= c.x && c.x <= width-1

}
