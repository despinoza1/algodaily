package questions

// DetectSubStr returns the index of the first substring
// otherwise it returns -1 if not found
func DetectSubStr(str, substr string) int {
	idx, j := 0, 0
	flag := false

	for i := range str {
		if str[i] == substr[j] && !flag {
			idx = i
			j++
			flag = true
		} else if str[i] == substr[j] && flag {
			j++
		} else if str[i] != substr[j] && flag {
			j = 0
			flag = false
		}

		if j == len(substr)-1 {
			break
		}
	}

	if flag {
		return idx
	}

	return -1
}
