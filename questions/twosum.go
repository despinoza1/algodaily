package questions

// TwoSum finds the indices of the array that sum to the goal
func TwoSum(arr []int, goal int) [2]int {
	hash := make(map[int]int)
	var res [2]int

	for i, v := range arr {
		diff := goal - v
		if _, ok := hash[v]; ok {
			res[0] = hash[v]
			res[1] = i
		} else {
			hash[diff] = i
		}
	}

	return res
}
