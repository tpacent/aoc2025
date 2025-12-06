package day6

import (
	"aoc2025/lib"
	"bufio"
	"bytes"
	"io"
	"iter"
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
	var lines [][]byte
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		if line := scanner.Bytes(); len(line) > 0 {
			lines = append(lines, bytes.Clone(line))
		}
	}

	var opNums []int
	for chunk := range reverseLineSeq(lines) {
		num, op, ok := ParseChunk(chunk)
		if !ok {
			continue
		}
		opNums = append(opNums, num)
		if op == OpAdd || op == OpMul {
			nums = append(nums, opNums)
			ops = append(ops, op)
			opNums = nil
		}
	}

	return
}

func ParseChunk(chunk []byte) (num int, op Op, ok bool) {
	digits := bytes.Trim(chunk[:len(chunk)-1], " ")
	if len(digits) == 0 {
		return
	}
	num = lib.MustAtoi(string(digits))
	op = Op(chunk[len(chunk)-1])
	return num, op, true

}

func reverseLineSeq(lines [][]byte) iter.Seq[[]byte] {
	pos := len(lines[0]) - 1
	return func(yield func([]byte) bool) {
		for pos >= 0 {
			chunk := make([]byte, 0, len(lines))
			for _, line := range lines {
				chunk = append(chunk, line[pos])
			}
			if !yield(chunk) {
				return
			}
			pos--
		}
	}
}
