/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package leetcode

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	for i := 0; i < len(digits); i++ {
		if digits[i] < '2' || digits[i] > '9' {
			return nil
		}
	}
	hash := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := make([]string, 0, 20)
	buf := make([]byte, 0, len(digits))
	dfs17(digits, hash, &ans, &buf)
	return ans
}

func dfs17(digits string, hash []string, ans *[]string, buf *[]byte) {
	if len(digits) == 0 {
		*ans = append(*ans, string(*buf))
		return
	}
	d := int(digits[0] - '0')
	for i := 0; i < len(hash[d]); i++ {
		*buf = append(*buf, hash[d][i])
		dfs17(digits[1:], hash, ans, buf)
		*buf = (*buf)[:len(*buf)-1]
	}
}

// @lc code=end
