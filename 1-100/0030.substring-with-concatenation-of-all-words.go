/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	ans := make([]int, 0, 16)
	wordLen := len(words[0])
	wordCnt := 0
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
		wordCnt++
	}
	for i := 0; i < wordLen; i++ {
		curCnt := 0
		curMap := make(map[string]int)
		for start, cur := i, i; start+wordCnt*wordLen <= len(s) && cur+wordLen <= len(s); {
			curWord := s[cur : cur+wordLen]
			if _, ok := wordMap[curWord]; ok {
				curCnt++
				curMap[curWord]++
				cur += wordLen
				for curMap[curWord] > wordMap[curWord] {
					preWord := s[start : start+wordLen]
					curCnt--
					curMap[preWord]--
					start += wordLen
				}
				if curCnt == wordCnt {
					ans = append(ans, start)
				}
			} else {
				start += wordLen
				cur = start
				curCnt = 0
				curMap = make(map[string]int)
			}

		}
	}
	return ans
}
