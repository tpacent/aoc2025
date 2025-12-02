package day1

import (
	"aoc2025/lib"
	"bufio"
	"io"
	"iter"
)

func NewDial(init, size int) Dial {
	return Dial{
		state: init,
		size:  size,
	}
}

type Dial struct {
	state int
	size  int
}

func (d *Dial) Rotate(n int, allClicks bool) (clicks int) {
	newState := d.state + n

	if allClicks {
		clicks += lib.Abs(newState / d.size)
		if d.state > 0 && newState <= 0 {
			clicks++
		}
	} else if newState%d.size == 0 {
		clicks++
	}

	newState %= d.size
	if newState < 0 {
		newState += d.size
	}

	d.state = newState
	return
}

func CountZero(stream iter.Seq[int], dial Dial, allClicks bool) (count int) {
	for r := range stream {
		count += dial.Rotate(r, allClicks)
	}

	return
}

func ParseInput(r io.Reader) iter.Seq[int] {
	return func(yield func(int) bool) {
		for scanner := bufio.NewScanner(r); scanner.Scan(); {
			line := scanner.Text()
			if len(line) < 2 {
				continue
			}

			n := lib.MustAtoi(line[1:])
			if line[0] == 'L' {
				n *= -1
			}

			if !yield(n) {
				break
			}
		}
	}
}
