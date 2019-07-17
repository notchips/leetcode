/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */
 func findSubstring(s string, words []string) []int {
	ret := make([]int, 0, 10)
	if len(s) == 0 || len(words) == 0 {
		return ret
	}

	wordLen := len(words[0])
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < wordLen; i++ {
		wordCnt := make(map[string]int)
		cnt := 0
		for start, cur := i, i; start <= len(s)-wordLen*len(words) && cur <= len(s)-wordLen; {
			subStr := s[cur : cur+wordLen]
			if _, ok := wordMap[subStr]; ok {
				wordCnt[subStr]++
				cnt++
				cur += wordLen
				for wordCnt[subStr] > wordMap[subStr] {
					preStr := s[start : start+wordLen]
					wordCnt[preStr]--
					cnt--
					start += wordLen
				}
				if cnt == len(words) {
					ret = append(ret, start)
				}
			} else { // reset
				wordCnt = make(map[string]int)
				cnt = 0
				start = cur + wordLen
				cur = start
			}
		}
	}
	return ret
}
