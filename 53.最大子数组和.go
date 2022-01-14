/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
	solution 1:

		handle all negative for loop

		handle positive for loop

	solution 2:
		divide it

		return left, right, max
*/

// @lc code=start
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}
	_, _, max, _ := maxPartArray(0, len(nums)-1, nums)
	return max
}

// return lmax, rmax, max, sum
// lmax, contains most left num's max serials
// rmax, contains most right num's max serials
// max most max serials
// current partition total sum
func maxPartArray(l, r int, nums []int) (int, int, int, int) {
	// edge case
	if l >= r {
		return nums[l], nums[l], nums[l], nums[l]
	}

	// partition lp: leftpart, rp: rightpart
	p := (r-l)/2 + l
	lplmax, lprmax, lpmax, lpsum := maxPartArray(l, p, nums)
	rplmax, rprmax, rpmax, rpsum := maxPartArray(p+1, r, nums)

	curlmax := GetMax(lpsum+rplmax, lplmax)
	currmax := GetMax(lprmax+rpsum, rprmax)
	cursum := lpsum + rpsum
	curmax := GetMax(GetMax(lpmax, rpmax), lprmax+rplmax)

	return curlmax, currmax, curmax, cursum
}

// dynamic programing
// arr[i] means maxSubArray of arr[0:i]
// arr[i+1] = max(num[i+1] + arr[i], a[i])
func maxSubArray1(nums []int) int {
	max := nums[0]
	// assume nums can be modified
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// normal max
func maxSubArray2(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}

	max, i := nums[0], 0

	// handle all negative case
	for ; i < len(nums); i++ {
		if nums[i] > 0 {
			max = nums[i]
			break
		}
		max = GetMax(max, nums[i])
	}
	if i == len(nums) {
		return max
	}

	curMax := 0
	for ; i < len(nums); i++ {
		if nums[i] < 0 {
			max = GetMax(curMax, max)
		}

		curMax += nums[i]
		if curMax < 0 {
			curMax = 0
		}
	}
	return GetMax(max, curMax)
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
