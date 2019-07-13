package main

// Palindrome checks if palindrome
func Palindrome(str string) bool {
	j := len(str) - 1

	for i := 0; i != j; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}

	return true
}
