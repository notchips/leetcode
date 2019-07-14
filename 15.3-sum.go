/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
 func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0, 20)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || i > 0 && nums[i] != nums[i-1] {
			for lo, hi, sum := i+1, len(nums)-1, -nums[i]; lo < hi; {
				if nums[lo]+nums[hi] == sum {
					ret = append(ret, []int{nums[i], nums[lo], nums[hi]})
					for lo+1 < hi && nums[lo+1] == nums[lo] {
						lo++
					}
					for lo < hi-1 && nums[hi-1] == nums[hi] {
						hi--
					}
					lo, hi = lo+1, hi-1
				} else if nums[lo]+nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}
	return ret
}
