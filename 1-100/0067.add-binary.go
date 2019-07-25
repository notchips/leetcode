/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	if len(a) == 0 {
		return ""
	}

	// 最多进2位
	bufA := make([]byte, 2, len(a)+2)
	bufB := make([]byte, 2+len(a)-len(b), len(a)+2)
	bufA = append(bufA, a...)
	bufB = append(bufB, b...)

	carry := 0
	for i := len(bufA) - 1; i >= 0; i-- {
		sum := toInt(bufA[i]) + toInt(bufB[i]) + carry
		bufA[i], carry = toChar(sum%2), sum/2
	}

	if bufA[0] != '0' {
		return string(bufA)
	} else if bufA[1] != '0' {
		return string(bufA[1:])
	}
	return string(bufA[2:])
}

func toChar(a int) byte {
	return byte(a + '0')
}

func toInt(c byte) int {
	if c == 0 {
		return 0
	}
	return int(c - '0')
}
