/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
 func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			lo, hi := j+1, len(nums)-1
			for lo <= hi {
				mid := (lo + hi) / 2
				sum := nums[i] + nums[j] + nums[mid]
				if abs(sum - target) < abs(minDiff) {
					minDiff = sum - target
				}
				if sum == target {
					return target
				} else if sum < target {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
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
