package day6_test

import (
	"aoc2025/day6"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay6Part1(t *testing.T) {
	nums, ops := day6.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day6.GrandTotal(nums, ops)
	t.Log(actual)
}

func TestSolveDay6Part2(t *testing.T) {
	nums, ops := day6.ParseInputCols(lib.GetFile(t, "input.txt"))
	actual := day6.GrandTotal(nums, ops)
	t.Log(actual)
}

const example = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func TestExample(t *testing.T) {
	nums, ops := day6.ParseInput(strings.NewReader(example))
	if actual := day6.GrandTotal(nums, ops); actual != 4277556 {
		t.Errorf("unexpected value: %d", actual)
	}
}

func TestExample2(t *testing.T) {
	nums, ops := day6.ParseInputCols(strings.NewReader(example))
	if actual := day6.GrandTotal(nums, ops); actual != 3263827 {
		t.Errorf("unexpected value: %d", actual)
	}
}
