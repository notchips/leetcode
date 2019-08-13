/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	hash := make(map[int]int) // value to index
	for i, num := range nums {
		toFind := target - num
		if _, ok := hash[toFind]; ok {
			ans[0], ans[1] = hash[toFind], i
			break
		}
		hash[num] = i
	}
	return ans
}
