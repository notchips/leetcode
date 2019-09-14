/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */
func countAndSay(n int) string {
	say := []byte{'1'}
	for ; n > 1; n-- {
		var newSay []byte
		for i, cnt := 0, 1; i < len(say); i, cnt = i+1, cnt+1 {
			if i == len(say)-1 || say[i] != say[i+1] {
				newSay = append(append(newSay, strconv.Itoa(cnt)...), say[i])
				cnt = 0
			}
		}
		say = newSay
	}
	return string(say)
}
