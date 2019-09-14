/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	pos := 0
LOOP:
	for ; ; pos++ {
		for _, str := range strs {
			if pos >= len(str) || str[pos] != strs[0][pos] {
				break LOOP
			}
		}
	}
	return strs[0][:pos]
}
