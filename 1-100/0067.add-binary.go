/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	n := len(a) + 1
	bufA := make([]byte, n)
	bufB := make([]byte, n)
	for i := 0; i < n; i++ {
		bufA[i], bufB[i] = '0', '0'
	}
	copy(bufA[n-len(a):], a)
	copy(bufB[n-len(b):], b)
	carry := 0
	for i := n - 1; i > 0; i-- {
		sum := charToInt(bufA[i]) + charToInt(bufB[i]) + carry
		bufA[i], carry = intToChar(sum%2), sum/2
	}
	bufA[0] = intToChar(carry)
	if bufA[0] == '0' {
		return string(bufA[1:])
	}
	return string(bufA)
}

func intToChar(i int) byte {
	return byte(i) + '0'
}

func charToInt(c byte) int {
	return int(c) - '0'
}
