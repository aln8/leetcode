/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
/*

	first missing must in [1, N+1]
	because if 1 to N exist, i will be N+1

	1. first time move all not in [1, N] to 0
	2. orderly check which hole is first 0
*/
func firstMissingPositive3(nums []int) int {
	// first time
	l := len(nums)
	for i := 0; i < l; i++ {
		for nums[i] != 0 && nums[i] != i+1 {
			cur := nums[i]
			// end since find beyond boundary
			if cur > l || cur < 1 || nums[cur-1] == nums[i] {
				nums[i] = 0
				break
			}
			// swap (cur, nums[cur-1]), go find next
			nums[i], nums[cur-1] = nums[cur-1], nums[i]
		}
	}

	// second time
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return l + 1
}

// loop, put num > 0 and < size into correct spot in new array, check first value/spot not match element
// Time O(n), space O(n)
func firstMissingPositive2(nums []int) int {
	l := len(nums)
	h := make([]int, l)
	for _, n := range nums {
		if n > 0 && n <= l {
			h[n-1] = n
		}
	}
	for i, n := range h {
		if n-1 != i {
			return i + 1
		}
	}
	return l + 1
}

// instead put into new array, now do swap.
// 1. after swap, should re-check new value
// 2. check if current spot value = swap value, skip it, otherwise it will go infinite loop
// Time O(n), Space O(1)
func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < l; {
		if nums[i] > 0 && nums[i] <= l && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			continue
		}
		i++
	}

	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}
	return l + 1
}

// @lc code=end

