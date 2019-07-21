/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
	neg := false
	if n < 0 {
		n = -n
		neg = true
	}

	p := 1
	ret := 1.0
	tx := x
	for n > 0 { // 倍数逼近，同1029
		if n-p >= 0 {
			n -= p
			p *= 2
			ret *= tx
			tx *= tx
		} else {
			p = 1
			tx = x
		}
	}

	if neg {
		return 1 / ret
	}
	return ret
}
