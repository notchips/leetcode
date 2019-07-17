/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
 func nextPermutation(nums []int) {
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				asc(nums[i+1:])
				return
			}
		}
	}
	asc(nums)
}

func asc(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
