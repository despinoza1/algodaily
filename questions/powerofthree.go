package questions

// PowerOfThree checks if a number is a power of three
func PowerOfThree(num int) bool {
	for num > 0 {
		if num == 1 {
			return true
		}
		num /= 3
	}
	return false
}
