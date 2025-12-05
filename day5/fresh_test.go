package day5_test

import (
	"aoc2025/day5"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay1Part1(t *testing.T) {
	file := lib.GetFile(t, "input.txt")
	t.Log(day5.CountFresh(day5.ParseInput(file)))
}

func TestSolveDay1Part2(t *testing.T) {
	file := lib.GetFile(t, "input.txt")
	_, ranges := day5.ParseInput(file)
	t.Log(day5.TotalFresh(ranges))
}

const example = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestExample(t *testing.T) {
	ingredients, ranges := day5.ParseInput(strings.NewReader(example))
	if actual := day5.CountFresh(ingredients, ranges); actual != 3 {
		t.Errorf("unexpected value: %d", actual)
	}
}

func TestExample2(t *testing.T) {
	_, ranges := day5.ParseInput(strings.NewReader(example))
	if actual := day5.TotalFresh(ranges); actual != 14 {
		t.Errorf("unexpected value: %d", actual)
	}
}
