/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 */
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	return checkInteger(s) || checkDecimal(s) || checkScientificNum(s)
}

func checkInteger(s string) bool {
	s = removeSign(s)
	return checkNoSignInteger(s)
}

func checkNoSignInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func checkDecimal(s string) bool {
	s = removeSign(s)
	return checkNoSignDecimal(s)
}

func checkNoSignDecimal(s string) bool {
	num := strings.Split(s, ".")
	if len(num) != 2 {
		return false
	}
	// 小数点前半部分可以为空或整数，后半部分也可以为空或整数，但不能全为空
	if len(num[0]) == 0 && len(num[1]) == 0 {
		return false
	}
	return (len(num[0]) == 0 || checkNoSignInteger(num[0])) && (len(num[1]) == 0 || checkNoSignInteger(num[1]))
}

func checkScientificNum(s string) bool {
	num := strings.Split(s, "e")
	if len(num) != 2 {
		return false
	}
	// e前半部分可以为整数或小数，后半部分必须为整数
	return (checkInteger(num[0]) || checkDecimal(num[0])) && checkInteger(num[1])
}

func removeSign(s string) string {
	if len(s) == 0 {
		return ""
	}
	if s[0] == '+' || s[0] == '-' {
		return s[1:]
	}
	return s
}
