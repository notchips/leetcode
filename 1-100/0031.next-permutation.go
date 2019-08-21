/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] { //has answer
			reverse(nums[i+1:])
			j := firstBigger(nums, i)
			nums[i], nums[j] = nums[j], nums[i]
			return
		}
	}
	reverse(nums)
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func firstBigger(nums []int, i int) int {
	target := nums[i]
	left, right := i+1, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
