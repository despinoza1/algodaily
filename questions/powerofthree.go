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

// IsPowerOf checks if a number is a power of base
func IsPowerOf(num, base int) bool {
	if base%2 == 0 && num%2 != 0 {
		return false
	}

	for num > 0 {
		if num == 1 {
			return true
		}
		num /= base
	}
	return false
}
