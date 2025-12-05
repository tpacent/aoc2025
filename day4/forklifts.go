package day4

import (
	"bufio"
	"io"
)

type Floor map[Coord]struct{}

type Coord struct {
	X, Y uint8
}

func RemoveMovables(floor Floor, limit int) (total int) {
	for {
		removed := RemoveMovablesStep(floor, limit)
		total += removed
		if removed == 0 {
			return
		}
	}
}

func RemoveMovablesStep(floor Floor, limit int) int {
	toRemove := []Coord{}

	for coord := range floor {
		if count := Around(floor, coord); count < limit {
			toRemove = append(toRemove, coord)
		}
	}

	for _, coords := range toRemove {
		delete(floor, coords)
	}

	return len(toRemove)
}

var deltas = []func(Coord) Coord{
	func(c Coord) Coord { return Coord{c.X - 1, c.Y - 1} },
	func(c Coord) Coord { return Coord{c.X, c.Y - 1} },
	func(c Coord) Coord { return Coord{c.X + 1, c.Y - 1} },
	func(c Coord) Coord { return Coord{c.X - 1, c.Y} },
	func(c Coord) Coord { return Coord{c.X + 1, c.Y} },
	func(c Coord) Coord { return Coord{c.X - 1, c.Y + 1} },
	func(c Coord) Coord { return Coord{c.X, c.Y + 1} },
	func(c Coord) Coord { return Coord{c.X + 1, c.Y + 1} },
}

func Around(floor Floor, coord Coord) (count int) {
	for _, dfunc := range deltas {
		if _, ok := floor[dfunc(coord)]; ok {
			count++
		}
	}
	return
}

func ParseInput(r io.Reader) Floor {
	floor := make(Floor)
	var y uint8

	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		for x, char := range line {
			if char == '@' {
				floor[Coord{X: uint8(x), Y: y}] = struct{}{}
			}
		}

		y++
	}

	return floor
}
