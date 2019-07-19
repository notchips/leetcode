/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
func jump(nums []int) int {
	n := len(nums)
	cnt := 0
	for i := 0; i < n-1; {
		if i+nums[i] >= n-1 {
			cnt++
			break
		} else {
			max := i + 1
			for j := i + 2; j < n && j <= i+nums[i]; j++ {
				if nums[j]+j >= nums[max]+max {
					max = j
				}
			}
			cnt++
			i = max
		}
	}
	return cnt
}
