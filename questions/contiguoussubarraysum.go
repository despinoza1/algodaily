package questions

// SubArraySum checks if there is a contiguous subarray that sum to n
func SubArraySum(arr []int, n int) bool {
	sumMap := make(map[int]bool)
	sum := 0

	for _, v := range arr {
		sum += v

		if _, ok := sumMap[sum-n]; ok {
			return true
		}

		sumMap[sum] = true
	}

	return false
}
