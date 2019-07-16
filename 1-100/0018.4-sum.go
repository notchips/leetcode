/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
 func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0, 10)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1; j < len(nums)-2; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					sum := target - nums[i] - nums[j]
					lo, hi := j+1, len(nums)-1
					for lo < hi {
						if nums[lo]+nums[hi] == sum {
							ret = append(ret, []int{nums[i], nums[j], nums[lo], nums[hi]})
							for lo < hi && nums[lo] == nums[lo+1] {
								lo++
							}
							for lo < hi && nums[hi] == nums[hi-1] {
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
		}
	}
	return ret
}
