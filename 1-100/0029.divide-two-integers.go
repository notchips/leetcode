/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
package leetcode

import "math"

// @lc code=start
func divide(dividend int, divisor int) int {
	// 参数不合法
	if dividend < math.MinInt32 || dividend > math.MaxInt32 ||
		divisor < math.MinInt32 || divisor > math.MaxInt32 ||
		divisor == 0 {
		return 0
	}

	// 结果溢出
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// 除数和被除数都转换为正数
	neg := (dividend < 0 && divisor > 0) ||
		(dividend > 0 && divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	ans := 0
	for dividend >= divisor {
		maxDivisor, p := divisor, 1
		for (maxDivisor << 1) <= dividend {
			maxDivisor <<= 1
			p <<= 1
		}
		dividend -= maxDivisor
		ans += p
	}

	if neg {
		return -ans
	}
	return ans
}

// @lc code=end
