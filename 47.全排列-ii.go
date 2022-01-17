/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
	dfs
		123
	1(23)  1(32)

*/
package main

import "sort"

func main() {
	input := []int{1, 1, 2}
	permuteUnique(input)
}

// @lc code=start
func permuteUnique(nums []int) [][]int {
	re := [][]int{}
	sort.Ints(nums)

	var dfs func(layer int)
	dfs = func(layer int) {
		if layer == len(nums)-1 {
			d := make([]int, len(nums))
			copy(d, nums)
			re = append(re, d)
			return
		}

		vized := make(map[int]struct{}, len(nums))
		for i := layer; i < len(nums); i++ {
			if _, ok := vized[nums[i]]; ok {
				continue
			}
			vized[nums[i]] = struct{}{}
			nums[layer], nums[i] = nums[i], nums[layer]
			dfs(layer + 1)
			// delete(vized, nums[i])
			nums[layer], nums[i] = nums[i], nums[layer]
		}
	}
	dfs(0)
	return re
}

// @lc code=end
