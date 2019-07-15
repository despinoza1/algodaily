package questions

func PowerOfThree(num int) bool {
	for num > 0 {
		if num == 1 {
			return true
		}
		num /= 3
	}
	return false
}
