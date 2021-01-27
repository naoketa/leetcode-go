package medium338

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := range ans {
		idx := i
		sum := 0
		for i != 0 {
			if i%2 == 1 {
				sum++
			}
			i = (i - i%2) / 2
		}
		ans[idx] = sum
	}
	return ans
}
