/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// init
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil) // append nil for layer+1
	layerArr := []int{}
	result := [][]int{}

	var cur *TreeNode
	for len(queue) != 0 {
		// popup
		cur, queue = queue[0], queue[1:]
		// handle layer change
		if cur == nil {
			result = append(result, layerArr)
			layerArr = []int{}
			if len(queue) != 0 {
				queue = append(queue, nil) // append next layer
			}
			continue
		}

		// handle normal val
		layerArr = append(layerArr, cur.Val)

		// append generation
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

	return result
}

// @lc code=end

