package advantage_shuffle

func search(nums []int, target int) int {
	if len(nums) == 0 || nums[len(nums)-1] < target || nums[0] > target {
		return -1
	}

	for left, right := 0, len(nums); left < right; {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
