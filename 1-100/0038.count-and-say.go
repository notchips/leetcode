/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */
func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		say := ""
		cnt := 1
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				cnt++
			} else {
				say += fmt.Sprintf("%d%c", cnt, s[i-1])
				cnt = 1
			}
		}
		say += fmt.Sprintf("%d%c", cnt, s[len(s)-1])
		s = say
	}
	return s
}
