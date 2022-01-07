package searchInsertPosition

// Binary search
// left <= right : left = mid + 1; right = mid - 1. Target one element in the last nums[mid] == target
// left < right: left = mid + 1; right = mid. Target one element with left.
// using right = length, making the most right avaible
// left < right - 1: left = mid, right = mid.
// Target will fall into [left, right], unless target < nums[0] pr target > nums[len(nums) - 1]
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	var left, right int
	for left, right = 0, len(nums); left < right; {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
