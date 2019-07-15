/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	v2i := make(map[int]int, len(nums)) // value to index
	for i := 0; i < len(nums); i++{
		dif := target - nums[i]
		if _, ok := v2i[dif]; ok{
			ret[0], ret[1] = v2i[dif], i
			break
		}
		v2i[nums[i]] = i
	}
	return ret
}
