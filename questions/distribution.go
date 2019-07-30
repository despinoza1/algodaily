package questions

// TreatDistribution returns maximum number of unique kinds of snacks
func TreatDistribution(arr []int) int {
	set := NewSet(arr)

	if len(set.list) >= (len(arr) / 2) {
		return (len(arr) / 2)
	}

	return len(set.list)
}
