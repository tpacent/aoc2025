package day6

import (
	"aoc2025/lib"
	"bufio"
	"io"
	"strings"
)

type Op byte

const (
	OpMul Op = '*'
	OpAdd Op = '+'
)

func GrandTotal(nums [][]int, ops []Op) (total int) {
	var result int
	for index, op := range ops {
		switch op {
		case OpMul:
			result = 1
		default:
			result = 0
		}

		for _, n := range nums[index] {
			switch op {
			case OpMul:
				result *= n
			case OpAdd:
				result += n
			}
		}
		total += result
	}
	return
}

func ParseInput(r io.Reader) (nums [][]int, ops []Op) {
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if nums == nil {
			nums = make([][]int, len(strings.Fields(line)))
		}

		if strings.ContainsAny(line, "+*") {
			for _, op := range strings.Fields(line) {
				ops = append(ops, Op(op[0]))
			}
			break
		}

		for index, n := range strings.Fields(line) {
			nums[index] = append(nums[index], lib.MustAtoi(n))
		}
	}

	return
}

func ParseInputCols(r io.Reader) (nums [][]int, ops []Op) {
	var lines []string

	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	chunk := []int{}
	var op Op
	lastLine := len(lines) - 1

	for col := range len(lines[0]) {
		end := true
		last := col == len(lines[0])-1
		digits := make([]byte, 0, len(lines))

		if maybeOp := lines[lastLine][col]; maybeOp != ' ' {
			op = Op(maybeOp)
		}

		for row := range len(lines) - 1 {
			if c := lines[row][col]; c != ' ' {
				digits = append(digits, c)
				end = last
			}
		}

		if !end || last {
			chunk = append(chunk, lib.MustAtoi(string(digits)))
		}

		if end {
			nums = append(nums, chunk)
			ops = append(ops, op)
			chunk = []int{}
		}
	}

	return
}
