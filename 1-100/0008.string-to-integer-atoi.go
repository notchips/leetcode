/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package leetcode

import (
	"math"
)

// @lc code=start
func myAtoi(str string) int {
	// 移除前置空格
	pos := 0
	for pos < len(str) && str[pos] == ' ' {
		pos++
	}
	if pos >= len(str) {
		return 0
	}
	str = str[pos:]

	ans := 0
	neg := false
	for i := 0; i < len(str); i++ {
		if i == 0 && (str[i] == '-' || str[i] == '+') {
			neg = str[i] == '-'
			continue
		}

		if str[i] < '0' || str[i] > '9' {
			break
		}

		ans = ans*10 + byteToInt(str[i], neg)

		if ans < math.MinInt32 {
			return math.MinInt32
		}
		if ans > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return ans
}

func byteToInt(c byte, neg bool) int {
	i := int(c) - '0'
	if neg {
		return -i
	}
	return i
}

// @lc code=end
