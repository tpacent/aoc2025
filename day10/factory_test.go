package day10_test

import (
	"aoc2025/day10"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay10Part1(t *testing.T) {
	actual := day10.CountPresses(lib.GetFile(t, "input.txt"))
	t.Log(actual)
}

func TestSolveDay10Part2(t *testing.T) {
	t.Log("TODO")
}

const input = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

func TestExample(t *testing.T) {
	actual := day10.CountPresses(strings.NewReader(input))
	if actual != 7 {
		t.Error("unexpected value")
	}
}

func TestExample2(t *testing.T) {
	actual := day10.CountJoltPresses(strings.NewReader(input))
	if actual != 33 {
		t.Error("unexpected value")
	}
}
