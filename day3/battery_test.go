package day3_test

import (
	"aoc2025/day3"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay3Part1(t *testing.T) {
	file := lib.GetFile(t, "input.txt")
	actual := day3.Joltage(file, 2)
	t.Log(actual)
}

func TestSolveDay3Part2(t *testing.T) {
	file := lib.GetFile(t, "input.txt")
	actual := day3.Joltage(file, 12)
	t.Log(actual)
}

const testinput = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func TestExample(t *testing.T) {
	if actual := day3.Joltage(strings.NewReader(testinput), 2); actual != 357 {
		t.Error("unexpected value")
	}
}
