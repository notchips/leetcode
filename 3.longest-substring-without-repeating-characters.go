/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	var (
		start = 0
		ret   = 0
		c2i = make(map[rune]int) // char to index
	)
	for i, c := range s {
		if oi, ok := c2i[c]; ok && oi >= start {
			start = oi + 1
		}
		c2i[c] = i
		ret = max(ret, i-start+1)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

