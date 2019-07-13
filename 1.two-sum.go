/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	v2i := make(map[int]int, len(nums)) // value to index
	for i := 0; i < len(nums); i++{
		diff := target - nums[i]
		if _, ok := v2i[diff]; ok{
			ret[0] = v2i[diff]
			ret[1] = i
			break
		}
		v2i[nums[i]] = i
	}
	return ret
}

