/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
 func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := math.MaxInt32
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	prefix := make([]byte, 0, 30)
	for i := 0; i < minLen; i++ {
		f := true
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				f = false
				break
			}
		}
		if f {
			prefix = append(prefix, strs[0][i])
		} else {
			break
		}
	}
	return string(prefix)
}

