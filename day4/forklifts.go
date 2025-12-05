package day4

import (
	"bufio"
	"io"
)

type Floor map[Coord]struct{}

type Coord [2]uint8

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

var deltas = []Coord{
	// adding 255 causes uint8 to wraparound and subtract 1
	{255, 255}, {0, 255}, {1, 255},
	{255, 0} /*     */, {1, 0},
	{255, 1}, {0, 1}, {1, 1},
}

func Around(floor Floor, c Coord) (count int) {
	for _, d := range deltas {
		item := Coord{c[0] + d[0], c[1] + d[1]}
		if _, ok := floor[item]; ok {
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
				floor[Coord{uint8(x), y}] = struct{}{}
			}
		}

		y++
	}

	return floor
}
