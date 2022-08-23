/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(node *TreeNode, a, b int) bool {
	re := false
	recur(node, a, b, 0, &re)
	return re
}

func recur(node *TreeNode, a, b, level int, re *bool) int {
	if node == nil {
		return -1
	}
	if node.Val == a || node.Val == b {
		return level
	}

	left := recur(node.Left, a, b, level+1, re)
	right := recur(node.Right, a, b, level+1, re)
	if left == right && left != -1 &&
		(left-1 != level ||
			right-1 != level) {
		*re = true
	}
	if left != -1 {
		return left
	}
	return right
}

// @lc code=end

