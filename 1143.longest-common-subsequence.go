/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
/*

dp[i][j]: source i to target j, LCS(size)

		a  b  d  e  c
	a	1  1  1  1  1
	d	1  1  2  2  2
	h	1  1  2  2  2
	e	1  1  2	 3  3
	e   1  1  2  3  3

since update only by last line, could use 1d-array store dp
*/
func longestCommonSubsequence(source string, target string) int {
	// corner
	if len(source) == 0 || len(target) == 0 {
		return 0
	}

	// initial
	dp := make([]int, len(source))
	cur := make([]int, len(source))
	for i := range target {
		// fill cur if same
		for j := range source {
			if source[j] == target[i] {
				if j == 0 {
					cur[j] = 1
				} else {
					cur[j] = max(dp[j-1]+1, cur[j-1])
				}
			} else {
				if j > 0 && cur[j] < cur[j-1] {
					cur[j] = cur[j-1]
				}
			}
		}

		for k := range dp {
			dp[k] = cur[k]
		}
	}

	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
