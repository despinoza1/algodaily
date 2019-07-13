package main

func isAlpha(char rune) bool {
	if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
		return true
	}

	return false
}

// ReverseAlphaChars reverses string but only the alphabetic characters
func ReverseAlphaChars(str string) string {
	s := []rune(str)

	j := len(s) - 1
	for i := 0; i < j; i++ {
		if isAlpha(s[i]) {
			for !isAlpha(s[j]) {
				j--
				if j < i {
					return string(s)
				}
			}
			temp := s[i]
			s[i] = s[j]
			s[j] = temp
			j--
		}
	}

	return string(s)
}
