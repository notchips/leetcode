/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
type charInfo struct {
	char  byte
	index int
}

func longestValidParentheses(s string) int {
	stack := make([]*charInfo, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1].char == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, &charInfo{s[i], i})
		}
	}
	if len(stack) == 0 {
		return len(s)
	}
	maxLen := max(stack[0].index, len(s)-stack[len(stack)-1].index-1)
	for i := 0; i < len(stack)-1; i++ {
		maxLen = max(maxLen, stack[i+1].index-stack[i].index-1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
