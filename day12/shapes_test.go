package day12_test

import (
	"aoc2025/day12"
	"aoc2025/lib"
	"testing"
)

func TestSolveDay12(t *testing.T) {
	actual := day12.Solve(lib.GetFile(t, "input.txt"))
	t.Log(actual)
}
