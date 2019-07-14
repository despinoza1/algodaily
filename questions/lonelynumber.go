package questions

// LonelyNumber finds the lonely number in an array
func LonelyNumber(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	hash := make(map[int]int)

	for _, v := range arr {
		hash[v]++
	}

	for num, count := range hash {
		if count == 1 {
			return num
		}
	}

	return 0
}
