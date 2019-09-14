/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int) // num到index的映射
	for i, num := range nums {
		toFind := target - num
		if j, ok := hash[toFind]; ok {
			return []int{j, i}
		}
		hash[num] = i
	}
	return nil
}

// @lc code=end
