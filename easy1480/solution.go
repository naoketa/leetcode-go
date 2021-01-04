package easy1480

func runningSum(nums []int) []int {
	for i := range nums {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}

	return nums
}
