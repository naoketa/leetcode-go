package medium560

func subarraySum(nums []int, k int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			ans++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
