package util

import "strconv"

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func IntDivide(numerator int, denominator int) int {
	remainder := numerator % denominator
	return (numerator - remainder) / denominator
}
