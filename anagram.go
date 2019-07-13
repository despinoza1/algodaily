package main

import "strings"

// IsAnagram checks if two strings are anagrams
func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	set1 := NewSet()
	set2 := NewSet()

	for _, char := range str1 {
		set1.list[int(char)] = struct{}{}
	}

	for _, char := range str2 {
		set2.list[int(char)] = struct{}{}
	}

	return set1.Intersects(set2)
}
