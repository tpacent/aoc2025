package day3

import (
	"aoc2025/lib"
	"bufio"
	"io"
	"iter"
)

func Joltage(r io.Reader, size int) (total int) {
	for digits := range ParseInput(r) {
		total += MaxNumber(digits, size)
	}
	return
}

func MaxNumber(digits []uint8, size int) (value int) {
	num := make([]uint8, 0, size)

	for size > 0 {
		index := MaxIndex(digits[:len(digits)-size+1])
		num = append(num, digits[index])
		digits = digits[index+1:]
		size--
	}

	for index, n := range num {
		value += lib.Pow(10, len(num)-index-1) * int(n)
	}

	return
}

func MaxIndex(digits []uint8) (index int) {
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
	scanner := bufio.NewScanner(r)

	return func(yield func([]uint8) bool) {
		for scanner.Scan() {
			line := scanner.Text()

			if len(line) == 0 {
				continue
			}

			digits := make([]uint8, 0, len(line))

			for _, c := range line {
				digits = append(digits, uint8(lib.MustAtoi(string(c))))
			}

			if !yield(digits) {
				return
			}

		}
	}

}
