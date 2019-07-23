/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	n := len(words)
	for i := n - 1; i >= 0; i-- {
		if len(words[i]) != 0 {
			return len(words[i])
		}
	}
	return 0
}
