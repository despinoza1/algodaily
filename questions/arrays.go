package questions

import (
	"algodaily/questions/datastructure"
	"sort"
)

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

// PowerOfThree checks if a number is a power of three
func PowerOfThree(num int) bool {
	for num > 0 {
		if num == 1 {
			return true
		}
		num /= 3
	}
	return false
}

// IsPowerOf checks if a number is a power of base
func IsPowerOf(num, base int) bool {
	if base%2 == 0 && num%2 != 0 {
		return false
	}

	for num > 0 {
		if num == 1 {
			return true
		}
		num /= base
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MissingInUnsorted finds the missing number from an unsorted array
func MissingInUnsorted(arr []int, lower, upper int) int {
	upperSum := upper * (upper + 1) / 2
	lowerSum := lower * (lower - 1) / 2

	sum := 0
	for _, v := range arr {
		sum += v
	}

	return (upperSum - lowerSum) - sum
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

// SubArraySum checks if there is a contiguous subarray that sum to n
func SubArraySum(arr []int, n int) bool {
	sumMap := make(map[int]bool)
	sum := 0

	for _, v := range arr {
		sum += v

		if _, ok := sumMap[sum-n]; ok {
			return true
		}

		sumMap[sum] = true
	}

	return false
}

func genPrimes(n int) map[int]bool {
	primes := make(map[int]bool)

OUTER_LOOP:
	for i := 1; i < n; i++ {
		for k := range primes {
			if (i+1)%k == 0 {
				continue OUTER_LOOP
			}
		}
		primes[(i + 1)] = true
	}

	return primes
}

// SumPrimes sums all prime numbers less than n
func SumPrimes(n int) int {
	sum := 0

	primes := genPrimes(n)
	for k := range primes {
		sum += k
	}

	return sum
}

// MinInRotArray finds the min of a rotated array
func MinInRotArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		if arr[start] < arr[end] {
			return arr[start]
		}

		mid := (start + end) / 2

		if arr[mid] >= arr[start] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return arr[start]
}

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

// TreatDistribution returns maximum number of unique kinds of snacks
func TreatDistribution(arr []int) int {
	set := datastructure.NewSet(arr)

	if len(set.List) >= (len(arr) / 2) {
		return (len(arr) / 2)
	}

	return len(set.List)
}

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

// MaxProductOfThree returns the max product from an array using 3 elements
func MaxProductOfThree(arr []int) int {
	sort.Ints(arr)

	prod1 := arr[0] * arr[1] * arr[len(arr)-1]
	prod2 := arr[len(arr)-1] * arr[len(arr)-2] * arr[len(arr)-3]

	if prod1 > prod2 {
		return prod1
	}

	return prod2
}

// SortedTwoSum returns the position of two numbers that sum to goal from a sorted array
func SortedTwoSum(arr []int, goal int) []int {
	i, j := 0, len(arr)-1

	for i < j {
		sum := arr[i] + arr[j]

		if sum == goal {
			return []int{i, j}
		} else if sum > goal {
			j--
		} else {
			i++
		}
	}

	return []int{}
}
