package easy1512

func numIdenticalPairs(nums []int) int {
	ans := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ans++
			}
		}
	}
	return ans
}
