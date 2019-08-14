/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
 func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	ans := 0
	for i, neg := 0, false; i < len(str); i++ {
		// 检查正负号
		if i == 0 && (str[0] == '-' || str[0] == '+') {
			if str[0] == '-' {
				neg = true
			}
			continue
		}

		// 遇到非数字直接返回结果
		if str[i] < '0' || str[i] > '9' {
			return ans
		}

		ans = ans*10 + charToInt(str[i], neg)

		// 超过int32直接返回int32边界值
		if ans > math.MaxInt32 {
			return math.MaxInt32
		} else if ans < math.MinInt32 {
			return math.MinInt32
		}
	}
	return ans
}

func charToInt(c byte, neg bool) int {
	if neg {
		return -1 * int(c-'0')
	}
	return int(c - '0')
}
