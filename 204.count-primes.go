/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
/*

	Eular prime filter
	for n, fill n*n, n*(n-1), n*(n-2)....
	it make sure, every item only filled once,
	so its better than ethos method
*/
func countPrimes(n int) int {
	m := make([]bool, n+1)
	count := 0
	for i := 2; i < n; i++ {
		// if prime
		if !m[i] {
			count++
		}
		// fill
		for j := 2; j <= i && j*i <= n; j++ {
			m[j*i] = true
		}
	}
	return count
}

// @lc code=end

