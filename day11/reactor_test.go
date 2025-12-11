package day11_test

import (
	"aoc2025/day11"
	"aoc2025/lib"
	"strings"
	"testing"
)

func TestSolveDay11Part1(t *testing.T) {
	data := day11.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day11.Walk("you", "out", data)
	t.Log(actual)
}

func TestSolveDay11Part2(t *testing.T) {
	data := day11.ParseInput(lib.GetFile(t, "input.txt"))
	actual := day11.WalkThrough("svr", "out", data, "dac", "fft")
	t.Log(actual)
}

const example = `
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`

func TestExample(t *testing.T) {
	data := day11.ParseInput(strings.NewReader(example))
	if actual := day11.Walk("you", "out", data); actual != 5 {
		t.Error("unexpected value")
	}
}

const example2 = `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`

func TestExample2(t *testing.T) {
	data := day11.ParseInput(strings.NewReader(example2))
	if actual := day11.WalkThrough("svr", "out", data, "fft", "dac"); actual != 2 {
		t.Error("unexpected value")
	}
}
