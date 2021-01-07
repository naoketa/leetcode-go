package medium153

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		mid := (start + end) / 2
		if nums[end] < nums[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
