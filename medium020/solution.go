package medium020

func myPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1
	case n > 0:
		return x * myPow(x, n-1)
	case n < 0:
		return 1 / x * myPow(x, n+1)
	default:
		return -1
	}
}
