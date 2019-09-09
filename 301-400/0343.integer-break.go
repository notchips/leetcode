/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */
// 整数i可以拆分为2和i-2，2*(i-2) = 2i-4，当i>=4时,2i-4>=i，
// 故对于任意一个大于3的整数i都可以拆分为2和i-2，乘积更大
// 因此因子只需要2或3（1是最差的），又由于2*2*2 < 3*3
// 综上因子2不超过两个，其它因子全为3
func integerBreak(n int) int {
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 1
	case n == 3:
		return 2
	}

	switch n % 3 {
	case 1:
		return pow(3, n/3-1) * pow(2, 2)
	case 2:
		return pow(3, n/3) * 2
	default:
		return pow(3, n/3)
	}
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	if b&1 == 0 {
		t := pow(a, b/2)
		return t * t
	}
	return a * pow(a, b-1)
}
