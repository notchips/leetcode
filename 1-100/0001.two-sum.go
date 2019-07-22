/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
 
/*
  时间复杂度：O(nlogn)；空间复杂度：O(n)
  √ 29/29 cases passed (4 ms)
  √ Your runtime beats 96.72 % of golang submissions
  √ Your memory usage beats 41.93 % of golang submissions (3.7 MB)
 */
func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	v2i := make(map[int]int) // value to index
	for i, v := range nums {
		if j, ok := v2i[target-v]; ok {
			ans[0], ans[1] = j, i
			break
		}
		v2i[v] = i
	}
	return ans
}
