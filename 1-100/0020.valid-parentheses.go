/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package leetcode

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if n := len(stack); n > 0 && isPair(stack[n-1], s[i]) {
			stack = stack[:n-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isPair(c1, c2 byte) bool {
	return (c1 == '(' && c2 == ')') ||
		(c1 == '{' && c2 == '}') ||
		(c1 == '[' && c2 == ']')
}

// @lc code=end
