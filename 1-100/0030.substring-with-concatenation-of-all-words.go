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
	totalCnt := 0
	for _, word := range words {
		if len(word) != wordLen {
			return nil
		}
		wordMap[word]++
		totalCnt++
	}

	var ans []int
	for i := 0; i < wordLen; i++ {
		curMap := make(map[string]int)
		curCnt := 0
		for l, r := i, i; l+wordLen*totalCnt <= len(s) && r+wordLen <= len(s); r += wordLen {
			rWord := s[r : r+wordLen]
			if _, ok := wordMap[rWord]; ok {
				curMap[rWord]++
				curCnt++
				for curMap[rWord] > wordMap[rWord] { // 移动左指针l，确保curMap中不存在多余单词
					curMap[s[l:l+wordLen]]--
					l += wordLen
					curCnt--
				}
				if curCnt == totalCnt {
					ans = append(ans, l)
				}
			} else { // 遇到不存在的单词，则重置起点l、cnt和map
				l = r + wordLen
				curCnt = 0
				curMap = make(map[string]int)
			}
		}
	}
	return ans
}

// @lc code=end
