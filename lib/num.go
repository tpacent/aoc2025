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

func NumDigits(n int) (count int) {
	for n > 0 {
		count++
		n /= 10
	}
	return
}

func Pow(a, b int) int {
	result := 1
	for b != 0 {
		if (b & 1) != 0 {
			result *= a
		}
		b >>= 1
		a *= a
	}
	return result
}
