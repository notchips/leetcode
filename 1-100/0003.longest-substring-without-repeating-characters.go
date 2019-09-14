/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	var hash [256]int // char到index的映射
	for i := range hash {
		hash[i] = -1
	}
	maxLen := 0
	for l, r := 0, 0; r < len(s); r++ {
		c := s[r]
		if hash[c] >= l {
			l = hash[c] + 1
		}
		hash[c] = r
		if curLen := r - l + 1; curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

// @lc code=end
