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

	buff1 := make([]byte, n+1)
	buff2 := make([]byte, n+1)
	copy(buff1[n+1-len(num1):], num1)
	copy(buff2[n+1-len(num2):], num2)

	t := 0
	for i := n; i >= 0; i-- {
		sum := charToInt(buff1[i]) + charToInt(buff2[i]) + t
		buff1[i] = intToChar(sum % 10)
		t = sum / 10
	}
	if buff1[0] == '0' {
		return string(buff1[1:])
	}
	return string(buff1)
}

func multiplySingle(s string, c byte) string {
	if c == '0' {
		return "0"
	}
	if c == '1' {
		return s
	}
	buff := make([]byte, len(s)+1)
	t := 0
	for i := len(s) - 1; i >= 0; i-- {
		p := charToInt(s[i])*charToInt(c) + t
		buff[i+1] = intToChar(p % 10)
		t = p / 10
	}
	buff[0] = intToChar(t)
	if buff[0] == '0' {
		return string(buff[1:])
	}
	return string(buff)
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
