/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

 var n2s = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"} // num to string

 func letterCombinations(digits string) []string {
	 var ret []string
	 if len(digits) == 0 {
		 return ret
	 }
	 dfs(&ret, "", digits, 0, len(digits)-1)
	 return ret
 }
 
 func dfs(ret *[]string, str string, digits string, i, n int) {
	 for _, c := range n2s[digits[i]-'0'] {
		 newStr := str +  string(c)
		 if i == n {
			 *ret = append(*ret, newStr)
		 } else {
			 dfs(ret, newStr, digits, i+1, n)
		 }
	 }
 }
