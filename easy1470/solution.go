package easy1470

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i := 1; i <= n; i++ {
		ans[2*i-2] = nums[i-1]
	}
	for i := 1; i <= n; i++ {
		ans[2*i-1] = nums[i+n-1]
	}
	return ans
}
