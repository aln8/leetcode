/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
// for int[m][n], it's 1d index = len(int[0]) * n + m
// for index, m = index % len(int[0])
// 			  n = index / len(int[0])
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for left, right := 0, len(matrix[0])*len(matrix)-1; left <= right; {
		mid := left + (right-left)/2
		cur := matrix[mid/len(matrix[0])][mid%len(matrix[0])]
		if cur == target {
			return true
		} else if cur > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end

