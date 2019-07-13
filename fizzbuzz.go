package main

import "strconv"

// FizzBuzz function
func FizzBuzz(n int) string {
	str := []rune("")

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			str = append(str, []rune("fizz")...)
		}
		if i%5 == 0 {
			str = append(str, []rune("buzz")...)
		}
		if i%3 != 0 && i%5 != 0 {
			str = append(str, []rune(strconv.Itoa(i))...)
		}
	}

	return string(str)
}
