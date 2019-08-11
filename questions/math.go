package questions

var memo = make([]int, 50)

// FibonacciSequence returns the nth fibonacci number
func FibonacciSequence(n int) int {
	if len(memo) < n {
		tmp := make([]int, 2*(n-len(memo)))
		for i := range tmp {
			tmp[i] = -1
		}
		memo = append(memo, tmp...)
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = FibonacciSequence(n-1) + FibonacciSequence(n-2)
	return memo[n]
}

func init() {
	for i := range memo {
		memo[i] = -1
	}

	memo[0] = 0
	memo[1] = 1
}
