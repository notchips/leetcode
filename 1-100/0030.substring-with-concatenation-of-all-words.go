/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */
package leetcode

// @lc code=start
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordLen := len(words[0])
	wordMap := make(map[string]int)
	wordCnt := 0
	for _, word := range words {
		wordMap[word]++
		wordCnt++
	}

	var ans []int
	for i := 0; i < wordLen; i++ {
		curMap := make(map[string]int)
		curCnt := 0
		for l, r := i, i; l+wordLen*wordCnt <= len(s) && r+wordLen <= len(s); r += wordLen {
			rWord := s[r : r+wordLen]
			if _, ok := wordMap[rWord]; ok {
				curMap[rWord]++
				curCnt++
				for curMap[rWord] > wordMap[rWord] { // 移动左指针l，确保curMap中不存在多余单词
					curMap[s[l:l+wordLen]]--
					curCnt--
					l += wordLen
				}
				if curCnt == wordCnt {
					ans = append(ans, l)
				}
			} else { // 遇到不存在的单词，则重置起点l、cnt和map
				l = r + wordLen
				curMap = make(map[string]int)
				curCnt = 0
			}
		}
	}
	return ans
}

// @lc code=end
