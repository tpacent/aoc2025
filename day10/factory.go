package day10

import (
	"aoc2025/lib"
	"bufio"
	"io"
	"iter"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func CountPresses(r io.Reader) (total int) {
	for task := range ParseInput(r) {
		total += ConfigureIndicators([]int{0}, task.Buttons, task.Target, 1, map[int]struct{}{})
	}
	return
}

func CountJoltPresses(r io.Reader) (total int) {
	for task := range ParseInput(r) {
		total += ConfigureJoltage(
			make([]int8, len(task.Jolts)),
			task.Jolts,
			task.JoltButtons,
		)

	}
	return
}

func ConfigureIndicators(states []int, buttons []LightButton, target, level int, visited map[int]struct{}) int {
	if len(states) == 0 {
		return 0
	}

	for _, state := range states {
		visited[state] = struct{}{}
	}

	nextStates := make([]int, 0, len(states)*len(buttons))
	for _, state := range states {
		for _, btn := range buttons {
			nextState := btn(state)
			if nextState == target {
				return level
			}

			if _, ok := visited[nextState]; !ok {
				nextStates = append(nextStates, nextState)
			}
		}
	}

	return ConfigureIndicators(nextStates, buttons, target, level+1, visited)
}

func ConfigureJoltage(state, target []int8, buttons []JoltButton) int {
	return joltage(state, target, buttons, 0)
}

func joltage(state, target []int8, buttons []JoltButton, presses int) int {
	if slices.Equal(state, target) {
		return presses
	}

	if len(buttons) == 0 {
		return 0 // dead end
	}

	button := buttons[0]
	buttons = buttons[1:]

	minPresses := 0

	upperBound := getUpperBound(state, target, button)
	for pressCount := upperBound; pressCount >= 0; pressCount-- {
		nextState := button.Press(slices.Clone(state), pressCount)

		count := joltage(nextState, target, buttons, presses+int(pressCount))

		if count == 0 {
			continue
		}

		if minPresses == 0 || count < minPresses {
			minPresses = count
		}
	}

	return minPresses
}

func getUpperBound(state, target []int8, button JoltButton) (presses int8) {
	for _, index := range button {
		diff := target[index] - state[index]
		if diff <= 0 {
			return 0 // cannot push this
		}

		if presses == 0 || diff < presses {
			presses = diff
		}
	}
	return
}

type LightButton func(state int) int

type JoltButton []int

// Press button specified number of times.
// Does not allocate: the caller should clone the slice if desired.
func (jb JoltButton) Press(state []int8, times int8) []int8 {
	for _, pos := range jb {
		state[pos] += times
	}
	return state
}

type Task struct {
	Buttons     []LightButton
	Target      int
	Jolts       []int8
	JoltButtons []JoltButton
}

func ParseInput(r io.Reader) iter.Seq[Task] {
	scanner := bufio.NewScanner(r)

	return func(yield func(Task) bool) {
		for scanner.Scan() {
			line := scanner.Text()

			if len(line) == 0 {
				continue
			}

			if !yield(ParseLine(line)) {
				return
			}
		}
	}
}

var linetask = regexp.MustCompile(`^\[([#.]+)\] ((?:\((?:\d+,?)+\) ?)+){((?:\d+,?)+)}$`)

func ParseLine(line string) (task Task) {
	matches := linetask.FindStringSubmatch(line)
	if matches == nil {
		panic("unexpected")
	}

	targetstr := matches[1]
	buttonsstr := matches[2]
	joltagestr := matches[3]

	task.Target = parseTarget(targetstr)
	task.Jolts = parseSeq[int8](joltagestr)

	for _, buttonspec := range strings.Fields(buttonsstr) {
		toggles := parseSeq[int](strings.Trim(buttonspec, "()"))
		task.Buttons = append(task.Buttons, NewButton(toggles, len(targetstr)))
		task.JoltButtons = append(task.JoltButtons, NewJoltButton(toggles))
	}

	return
}

func parseSeq[T lib.Intlike](s string) (out []T) {
	for nstr := range strings.SplitSeq(s, ",") {
		out = append(out, T(lib.MustAtoi(nstr)))
	}
	return
}

func parseTarget(s string) int {
	s = strings.Map(func(r rune) rune {
		switch r {
		case '.':
			return '0'
		case '#':
			return '1'
		default:
			return -1
		}
	}, s)
	target, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(target)
}

func NewButton(toggles []int, size int) LightButton {
	mask := 0
	for _, toggle := range toggles {
		mask |= 1 << (size - 1 - toggle)
	}

	return func(state int) int {
		return state ^ mask
	}
}

func NewJoltButton(toggles []int) JoltButton {
	return JoltButton(slices.Clone(toggles))
}
