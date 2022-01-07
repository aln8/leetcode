/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
	solution 1:

		handle all negative for loop

		handle positive for loop
*/

// @lc code=start
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}

	max, i := nums[0], 0

	// handle all negative case
	for ; i < len(nums); i++ {
		if nums[i] > 0 {
			max = nums[i]
			break
		}
		max = GetMax(max, nums[i])
	}
	if i == len(nums) {
		return max
	}

	curMax := 0
	for ; i < len(nums); i++ {
		if nums[i] < 0 {
			max = GetMax(curMax, max)
		}

		curMax += nums[i]
		if curMax < 0 {
			curMax = 0
		}
	}
	return GetMax(max, curMax)
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

