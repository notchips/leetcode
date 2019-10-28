/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
package leetcode

// @lc code=start
type char struct {
	c byte
	i int
}

func longestValidParentheses(s string) int {
	// stack存储所有不合法字符及其索引
	stack := make([]*char, 0, len(s))
	for i := 0; i < len(s); i++ {
		if n := len(stack); n > 0 && stack[n-1].c == '(' && s[i] == ')' {
			stack = stack[:n-1]
		} else {
			stack = append(stack, &char{s[i], i})
		}
	}

	// 字符全部合法
	if len(stack) == 0 {
		return len(s)
	}

	// 根据不合法字符索引之间的间隔，可以得出所有合法字符长度
	// 例如"()(()()("，存在两个不合法字符，对应索引为2和7，则可以得出合法序列有s[:2]和s[3:7]
	maxGap := stack[0].i
	postGap := len(s) - stack[len(stack)-1].i - 1
	if postGap > maxGap {
		maxGap = postGap
	}
	for i := 0; i < len(stack)-1; i++ {
		gap := stack[i+1].i - stack[i].i - 1
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

// @lc code=end
