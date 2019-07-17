/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
func generateParenthesis(n int) []string {
	ret := make([]string, 0, 20)
	helper(&ret, "", 0, 0, 2*n)
	return ret
}

// cnt1: count of '('
// cnt2: count of ')'
func helper(ret *[]string, str string, cnt1, cnt2, total int) {
	if cnt1 < total/2 {
		helper(ret, str+"(", cnt1+1, cnt2, total)
	}
	if cnt2 < cnt1 {
		helper(ret, str+")", cnt1, cnt2+1, total)
	}
	if cnt1 == total/2 && cnt1 == cnt2 {
		*ret = append(*ret, str)
	}
}
