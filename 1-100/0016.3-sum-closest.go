/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for lo, hi := i+1, n-1; lo < hi; {
			sum := nums[i] + nums[lo] + nums[hi]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum == target {
				break
			} else if sum < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
