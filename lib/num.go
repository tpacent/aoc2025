package lib

import "strconv"

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return n
}

func Abs(n int) int {
	if n < 0 {
		n *= -1
	}

	return n
}
