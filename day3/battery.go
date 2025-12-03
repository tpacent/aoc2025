package day3

import (
	"aoc2025/lib"
	"bufio"
	"io"
	"iter"
)

func Joltage(r io.Reader, size int) (total int) {
	for digits := range ParseInput(r) {
		total += maxNumber(digits, size)
	}
	return
}

func maxNumber(digits []uint8, size int) (value int) {
	if size == 0 {
		return 0
	}
	index := maxIndex(digits[:len(digits)-size+1])
	value += lib.Pow(10, size-1) * int(digits[index])
	value += maxNumber(digits[index+1:], size-1)
	return
}

func maxIndex(digits []uint8) (index int) {
	var value uint8

	for i, d := range digits {
		if d > value {
			value = d
			index = i
		}
	}

	return
}

func ParseInput(r io.Reader) iter.Seq[[]uint8] {
	return func(yield func([]uint8) bool) {
		for scanner := bufio.NewScanner(r); scanner.Scan(); {
			line := scanner.Text()

			if len(line) == 0 {
				continue
			}

			digits := make([]uint8, len(line))
			for index := range len(line) {
				digits[index] = line[index] - '0'
			}

			if !yield(digits) {
				return
			}
		}
	}
}
