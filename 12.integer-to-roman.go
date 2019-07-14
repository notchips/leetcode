/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func intToRoman(num int) string {
	buff := make([]byte, 0, 20)
	for i := 1000; i > 0; i /= 10 {
		n := num / i
		num %= i
		switch i {
		case 1:
			singleToRoman(&buff, 'I', 'V', 'X', n)
		case 10:
			singleToRoman(&buff, 'X', 'L', 'C', n)
		case 100:
			singleToRoman(&buff, 'C', 'D', 'M', n)
		case 1000:
			singleToRoman(&buff, 'M', '-', '-', n)
		}
	}
	return string(buff)
}

/*
a b c
I V X
X L C
C D M
M - -
*/
func singleToRoman(buff *[]byte, a, b, c byte, n int) {
	switch {
	case 1 <= n && n <= 3:
		addChar(buff, a, n)
	case n == 4:
		addChar(buff, a, 1)
		addChar(buff, b, 1)
	case 5 <= n && n <= 8:
		addChar(buff, b, 1)
		addChar(buff, a, n-5)
	case n == 9:
		addChar(buff, a, 1)
		addChar(buff, c, 1)
	}
}

func addChar(buff *[]byte, char byte, n int) {
	for i := 0; i < n; i++ {
		*buff = append(*buff, char)
	}
}


