package medium020

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		d := myPow(x, n/2)
		return d * d
	}
	if n > 0 {
		return x * myPow(x, n-1)
	}
	return 1 / x * myPow(x, n+1)

}
