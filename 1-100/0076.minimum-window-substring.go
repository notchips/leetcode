/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
package leetcode

import (
	"math"
)

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	var charMap [256]int // charMap记录t中字符出现的次数
	for i := 0; i < len(t); i++ {
		charMap[t[i]]++
	}

	minLen, minIdx := math.MaxInt64, 0
	var curMap [256]int // curMap记录遍历过程中，存在于t中的字符的次数
	cnt := 0
	// 双指针遍历
	for left, right := 0, 0; right < len(s); right++ {
		// 如果s[right]存在于t中，则记录其出现次数
		if charMap[s[right]] > 0 {
			curMap[s[right]]++
			if curMap[s[right]] <= charMap[s[right]] {
				cnt++
			}
			// 存在解
			if cnt == len(t) {
				// 移动left指针
				for charMap[s[left]] == 0 || curMap[s[left]] > charMap[s[left]] {
					if curMap[s[left]] > charMap[s[left]] {
						curMap[s[left]]--
					}
					left++
				}
				if right-left+1 <= minLen {
					minLen = right - left + 1
					minIdx = left
				}
			}
		}
	}
	if cnt < len(t) {
		return ""
	}
	return s[minIdx : minIdx+minLen]
}

// @lc code=end
