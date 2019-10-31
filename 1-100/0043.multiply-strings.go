/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
package leetcode

// @lc code=start
func multiply(num1 string, num2 string) string {
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	buf1, buf2 := []byte(num1), []byte(num2)
	var ret []byte
	for i := len(buf2) - 1; i >= 0; i-- {
		ret = add(ret, multiplySingle(buf1, buf2[i]))
		buf1 = append(buf1, '0')
	}
	return string(ret)
}

func multiplySingle(num []byte, d byte) []byte {
	buf := append([]byte{'0'}, num...)
	carry := 0
	for i := len(buf) - 1; i > 0; i-- {
		product := charToInt(buf[i])*charToInt(d) + carry
		buf[i], carry = intToByte(product%10), product/10
	}
	buf[0] = intToByte(carry)
	return removeLeadingZero(buf)
}

func add(num1, num2 []byte) []byte {
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	buf := append([]byte{'0'}, num1...)
	carry := 0
	for i, j := len(buf)-1, len(num2)-1; i > 0; i-- {
		if j >= 0 {
			sum := charToInt(buf[i]) + charToInt(num2[j]) + carry
			buf[i], carry = intToByte(sum%10), sum/10
			j--
		} else if carry > 0 {
			sum := charToInt(buf[i]) + carry
			buf[i], carry = intToByte(sum%10), sum/10
		}
	}
	buf[0] = intToByte(carry)
	return removeLeadingZero(buf)
}

func charToInt(c byte) int {
	return int(c - '0')
}

func intToByte(i int) byte {
	return byte(i + '0')
}

func removeLeadingZero(num []byte) []byte {
	pos := 0
	for ; pos < len(num)-1; pos++ {
		if num[pos] != '0' {
			break
		}
	}
	return num[pos:]
}

// @lc code=end
