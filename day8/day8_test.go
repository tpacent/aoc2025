package day8_test

import (
	"aoc2025/day8"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay8Part1(t *testing.T) {
	data := day8.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day8.CalcLargestCircuits(data, 1000, 3)
	t.Log(actual)
}

func TestSolveDay8Part2(t *testing.T) {
	data := day8.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day8.CalcLastJunctions(data)
	t.Log(actual)
}

const example = `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

func TestExample(t *testing.T) {
	data := day8.ParseInput(strings.NewReader(example))
	actual := day8.CalcLargestCircuits(data, 10, 3)
	if actual != 40 {
		t.Error("unexpected value")
	}
}
