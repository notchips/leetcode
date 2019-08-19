/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	i := 0
LOOP:
	for ; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				break LOOP
			}
		}
	}
	return strs[0][:i]
}
