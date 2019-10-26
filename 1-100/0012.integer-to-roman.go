/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
package leetcode

// @lc code=start

func intToRoman(num int) string {
	hash := map[int]string{
		1:    "IVX",
		10:   "XLC",
		100:  "CDM",
		1000: "M--",
	}
	buff := make([]byte, 0, 20)
	for m := 1000; m > 0; m /= 10 {
		singleToRoman(&buff, hash[m][0], hash[m][1], hash[m][2], num/m)
		num %= m
	}
	return string(buff)
}

func singleToRoman(buff *[]byte, a, b, c byte, n int) {
	switch {
	case 1 <= n && n <= 3:
		addChars(buff, a, n)
	case n == 4:
		addChars(buff, a, 1)
		addChars(buff, b, 1)
	case 5 <= n && n <= 8:
		addChars(buff, b, 1)
		addChars(buff, a, n-5)
	case n == 9:
		addChars(buff, a, 1)
		addChars(buff, c, 1)
	}
}

func addChars(buff *[]byte, char byte, n int) {
	for i := 0; i < n; i++ {
		*buff = append(*buff, char)
	}
}

// @lc code=end
