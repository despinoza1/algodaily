package questions

import (
	"algodaily/questions/datastructure"
	"strconv"
	"strings"
)

// FizzBuzz function
func FizzBuzz(n int) string {
	str := []rune("")

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			str = append(str, []rune("fizz")...)
		}
		if i%5 == 0 {
			str = append(str, []rune("buzz")...)
		}
		if i%3 != 0 && i%5 != 0 {
			str = append(str, []rune(strconv.Itoa(i))...)
		}
	}

	return string(str)
}

func isAlpha(char rune) bool {
	if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
		return true
	}

	return false
}

// ReverseAlphaChars reverses string but only the alphabetic characters
func ReverseAlphaChars(str string) string {
	s := []rune(str)

	j := len(s) - 1
	for i := 0; i < j; i++ {
		if isAlpha(s[i]) {
			for !isAlpha(s[j]) {
				j--
				if j < i {
					return string(s)
				}
			}
			temp := s[i]
			s[i] = s[j]
			s[j] = temp
			j--
		}
	}

	return string(s)
}

// IsAnagram checks if two strings are anagrams
func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	set1 := datastructure.NewSet()
	set2 := datastructure.NewSet()

	for _, char := range str1 {
		set1.List[int(char)] = struct{}{}
	}

	for _, char := range str2 {
		set2.List[int(char)] = struct{}{}
	}

	return set1.Intersects(set2)
}

// Palindrome checks if palindrome
func Palindrome(str string) bool {
	j := len(str) - 1

	for i := 0; i != j; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}

	return true
}

// SumDigits sums the digits of a number until only one remains
func SumDigits(num int) int {
	digits := []rune(strconv.Itoa(num))

	for len(digits) > 1 {
		res := 0
		for _, v := range digits {
			digit, _ := strconv.Atoi(string(v))
			res += digit
		}

		digits = []rune(strconv.Itoa(res))
	}

	num, _ = strconv.Atoi(string(digits[0]))
	return num
}

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

func getProccessedStr(str string) string {
	stk := make(datastructure.Stack, 0)

	for _, char := range str {
		if char == '$' {
			stk, _ = stk.Pop()
		} else {
			stk = stk.Push(char)
		}
	}

	res := ""
	var tmp interface{}
	for i := 0; i < len(stk); i++ {
		stk, tmp = stk.Pop()
		res += string(tmp.(rune))
	}

	return res
}

// DollarSignDeletion returns if after remove every character
// preceding a $ if all elements in array are equivalent
func DollarSignDeletion(strs []string) bool {
	tmp := getProccessedStr(strs[0])

	for _, str := range strs {
		if tmp != getProccessedStr(str) {
			return false
		}
	}

	return true
}

// FirstNonRepeat returns the first character that does not repeat
func FirstNonRepeat(str string) string {
	seen := make(map[rune]int)
	order := make([]rune, 0)

	for _, char := range str {
		seen[char]++
		if seen[char] == 1 {
			order = append(order, char)
		}
	}

	for _, char := range order {
		if seen[char] == 1 {
			return string(char)
		}
	}

	return ""
}
