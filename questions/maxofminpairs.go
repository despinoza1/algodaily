package questions

import "sort"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MaxOfMinPairs Gets the max of the min pairs
func MaxOfMinPairs(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	sum := 0
	for i := 0; i < len(arr); i += 2 {
		sum += arr[i]
	}

	return sum
}
