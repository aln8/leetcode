/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// this map hold index of meet element
	// if index > lt and exists in map means meet
	m := make(map[byte]int)
	r := []byte(s)
	// max & sliding window
	// rt: right most + 1
	// lt: left most
	max, lt, rt := 0, 0, 0
	for i, c := range r {
		ci, ok := m[c]
		if ok && ci >= lt {
			if rt-lt > max {
				max = rt - lt
			}
			lt = ci + 1
		}
		m[c] = i
		rt++
	}

	if rt-lt > max {
		max = rt - lt
	}

	return max
}

// @lc code=end


