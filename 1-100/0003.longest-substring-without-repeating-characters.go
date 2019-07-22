/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

/*
  时间复杂度：O(n)；空间复杂度：O(1)
  √ 987/987 cases passed (4 ms)
  √ Your runtime beats 91.89 % of golang submissions
  √ Your memory usage beats 82.72 % of golang submissions (2.6 MB)
*/
func lengthOfLongestSubstring(s string) int {
	var left, ans int
	c2i := make([]int, 256) // char to index
	for i := range c2i {
		c2i[i] = -1
	}
	for i, c := range s {
		if oi := c2i[c]; oi >= left {
			left = oi + 1
		}
		if i-left+1 > ans {
			ans = i - left + 1
		}
		c2i[c] = i
	}
	return ans
}
