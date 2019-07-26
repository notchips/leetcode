/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	hash := make(map[int]int) // value to index
	for i := 0; i < len(nums); i++ {
		toFind := target - nums[i]
		if _, ok := hash[toFind]; ok {
			ans[0], ans[1] = hash[toFind], i
			break
		}
		hash[nums[i]] = i
	}
	return ans
}
