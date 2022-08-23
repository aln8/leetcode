/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
/*
map 解法时间复杂度O(n), 空间复杂度O(n)
左右指针解法时间复杂度，排序O(nlogn), 找O(n)
*/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, n := range nums {
		if mi, ok := m[n]; ok {
			return []int{mi, i}
		}
		m[target-n] = i
	}
	return nil
}

// @lc code=end

