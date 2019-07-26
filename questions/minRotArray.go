package questions

// MinInRotArray finds the min of a rotated array
func MinInRotArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		if arr[start] < arr[end] {
			return arr[start]
		}

		mid := (start + end) / 2

		if arr[mid] >= arr[start] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return arr[start]
}
