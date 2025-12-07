package day7

import (
	"bufio"
	"io"
)

type Manifold map[Coord]struct{}

type Coord struct {
	X, Y int16
}

func CountAll(mf Manifold, beam Coord, rowmax int16) (splits, total int) {
	row := beam.Y
	beams := map[int16]int{beam.X: 1}

	for {
		if row++; row > rowmax {
			break
		}

		for x, v := range beams {
			if _, ok := mf[Coord{x, row}]; ok {
				splits++
				beams[x-1] += v
				beams[x+1] += v
				delete(beams, x)
			}
		}
	}

	for _, v := range beams {
		total += v
	}

	return
}

func ParseInput(r io.Reader) (mf Manifold, start Coord, row int16) {
	mf = make(Manifold)
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		row++

		for col, char := range line {
			switch char {
			case '^':
				mf[Coord{int16(col), row}] = struct{}{}
			case 'S':
				start = Coord{int16(col), row}
			}
		}
	}

	return
}
