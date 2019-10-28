/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */
package leetcode

// @lc code=start
func countAndSay(n int) string {
	say := []byte{'1'}
	for i := 1; i < n; i++ {
		newSay := make([]byte, 0, len(say))
		for j, cnt := 0, 1; j < len(say); j, cnt = j+1, cnt+1 {
			if j == len(say)-1 || say[j] != say[j+1] {
				newSay = append(newSay, byte(cnt+'0'), say[j])
				cnt = 0
			}
		}
		say = newSay
	}
	return string(say)
}

// @lc code=end
