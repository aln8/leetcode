/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
// @lc code=start
/*
	   a a b b b c a a
	a  1 0 0 0 0 0 0 0
	*  1 1 1 1 1 1 1 1
	b  * 0 1 1 1 0 0 0
	*  * 0 1 1 1 1 1 1
	*  * * 1 1 1 1 1 1
	?  * * * 1 1 1 1 1
	c  * * * * 0 1 0 0
	a  * * * * * 0 1 0



		a  b  c  d
	?   1  1  1  1
	*	1  1  1  1
	?	0  1  1  1
	*   0  1  1  1
	?	0  0  1  1
	*	0  0  1  1
	?	0  0  0  1
	*	0  0  0  1
*/
// package main

// import "fmt"

// func main() {
// 	fmt.Println(isMatch("", "*****"))
// 	fmt.Println(isMatch("ab", "*?****?*"))
// 	fmt.Println(isMatch("a", "a*"))
// 	fmt.Println(isMatch("a", "aa"))
// 	fmt.Println(isMatch("adceb", "*a***b"))
// 	fmt.Println(isMatch("abcabczzzde", "***abc???de***"))
// }

func isMatch(input string, pattern string) bool {
	pattern = trimRepeat(pattern, '*')
	// input empty
	if input == "" {
		if pattern == "*" || pattern == "" {
			return true
		}
		return false
	}

	// hold previous step
	dp := make([]bool, len(input))
	// hold idx increase beside '*'
	validIdx := 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '*' {
			// fill dp in ascending order
			for j := validIdx; j < len(input); j++ {
				dp[j] = j == 0 || dp[j-1] || dp[j]
			}
		} else {
			// fill with descending order
			for j := len(input) - 1; j >= validIdx; j-- {
				if pattern[i] == '?' || pattern[i] == input[j] {
					if j == 0 {
						dp[j] = true
					} else {
						dp[j] = dp[j-1]
					}
				} else {
					dp[j] = false
				}
			}
			validIdx++
			if validIdx > len(input) {
				return false
			}
		}
		fmt.Println(string(pattern[i]), validIdx, dp)
	}
	return dp[len(input)-1]
}

func trimRepeat(in string, r byte) string {
	// clean pattern prefix, suffix '*'
	nr := []byte{}
	for i := 0; i < len(in); i++ {
		// in[0] == r if trim by found, no worry about in[-1]
		if i == 0 || in[i] != r || in[i-1] != r {
			nr = append(nr, in[i])
		}
	}
	return string(nr)
}

// @lc code=end
