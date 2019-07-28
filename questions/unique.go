package questions

// UniqueInArray finds the unique elements of an array
// Order is preserved
func UniqueInArray(arr []int) []int {
	order := make([]int, 0)
	seen := make(map[int]bool)

	for _, elem := range arr {
		if _, ok := seen[elem]; !ok {
			seen[elem] = true
			order = append(order, elem)
		}
	}

	return order
}
