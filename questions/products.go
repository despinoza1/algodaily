package questions

// ProductExceptSelf returns an array where each of its elements
// is the product of all elements in the array except itself
func ProductExceptSelf(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	prod := make([]int, len(arr))
	for i := range prod {
		prod[i] = 1
	}

	temp := 1
	for i, v := range arr {
		prod[i] *= temp
		temp *= v
	}

	temp = 1
	size := len(arr) - 1
	for i := range arr {
		prod[size-i] *= temp
		temp *= arr[size-i]
	}

	return prod
}
