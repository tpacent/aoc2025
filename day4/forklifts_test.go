package day4_test

import (
	"aoc2025/day4"
	"aoc2025/lib"
	"bytes"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	file := lib.GetFile(t, "input.txt")
	floor := day4.ParseInput(file)
	actual := day4.RemoveMovablesStep(floor, 4)
	t.Log(actual)
}

func TestDay4Part2(t *testing.T) {
	file := lib.GetFile(t, "input.txt")
	floor := day4.ParseInput(file)
	actual := day4.RemoveMovables(floor, 4)
	t.Log(actual)
}

const example = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestExample(t *testing.T) {
	floor := day4.ParseInput(bytes.NewReader([]byte(example)))
	actual := day4.RemoveMovablesStep(floor, 4)
	if actual != 13 {
		t.Log("unexpected value")
	}
}

func TestExample2(t *testing.T) {
	floor := day4.ParseInput(bytes.NewReader([]byte(example)))
	actual := day4.RemoveMovables(floor, 4)
	if actual != 43 {
		t.Log("unexpected value")
	}
}
