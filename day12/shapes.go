package day12

import (
	"aoc2025/lib"
	"bufio"
	"bytes"
	"io"
	"strings"
)

func Solve(r io.ReadSeeker) (passed int) {
	shapes, tasks := ParseInput(r)
	for _, task := range tasks {
		area := 0
		for id, count := range task.Counts {
			area += shapes[id] * count
		}
		if area > task.Width*task.Height {
			continue
		}
		passed++
	}
	return
}

type Task struct {
	Width  int
	Height int
	Counts []int
}

func ParseInput(r io.ReadSeeker) (shapes []int, tasks []Task) {
	shape := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.Contains(line, []byte{'x'}) {
			break
		}
		if len(line) == 0 {
			if shape > 0 {
				shapes = append(shapes, shape)
				shape = 0
			}
			continue
		}
		if bytes.HasSuffix(line, []byte{':'}) {
			continue
		}
		shape += bytes.Count(line, []byte{'#'})
	}

	r.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "x") {
			continue
		}
		dim, counts, _ := strings.Cut(line, ":")
		wstr, hstr, _ := strings.Cut(dim, "x")
		task := Task{
			Width:  lib.MustAtoi(wstr),
			Height: lib.MustAtoi(hstr),
		}
		for count := range strings.FieldsSeq(counts) {
			task.Counts = append(task.Counts, lib.MustAtoi(count))
		}
		tasks = append(tasks, task)
	}

	return
}
