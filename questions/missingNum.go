package questions

// MissingInUnsorted finds the missing number from an unsorted array
func MissingInUnsorted(arr []int, lower, upper int) int {
	upperSum := upper * (upper + 1) / 2
	lowerSum := lower * (lower - 1) / 2

	sum := 0
	for _, v := range arr {
		sum += v
	}

	return (upperSum - lowerSum) - sum
}
