package day1_test

import (
	"aoc2025/day1"
	"os"
	"strings"
	"testing"
)

func TestSolveDay1Part1(t *testing.T) {
	actual := day1.CountZero(
		day1.ParseInput(getFile(t, "input.txt")),
		day1.NewDial(50, 100),
		false,
	)
	t.Log(actual)
}

func TestSolveDay1Part2(t *testing.T) {
	actual := day1.CountZero(
		day1.ParseInput(getFile(t, "input.txt")),
		day1.NewDial(50, 100),
		true,
	)
	t.Log(actual)
}

func getFile(t *testing.T, path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = file.Close() })
	return file
}

const example = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestDial(t *testing.T) {
	dial := day1.NewDial(50, 100)
	seq := day1.ParseInput(strings.NewReader(example))
	if actual := day1.CountZero(seq, dial, false); actual != 3 {
		t.Errorf("unexpected value: %d", actual)
	}
}

func TestDialClicks(t *testing.T) {
	dial := day1.NewDial(50, 100)
	seq := day1.ParseInput(strings.NewReader(example))
	if actual := day1.CountZero(seq, dial, true); actual != 6 {
		t.Errorf("unexpected value: %d", actual)
	}
}
