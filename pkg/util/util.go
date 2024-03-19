package util

import "strconv"

func MustAtoi(input string) int {
	ans, _ := strconv.Atoi(input)
	return ans
}
