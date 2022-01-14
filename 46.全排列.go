/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列

			dfs

		     	root
		a         b         c
	 b    c    a     c    a     b
     c    a    c     a    b     a
*/
// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	re := [][]int{}
	dfs(0, nums, &re)
	return re
}

func dfs(layer int, nums []int, re *[][]int) {
	// end
	if layer == len(nums)-1 {
		d := make([]int, len(nums))
		copy(d, nums)
		*re = append(*re, d)
		return
	}

	dfs(layer+1, nums, re)
	for i := layer + 1; i < len(nums); i++ {
		// swap
		nums[layer], nums[i] = nums[i], nums[layer]
		dfs(layer+1, nums, re)
		// swap back
		nums[layer], nums[i] = nums[i], nums[layer]
	}
}

// @lc code=end
