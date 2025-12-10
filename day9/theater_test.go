package day9_test

import (
	"aoc2025/day9"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay9Part1(t *testing.T) {
	points := day9.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day9.GreatestArea(points)
	t.Log(actual)
}

func TestSolveDay9Part2(t *testing.T) {
	points := day9.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day9.GreatestAreaInsidePerimeter(points)
	t.Log(actual)
}

const example = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestExample(t *testing.T) {
	points := day9.ParseInput(strings.NewReader(example))
	if actual := day9.GreatestArea(points); actual != 50 {
		t.Error("unexpected value")
	}
}

func TestExample2(t *testing.T) {
	points := day9.ParseInput(strings.NewReader(example))
	if actual := day9.GreatestAreaInsidePerimeter(points); actual != 24 {
		t.Error("unexpected value")
	}
}
