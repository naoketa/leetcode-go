package easy035

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 && nums[0] < target {
		return 1
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return start
	}
	mid := (start + end) / 2
	switch {
	case nums[mid] == target:
		return mid
	case nums[mid] > target:
		return binarySearch(nums, start, mid-1, target)
	case nums[mid] < target:
		return binarySearch(nums, mid+1, end, target)
	default:
		return -1
	}
}
