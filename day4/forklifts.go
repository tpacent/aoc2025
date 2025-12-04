package day4

import (
	"bufio"
	"io"
)

func CountMovables(floor *Floor[byte], limit int, sigil byte) (total int) {
	floor.Each(func(value byte, x, y int) {
		if value != sigil {
			return
		}
		if count := Around(floor, sigil, x, y); count < limit {
			total++
		}
	})

	return
}

func RemoveMovables(floor *Floor[byte], limit int, sigil byte) (total int) {
	for {
		removed := RemoveMovablesStep(floor, limit, sigil)
		total += removed
		if removed == 0 {
			return
		}
	}
}

func RemoveMovablesStep(floor *Floor[byte], limit int, sigil byte) int {
	toRemove := [][2]int{}

	floor.Each(func(value byte, x, y int) {
		if value != sigil {
			return
		}
		if count := Around(floor, sigil, x, y); count < limit {
			toRemove = append(toRemove, [2]int{x, y})
		}
	})

	for _, tuple := range toRemove {
		floor.Set('.', tuple[0], tuple[1])
	}

	return len(toRemove)
}

var deltas = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func Around[T comparable](floor *Floor[T], search T, x, y int) (count int) {
	for _, d := range deltas {
		if value := floor.SafeGet(x+d[0], y+d[1]); value == search {
			count++
		}
	}
	return
}

func NewFloor[T any](w, h int, data []T) *Floor[T] {
	return &Floor[T]{
		w:    w,
		h:    h,
		data: data,
	}
}

type Floor[T any] struct {
	w, h int
	data []T
}

func (f *Floor[T]) index(x, y int) int {
	return y*f.w + x
}

func (f *Floor[T]) Get(x, y int) T {
	return f.data[f.index(x, y)]
}

func (f *Floor[T]) SafeGet(x, y int) (value T) {
	if f.Has(x, y) {
		value = f.Get(x, y)
	}
	return
}

func (f *Floor[T]) Set(value T, x, y int) {
	f.data[f.index(x, y)] = value
}

func (f *Floor[T]) Has(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.w && y < f.h
}

func (f *Floor[T]) Each(cb func(value T, x, y int)) {
	for index := range len(f.data) {
		x := index % f.w
		y := index / f.h
		cb(f.data[index], x, y)
	}
}

func ParseInput(r io.Reader) *Floor[byte] {
	var w, h int
	data := make([]byte, 0)

	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

		h++
		if w == 0 {
			w = len(line)
		}

		data = append(data, line...)
	}

	return NewFloor(w, h, data)
}
