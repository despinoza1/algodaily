package questions

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
