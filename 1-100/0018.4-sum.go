/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0, 10)
	for i := 0; i < n-3; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			for j := i + 1; j < n-2; j++ {
				if j == i+1 || (j > i+1 && nums[j] != nums[j-1]) {
					for lo, hi := j+1, n-1; lo < hi; {
						sum := nums[i] + nums[j] + nums[lo] + nums[hi]
						if sum == target {
							ans = append(ans, []int{nums[i], nums[j], nums[lo], nums[hi]})
							for lo+1 < hi && nums[lo] == nums[lo+1] {
								lo++
							}
							for lo < hi-1 && nums[hi-1] == nums[hi] {
								hi--
							}
							lo, hi = lo+1, hi-1
						} else if sum < target {
							lo++
						} else {
							hi--
						}
					}
				}
			}
		}
	}
	return ans
}
