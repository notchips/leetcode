/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
 func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	neg := false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		neg = true
	}
	dividend, divisor = abs(dividend), abs(divisor)
	ans := 0
	for dividend >= divisor {
		maxDivisor, p := divisor, 1
		for maxDivisor<<1 <= dividend {
			maxDivisor <<= 1
			p <<= 1
		}
		dividend -= maxDivisor
		ans += p
	}
	if neg {
		ans = -ans
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
