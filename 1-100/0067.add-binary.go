/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */
package leetcode

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	if len(b) == 0 {
		return a
	}

	c := make([]byte, len(a)+1)
	carry := 0
	posA, posB, posC := len(a)-1, len(b)-1, len(c)-1
	for posB >= 0 {
		sum := byteToInt(a[posA]) + byteToInt(b[posB]) + carry
		c[posC], carry = intToByte(sum%2), sum/2
		posA, posB, posC = posA-1, posB-1, posC-1
	}
	for posA >= 0 {
		sum := byteToInt(a[posA]) + carry
		c[posC], carry = intToByte(sum%2), sum/2
		posA, posC = posA-1, posC-1
	}
	c[0] = intToByte(carry)
	if c[0] == '0' {
		return string(c[1:])
	}
	return string(c)
}

func intToByte(i int) byte {
	return byte(i) + '0'
}

func byteToInt(c byte) int {
	return int(c) - '0'
}

// @lc code=end
