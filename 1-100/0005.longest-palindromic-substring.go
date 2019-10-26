/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package leetcode

import "math"

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 马拉车算法：https://www.cnblogs.com/xiaochuan94/p/10971443.html
	// 预处理，将原字符串转换为新字符串，统一长度为奇数
	// 例如将aba转变成 ('\1' '\0' 'a' '\0' 'b' '\0' 'a' '\0' '\2')
	buf := make([]byte, 0, 2*len(s)+3)
	buf = append(buf, 1)
	for i := 0; i < len(s); i++ {
		buf = append(buf, 0, s[i])
	}
	buf = append(buf, 0, 2)

	p := make([]int, len(buf)) // p[i] 表示buf数组中，以i为中心的最长回文半径
	// 索引      0 1 2 3 4 5 6 7 8
	// buf[i]    1 0 a 0 b 0 a 0 2
	// p[i]        1 2 1 4 1 2 1
	// 原字符串和处理后字符串之间推导关系：回文子串i在s中长度为p[i]-1，起始索引为(i-p[i])/2

	mid, right := 0, 0     // mid是buf数组所有回文子串中，能延伸到最右端的回文子串的中心位置，right是其右开边界
	maxLen, maxIdx := 0, 0 // 记录s中最长回文子串长度及其起始索引

	for i := 1; i < len(buf)-1; i++ { // 每次遍历确定以i为中心的最长回文半径
		if right > i { // 如果i在右边界内，则直接更新其p[i]值
			j := 2*mid - i // j是i关于mid的对称点
			p[i] = int(math.Min(float64(p[j]), float64(right-i)))
		} else { // i超过了右边界，p[i]初始值为1
			p[i] = 1
		}

		// 向左右两边延伸，扩展以i为中点回文半径
		// 预处理头尾加入的 '\1' 和 '\2' 就是为了防止这里越界
		for buf[i+p[i]] == buf[i-p[i]] {
			p[i]++
		}

		// 如果当前回文子串i的右边界值超过了right，则更新right和mid
		if p[i]+i > right {
			right = p[i] + i
			mid = i
		}

		// 更新maxLen和maxIdx
		if p[i]-1 > maxLen {
			maxLen = p[i] - 1
			maxIdx = (i - p[i]) / 2
		}
	}
	return s[maxIdx : maxIdx+maxLen]
}

// @lc code=end
