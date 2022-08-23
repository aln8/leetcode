/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

// Time: O(n), worst case validate every node
// Space: O(height), const space for recurrsion on call stack
func isValidBST(root *TreeNode) bool {
	return isBST(root, minInt, maxInt)
}

func isBST(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	// min, cur, max
	return root.Val > min &&
		root.Val < max &&
		isBST(root.Left, min, root.Val) &&
		isBST(root.Right, root.Val, max)
}

// @lc code=end

