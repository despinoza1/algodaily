package questions

// FirstNonRepeat returns the first character that does not repeat
func FirstNonRepeat(str string) string {
	seen := make(map[rune]int)
	order := make([]rune, 0)

	for _, char := range str {
		seen[char]++
		if seen[char] == 1 {
			order = append(order, char)
		}
	}

	for _, char := range order {
		if seen[char] == 1 {
			return string(char)
		}
	}

	return ""
}
