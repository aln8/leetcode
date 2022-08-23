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

package main

func main() {
	input := []int{1, 2, 3}
	permute(input)
}

// @lc code=start
func permute(nums []int) [][]int {
	re := [][]int{}

	var dfs func(layer int)
	dfs = func(layer int) {
		if layer == len(nums)-1 {
			d := make([]int, len(nums))
			copy(d, nums)
			re = append(re, d)
			return
		}

		for i := layer; i < len(nums); i++ {
			// swap
			nums[layer], nums[i] = nums[i], nums[layer]
			dfs(layer + 1)
			// swap back
			nums[layer], nums[i] = nums[i], nums[layer]
		}
	}
	dfs(0)
	return re
}

// @lc code=end
