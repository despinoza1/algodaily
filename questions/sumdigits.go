package questions

import "strconv"

// SumDigits sums the digits of a number until only one remains
func SumDigits(num int) int {
	digits := []rune(strconv.Itoa(num))

	for len(digits) > 1 {
		res := 0
		for _, v := range digits {
			digit, _ := strconv.Atoi(string(v))
			res += digit
		}

		digits = []rune(strconv.Itoa(res))
	}

	num, _ = strconv.Atoi(string(digits[0]))
	return num
}
