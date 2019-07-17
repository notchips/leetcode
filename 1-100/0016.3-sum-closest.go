/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			diff := nums[i] + nums[lo] + nums[hi] - target
			if diff == 0 {
				return target
			} else if diff > 0 {
				hi--
			} else {
				lo++
			}
			if abs(diff) < abs(minDiff) {
				minDiff = diff
			}
		}
	}
	return target + minDiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
