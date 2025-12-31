package Book

func binerySearch(nums []int, terget int) bool {
	if len(nums) == 0 {
		return false
	}

	mid := len(nums) / 2
	if nums[mid] == terget {
		return true
	}
	if nums[mid] > terget {
		return binerySearch(nums[:mid], terget)
	}
	return binerySearch(nums[mid+1:], terget)

}
