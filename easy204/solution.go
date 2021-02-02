package easy204

func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	var cnt int
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			cnt++
			for j := 2; i*j < n; j++ {
				isNotPrime[i*j] = true
			}
		}
	}
	return cnt
}
