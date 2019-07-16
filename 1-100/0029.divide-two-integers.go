/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
 func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	neg := true
	if dividend < 0 && divisor < 0 || dividend >= 0 && divisor > 0 {
		neg = false
	}

	dividend, divisor = abs(dividend), abs(divisor)
	ret, p := 0, 1
	oldDivisor := divisor

	for {
		for dividend-divisor < 0 && divisor > oldDivisor { // 除数太大，寻找较小的除数
			divisor >>= 1
			p >>= 1
		}
		if dividend < divisor {
			break
		} else {
			dividend -= divisor
			ret += p
			divisor <<= 1 // 每次运算后，除数翻倍
			p <<= 1
		}
	}

	if neg {
		ret = - ret
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
