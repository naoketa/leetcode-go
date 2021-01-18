package easy1486

func xorOperation(n int, start int) int {
	var ans int
	for i := range make([]int, n) {
		ans ^= start + 2*i
	}
	return ans
}
