package day1_test

import (
	"aoc2025/day1"
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	expected1 = 969
	expected2 = 5887
)

func TestDay1(t *testing.T) {
	for _, tcase := range []struct {
		Clicks   bool
		Expected int
	}{
		{Clicks: false, Expected: expected1},
		{Clicks: true, Expected: expected2},
	} {
		t.Run(fmt.Sprintf("allclicks: %v", tcase.Clicks), func(t *testing.T) {
			src, err := os.Open("input.txt")
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { _ = src.Close() })

			actual := day1.CountZero(
				day1.ParseInput(src),
				day1.NewDial(50, 100),
				tcase.Clicks,
			)

			if actual != tcase.Expected {
				t.Errorf("Unexpected value: %d", actual)
			}
		})
	}
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
