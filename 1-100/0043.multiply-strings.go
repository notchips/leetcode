/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	ans := "0"
	for i := len(num2) - 1; i >= 0; i-- {
		p := multiplySingle(num1, num2[i])
		ans = add(ans, p)
		num1 += "0"
	}
	return ans
}

func add(num1, num2 string) string {
	n := max(len(num1), len(num2))
	buf1 := make([]byte, n+1)
	buf2 := make([]byte, n+1)
	copy(buf1[n+1-len(num1):], num1)
	copy(buf2[n+1-len(num2):], num2)
	t := 0
	for i := n; i >= 0; i-- {
		sum := charToInt(buf1[i]) + charToInt(buf2[i]) + t
		buf1[i] = intToChar(sum % 10)
		t = sum / 10
	}
	if buf1[0] == '0' {
		return string(buf1[1:])
	}
	return string(buf1)
}

func multiplySingle(num string, c byte) string {
	if c == '0' {
		return "0"
	}
	if c == '1' {
		return num
	}
	buf := make([]byte, len(num)+1)
	t := 0
	for i := len(num) - 1; i >= 0; i-- {
		p := charToInt(num[i])*charToInt(c) + t
		buf[i+1] = intToChar(p % 10)
		t = p / 10
	}
	buf[0] = intToChar(t)
	if buf[0] == '0' {
		return string(buf[1:])
	}
	return string(buf)
}

func charToInt(c byte) int {
	if c == 0 {
		return 0
	}
	return int(c - '0')
}

func intToChar(n int) byte {
	return byte(n + '0')
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
