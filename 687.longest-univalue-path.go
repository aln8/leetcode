/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
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
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	now := triverse(root, &max)
	fmt.Println(max, now)
	return MaxInt(max, now) - 1
}

// return longest path, and update max each step
func triverse(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	l := triverse(root.Left, max)
	r := triverse(root.Right, max)

	// both 0, mean no child, return 1
	if l == 0 && r == 0 {
		return 1
	}

	// only right child, check value eq
	if l == 0 {
		if root.Val == root.Right.Val {
			// no need to update max, it can be done in next iter
			return r + 1
		}
		// update max, since path break
		*max = MaxInt(*max, r)
		return 1
	}

	// only left child, check value eq
	if r == 0 {
		if root.Val == root.Left.Val {
			return l + 1
		}
		*max = MaxInt(*max, l)
		return 1
	}

	// both child exist, check if all eq
	if root.Val == root.Left.Val && root.Val == root.Right.Val {
		*max = MaxInt(*max, l+r+1)
		if l > r {
			return l + 1
		}
		return r + 1
	}

	// left eq
	if root.Val == root.Left.Val {
		return l + 1
	}

	// right eq
	if root.Val == root.Right.Val {
		return r + 1
	}

	// no eq
	if l > r {
		*max = MaxInt(*max, l)
	}
	*max = MaxInt(*max, r)
	return 1
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// 			1
// 		2       3
//   4    2   	   5
