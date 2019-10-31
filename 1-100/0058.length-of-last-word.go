/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
package leetcode

// @lc code=start
func lengthOfLastWord(s string) int {
	length := 0
	meetWord := false
	for i := len(s) - 1; i >= 0; i-- {
		if meetWord && s[i] == ' ' {
			break
		}
		if s[i] != ' ' {
			length++
			meetWord = true
		}
	}
	return length
}

// @lc code=end
