package day7_test

import (
	"aoc2025/day7"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay7Part1(t *testing.T) {
	mf, col, row := day7.ParseInput(lib.GetFile(t, "input.txt"))
	actual, _ := day7.CountAll(mf, col, row)
	t.Log(actual)
}

func TestSolveDay7Part2(t *testing.T) {
	mf, col, row := day7.ParseInput(lib.GetFile(t, "input.txt"))
	_, actual := day7.CountAll(mf, col, row)
	t.Log(actual)
}

const example = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestExample(t *testing.T) {
	splits, timelines := day7.CountAll(day7.ParseInput(strings.NewReader(example)))
	if splits != 21 {
		t.Errorf("unexpected value: %d", splits)
	}
	if timelines != 40 {
		t.Errorf("unexpected value: %d", timelines)
	}
}
