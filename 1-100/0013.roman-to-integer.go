/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
func romanToInt(s string) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		ret += singleToInt(s[i])
		if i > 0 {
			switch {
			case s[i-1] == 'I' && (s[i] == 'V' || s[i] == 'X'):
				ret -= 2
			case s[i-1] == 'X' && (s[i] == 'L' || s[i] == 'C'):
				ret -= 20
			case s[i-1] == 'C' && (s[i] == 'D' || s[i] == 'M'):
				ret -= 200
			}
		}
	}
	return ret
}

func singleToInt(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
