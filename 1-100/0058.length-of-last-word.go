/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	n := len(words)
	if n == 0 {
		return 0
	}
	return len(words[n-1])
}
