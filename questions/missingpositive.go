package questions

import "sort"

// LeastPositiveNumber finds the least positive integer missing from an array
func LeastPositiveNumber(arr []int) int {
	sort.Ints(arr)

	index := 0
	for i, v := range arr {
		if v > 0 {
			index = i
			break
		}
	}

	posArr := arr[index:]
	for i, v := range posArr {
		if v != (i + 1) {
			return i + 1
		}

	}

	return len(posArr) + 1
}
