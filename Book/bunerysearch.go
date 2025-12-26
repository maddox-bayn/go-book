package Book

func binarySearch(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	mid := len(nums) / 2

	if nums[mid] == target {
		return true
	} else if nums[mid] > target {
		return binarySearch(nums[:mid], target)
	}
	return binarySearch(nums[mid+1:], target)
}
