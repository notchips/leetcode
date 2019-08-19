/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			for lo, hi, toFind := i+1, n-1, -nums[i]; lo < hi; {
				if nums[lo]+nums[hi] == toFind {
					ans = append(ans, []int{nums[i], nums[lo], nums[hi]})
					for lo+1 < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi-1 && nums[hi-1] == nums[hi] {
						hi--
					}
					lo, hi = lo+1, hi-1
				} else if nums[lo]+nums[hi] < toFind {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return ans
}
