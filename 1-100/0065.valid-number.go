/*
 * @lc app=leetcode id=65 lang=golang
 *
 * [65] Valid Number
 */
package leetcode

import "strings"

// @lc code=start
// var reg = regexp.MustCompile(`^[\s]*[-+]?(([\d]+([.][\d]*)?)|([.][\d]+))([eE][-+]?[\d]+)?[\s]*$`)

// func isNumber(s string) bool {
// 	return reg.MatchString(s)
// }

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	return isInteger(s) || isDecimal(s) || isScientificNumber(s)
}

func isInteger(s string) bool {
	s = removeSign(s)
	return isUnsignedInteger(s)
}

func isUnsignedInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func isDecimal(s string) bool {
	s = removeSign(s)
	nums := strings.Split(s, ".")
	if len(nums) != 2 {
		return false
	}
	return (len(nums[0]) == 0 && isUnsignedInteger(nums[1])) ||
		(isUnsignedInteger(nums[0]) && len(nums[1]) == 0) ||
		(isUnsignedInteger(nums[0]) && isUnsignedInteger(nums[1]))
}

func isScientificNumber(s string) bool {
	nums := strings.Split(s, "e")
	if len(nums) != 2 {
		return false
	}
	return (isInteger(nums[0]) || isDecimal(nums[0])) && isInteger(nums[1])
}

func removeSign(s string) string {
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		return s[1:]
	}
	return s
}

// @lc code=end
