/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package leetcode

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	pos := 0
LOOP:
	for ; ; pos++ {
		for _, s := range strs {
			if pos >= len(s) || s[pos] != strs[0][pos] {
				break LOOP
			}
		}
	}
	return strs[0][:pos]
}

// @lc code=end
